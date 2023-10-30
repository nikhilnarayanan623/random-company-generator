package request

type CompanyRequest struct {
	Name           string `json:"company_name" binding:"required,min=3,max=15"`
	CEO            string `json:"ceo_name" binding:"required,min=3,max=25"`
	TotalEmployees int    `json:"total_employees" binding:"required,min=100,max=5000"`
}
