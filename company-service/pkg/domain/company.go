package domain

import "time"


type Company struct {
	Name           string
	Industry       string
	CEO            *Employee
	Departments    []*Department
	TotalEmployees int
}

type Department struct {
	Name           string
	Leader         *Employee
	Teams          []*Team
	TotalEmployees int
}

type Team struct {
	Name           string
	Manager        *Employee
	Members        []*Employee
	TotalEmployees int
}

type Employee struct {
	ID       string
	Name     string
	Age      int
	Email    string
	Salary   float64
	HireDate time.Time
	Role     string
}
