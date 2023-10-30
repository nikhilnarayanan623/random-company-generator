package interfaces

import (
	"company-service/pkg/domain"
	"context"
)

type CompanyRepo interface {
	// one
	CreateCompany(ctx context.Context, company *domain.Company) error
	CreateDepartment(ctx context.Context, department *domain.Department) error
	CreateTeam(ctx context.Context, team *domain.Team) error
	CreteEmployee(ctx context.Context, employee *domain.Employee) error

	//many
	CreateDepartments(ctx context.Context, departments []*domain.Department) error
	CreateTeams(ctx context.Context, team []*domain.Team) error
	CreteEmployees(ctx context.Context, employee []*domain.Employee) error
}
