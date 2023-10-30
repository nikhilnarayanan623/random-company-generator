package usecase

import (
	"company-service/pkg/domain"
	"company-service/pkg/models"
	repoInterfaces "company-service/pkg/repository/interfaces"
	"company-service/pkg/service/random"
	"company-service/pkg/store"
	"company-service/pkg/usecase/interfaces"
	"company-service/pkg/utils"
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type companyUseCase struct {
	randomGen   random.RandomGenerator
	companyRepo repoInterfaces.CompanyRepo
}

func NewCompanyUseCase(
	randomGen random.RandomGenerator,
	companyRepo repoInterfaces.CompanyRepo,
) interfaces.CompanyUseCase {

	usecase := companyUseCase{
		randomGen:   randomGen,
		companyRepo: companyRepo,
	}

	return &usecase
}

const (
	maxWaitDurationToStartAGoroutine = time.Second * 2

	// salary range in month(change the value according to your preference)
	CEOMinSalary            = float64(100_000)
	CEOMaxSalary            = float64(400_000)
	DepartmentLeadMinSalary = float64(70_000)
	DepartmentLeadMaxSalary = float64(200_000)
	TeamManagerMinSalary    = float64(50_000)
	TeamManagerMaxSalary    = float64(150_000)
	TeamMemberMinSalary     = float64(30_000)
	TeamMemberMaxSalary     = float64(150_000)

	// wait duration
	// maximum wait time for a goroutine to send an error to error channel.
	//in case of multiple goroutine try to send error through the un buffered error channel, then wait duration helps to avoid deadlock.
	MaxWaitDurationForErrSendToChannel = time.Millisecond * 100
)

func (c *companyUseCase) Create(ctx context.Context, companyReq models.CompanyRequest) (*domain.Company, error) {

	// create a copy of departments
	departmentDetails := store.GetAllDepartmentNamesAndDetails()

	// create departments same size for the department details
	departments := make([]*domain.Department, len(departmentDetails))

	// create each department for company
	for i := range departmentDetails {

		// make the department request
		departmentReq := makeDepartmentRequest(
			departmentDetails[i].Name,
			companyReq.TotalEmployees,
			departmentDetails[i].PercentageOfEmployees,
		)

		// before creating a new department check the request got cancelled
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// create a new department and append to departments.
			department, err := c.newGenerateDepartment(ctx, departmentReq)
			if err != nil {
				return nil, err
			}
			departments[i] = department
		}

	}

	// save the departments on db
	if err := c.companyRepo.CreateDepartments(ctx, departments); err != nil {
		return nil, fmt.Errorf("failed to save departments on db: %w", err)
	}

	// create random CEO
	ceo := c.randomGen.GenerateEmployee(CEOMinSalary, CEOMaxSalary)
	ceo.Role = fmt.Sprintf("CEO")

	// create a company with requested details and generated department.
	company := &domain.Company{
		Name:           companyReq.Name,
		CEO:            ceo,
		Industry:       "IT Industry",
		TotalEmployees: companyReq.TotalEmployees,
		Departments:    departments,
	}

	// save the company on db
	if err := c.companyRepo.CreateCompany(ctx, company); err != nil {

		return nil, fmt.Errorf("failed to save company details on db: %w", err)
	}

	return company, nil
}

