package employee

type CreateEmployeeRequest struct {
	FullName   string `json:"fullName" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Position   string `json:"position"`
	Department string `json:"department"`
}
