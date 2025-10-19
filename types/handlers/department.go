package handlers

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type DepartmentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
