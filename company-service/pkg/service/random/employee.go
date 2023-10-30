package random

import (
	"company-service/pkg/domain"
	"company-service/pkg/utils"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	employeeMinAge = 18
	employeeMaxAge = 50
)

var (
	companyStartTime = time.Date(2014, 11, 12, 0, 0, 0, 0, time.Local)
)

func (r *randomGenerator) GenerateEmployee(minSalary, maxSalary float64) *domain.Employee {

	// select a random name from names
	name := r.getRandomName()

	return &domain.Employee{
		ID:       utils.GenerateUUID(),
		Name:     name,
		Email:    generateEmail(name),
		Age:      utils.GetIntBetween(employeeMinAge, employeeMaxAge),
		Salary:   generateSalary(minSalary, maxSalary),
		HireDate: utils.GetTimeBetween(companyStartTime, time.Now().Add(-(time.Hour * 24 * 31 * 2))), // select start time as predefined and end time as 2 month before now.
	}
}

func (r *randomGenerator) GetRandomTeamNameByDepartmentName(department string) string {

	r.mu.RLock()
	defer r.mu.RUnlock()

	if teams, ok := r.departmentToRole[department]; ok {
		return teams[utils.GetRandomIndex(len(teams))]
	}

	return ""
}

func (r *randomGenerator) GetRandomRoleByTeamName(team string) string {

	r.mu.RLock()
	defer r.mu.RUnlock()

	if roles, ok := r.teamToRole[team]; ok {
		return roles[utils.GetRandomIndex(len(roles))]
	}

	return ""
}

func generateEmail(name string) string {

	// change the name to lowercase
	name = strings.ToLower(name)

	numChar := []byte("123456789")

	extraLetters := 4

	// cover the name to byte slice as email
	email := []byte(name)

	for i := 1; i <= extraLetters; i++ {
		// take a random number character from numChan and append to email
		email = append(email, numChar[rand.Intn(len(numChar))])
	}

	mailEnd := []byte("@email.com")
	// add the email ending to the email
	email = append(email, mailEnd...)

	// convert the mail slice to string and return
	return string(email)
}

func generateSalary(start, end float64) float64 {
	// get a random number in between start and end
	salary := utils.GetFloatBetween(start, end)

	// return the salary to a thousand multiplied value(like: 20000,30000)
	return math.Floor(salary/1000) * 1000
}
