package repository

import (
	"company-service/pkg/domain"
	"company-service/pkg/repository/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionCompanies   = "companies"
	CollectionDepartments = "departments"
	CollectionTeams       = "teams"
	CollectionEmployees   = "employees"
)

type companyDB struct {
	db *mongo.Database
}

func NewCompanyRepo(db *mongo.Database) interfaces.CompanyRepo {
	return &companyDB{
		db: db,
	}
}

// Create company
func (c *companyDB) CreateCompany(ctx context.Context, company *domain.Company) error {

	_, err := c.db.Collection(CollectionCompanies).InsertOne(ctx, company)

	return err
}

// Create one department
func (c *companyDB) CreateDepartment(ctx context.Context, department *domain.Department) error {

	_, err := c.db.Collection(CollectionDepartments).InsertOne(ctx, department)

	return err
}

// Create one team
func (c *companyDB) CreateTeam(ctx context.Context, team *domain.Team) error {

	_, err := c.db.Collection(CollectionTeams).InsertOne(ctx, team)

	return err
}

// Create one employee
func (c *companyDB) CreteEmployee(ctx context.Context, employee *domain.Employee) error {

	_, err := c.db.Collection(CollectionEmployees).InsertOne(ctx, employee)

	return err
}

// Create multiple departments
func (c *companyDB) CreateDepartments(ctx context.Context, departments []*domain.Department) error {

	departmentDocuments := make([]interface{}, len(departments))

	for i, department := range departments {
		// check the department is nil before dereferencing
		if department == nil {
			continue
		}
		departmentDocuments[i] = *department
	}

	_, err := c.db.Collection(CollectionDepartments).InsertMany(ctx, departmentDocuments)
	if err != nil {
		return err
	}
	return nil
}

// Create multiple teams
func (c *companyDB) CreateTeams(ctx context.Context, teams []*domain.Team) error {

	teamDocuments := make([]interface{}, len(teams))

	for i, team := range teams {
		// check the teams is nil before dereferencing
		if team == nil {
			continue
		}
		teamDocuments[i] = *team
	}

	_, err := c.db.Collection(CollectionTeams).InsertMany(ctx, teamDocuments)
	if err != nil {
		return err
	}
	return nil
}

// Create multiple employees
func (c *companyDB) CreteEmployees(ctx context.Context, employees []*domain.Employee) error {

	employeeDocuments := make([]interface{}, len(employees))

	for i, employee := range employees {
		// check the employees is nil before dereferencing
		if employee == nil {
			continue
		}
		employeeDocuments[i] = *employee
	}

	_, err := c.db.Collection(CollectionEmployees).InsertMany(ctx, employeeDocuments)
	if err != nil {
		return err
	}
	return nil
}