func (c *companyUseCase) newGenerateDepartment(
	ctx context.Context,
	departmentReq departmentRequest,
) (*domain.Department, error) {

	startChan := make(chan struct{}, runtime.NumCPU())
	errChan := make(chan error)

	// make a slice to store distributed team members count in team.
	randTeamMembersCounts := make([]int, 0)

	var (
		employeesNeedToCreate = departmentReq.requiredEmployees
		employeesForTeam      int
	)

	for employeesNeedToCreate > 1 { // the last employee is department leader

		// if required employees become lesser than max employees in a team then assign all employees count to this team
		if employeesNeedToCreate <= departmentReq.maxEmployeesInATeam {
			employeesForTeam = employeesNeedToCreate - 1 // exclude the leader
		}

		// selecting team count randomly making some team to be with 1 or 2 members to avoid that using the below switch
		switch {
		case employeesNeedToCreate <= departmentReq.maxEmployeesInATeam: // if the employee need to create is lesser than the max employees in a team then we can choose all left members to that team.
			employeesForTeam = employeesNeedToCreate - 1 // exclude the lead
		case employeesNeedToCreate <= (departmentReq.maxEmployeesInATeam * 2): // if employee need to create in a team is less than double of max employees in a team, then split the left over employees half for each team
			employeesForTeam = employeesNeedToCreate / 2 // and in the next iteration the above select will be selected
		default: //in all other case select a random employees count for the team
			employeesForTeam = utils.GetIntBetween(departmentReq.minEmployeesInATeam, departmentReq.maxEmployeesInATeam)
		}

		// append the new value
		randTeamMembersCounts = append(randTeamMembersCounts, employeesForTeam)

		// update the employeesNeedToCreate
		employeesNeedToCreate -= employeesForTeam
	}

	// now we have total team count and distributed the total employees count.
	// make a slice of team with the length of randTeamMembersCounts
	teams := make([]*domain.Team, len(randTeamMembersCounts))

	// create a wait group to wait for all new goroutines to complete
	var wg sync.WaitGroup
	// fmt.Println("2")
	// range the the teams and generate the teams concurrently
	for i := range teams {
		// create a team request
		teamReq := teamRequest{
			name:        c.randomGen.GetRandomTeamNameByDepartmentName(departmentReq.name),
			totalMember: randTeamMembersCounts[i], // select the team member count
			// pass the slice and index to to set the value
			team: &teams[i],
		}

		select {
		case <-ctx.Done(): // check the request is cancelled
			return nil, ctx.Err()

		case err := <-errChan: // check any error on running goroutine
			return nil, err

		case startChan <- struct{}{}: //check availability of firing new goroutine
			// generate the team concurrently
			wg.Add(1)
			go func() {

				defer wg.Done()
				c.generateTeam(ctx, errChan, teamReq)
				fmt.Println("team gen completed: ", teamReq)
				// release the start channel
				<-startChan
			}()
		}
	}
	// create leader fo department
	leader := c.randomGen.GenerateEmployee(DepartmentLeadMinSalary, DepartmentLeadMaxSalary)
	leader.Role = fmt.Sprintf("%s Leader", departmentReq.name)

	// save the department leader on employee db
	if err := c.companyRepo.CreteEmployee(ctx, leader); err != nil {

		return nil, fmt.Errorf("failed to save departments on db: %w", err)
	}

	// wait for all goroutines to complete
	wg.Wait()

	select {
	// check any error has been sended after firing all goroutine.
	case err := <-errChan:

		return nil, err
	default: // if no error then save teams on db and return the department

		// store the teams on db
		if err := c.companyRepo.CreateTeams(ctx, teams); err != nil {
			fmt.Println("before error: ", teams, "err: ", err)
			fmt.Println("h:::: ", departmentReq)
			return nil, fmt.Errorf("failed to save teams details on db: %w", err)
		}

		// create a department with request details and generated teams and return it.
		return &domain.Department{
			Name:           departmentReq.name,
			Leader:         leader,
			TotalEmployees: departmentReq.requiredEmployees,
			Teams:          teams,
		}, nil
	}
}

func (c *companyUseCase) generateTeam(
	ctx context.Context,
	errChan chan<- error,
	teamReq teamRequest,
) {

	// make slice of members to store random employees(exclude team manager)
	members := make([]*domain.Employee, teamReq.totalMember-1)

	// fill the employees
	for i := range members {
		select {
		case <-ctx.Done():
			return
		default:
			employee := c.randomGen.GenerateEmployee(TeamMemberMinSalary, TeamMemberMaxSalary)
			// randomly select a role for the employee based on the team
			employee.Role = c.randomGen.GetRandomRoleByTeamName(teamReq.name)
			members[i] = employee
		}
	}

	// save the team members(employees) on db
	if err := c.companyRepo.CreteEmployees(ctx, members); err != nil {
		// use a select to send error channel to avoid deadlock when no one is there's to receive error
		select {
		case <-time.After(MaxWaitDurationForErrSendToChannel):
		case errChan <- fmt.Errorf("failed to save employees on db: %w", err):
		}
		return
	}

	// create a manager as the last employee
	var manager *domain.Employee
	select {
	case <-ctx.Done():
		return
	default:
		manager = c.randomGen.GenerateEmployee(TeamManagerMinSalary, TeamManagerMaxSalary)
		manager.Role = fmt.Sprintf("%s Manager", teamReq.name)
	}

	// save the manager on employee db
	if err := c.companyRepo.CreteEmployee(ctx, manager); err != nil {
		select {
		case <-time.After(MaxWaitDurationForErrSendToChannel):
		case errChan <- fmt.Errorf("failed to save employee on db: %w", err):
		}
		return
	}

	// insert the generated team to to given team pointer object
	(*teamReq.team) = &domain.Team{
		Name:           teamReq.name,
		TotalEmployees: teamReq.totalMember,
		Manager:        manager,
		Members:        members,
	}
}

// extra data need to create a department
type departmentRequest struct {
	name                string
	requiredEmployees   int
	minEmployeesInATeam int
	maxEmployeesInATeam int
}

// To create new department based on the company employees and employees percentage in department
func makeDepartmentRequest(
	departmentName string,
	totalEmployeesInCompany int,
	employeePercentageInDepartment int,
) departmentRequest {

	request := departmentRequest{
		name:              departmentName,
		requiredEmployees: (totalEmployeesInCompany / 100) * employeePercentageInDepartment, // calculate the total employees for this department
	}

	// according to the employee percentage assign the min and max employees in a team
	switch {
	case employeePercentageInDepartment <= 10: // department like HR
		request.minEmployeesInATeam = 5
		request.maxEmployeesInATeam = 8
	case employeePercentageInDepartment <= 40:
		request.minEmployeesInATeam = 8
		request.maxEmployeesInATeam = 12
	default: // department like Development
		request.minEmployeesInATeam = 12
		request.maxEmployeesInATeam = 22
	}

	return request
}

// details need to create a team
type teamRequest struct {
	name        string
	totalMember int
	team        **domain.Team
}

// To create a new team request
func makeTeamRequest(name string, totalMembers int) teamRequest {

	return teamRequest{
		name:        name,
		totalMember: totalMembers,
	}
}
