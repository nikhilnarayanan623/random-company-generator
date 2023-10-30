package store

import "sync"

type departmentDetail struct {
	Name                  string
	PercentageOfEmployees int
}

var (
	mu sync.RWMutex
	// department and percentage should be calculated logically by ourself
	departmentsDetails = [...]departmentDetail{ // distributing 100 percentage to all departments
		{
			Name:                  "Development",
			PercentageOfEmployees: 65,
		},
		{
			Name:                  "HR",
			PercentageOfEmployees: 9,
		},
		{
			Name:                  "Sales",
			PercentageOfEmployees: 13,
		},
		{
			Name:                  "Marketing",
			PercentageOfEmployees: 13,
		},
	}

	departmentTeamNames = map[string][]string{
		"Development": {"Frontend Team", "Backend Team", "QA Team"},
		"HR":          {"Recruitment Team", "Employee Relations Team"},
		"Sales":       {"Sales Team"},
		"Marketing":   {"Marketing Team"},
	}

	teamRoles = map[string][]string{
		// Development
		"Frontend Team": {"Frontend Developer", "UI/UX Designer", "Quality Assurance Engineer"},
		"Backend Team":  {"Backend Developer", "Database Administrator", "DevOps Engineer"},
		"QA Team":       {"Quality Assurance Engineer", "Test Automation Engineer"},
		// HR
		"Recruitment Team":        {"HR Manager", "Recruiter", "HR Coordinator"},
		"Employee Relations Team": {"HR Manager", "Employee Relations Specialist"},
		// Sales
		"Sales Team": {"Sales Manager", "Account Executive"},
		// Marketing
		"Marketing Team": {"Marketing Manager", "Content Specialist"},
	}
)
