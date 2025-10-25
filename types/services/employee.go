package employee

type CreateEmployeeRequest struct {
	FullName         string `json:"fullName" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	DepartmentId     *uint  `json:"departmentId"`
	EmploymentTypeId *uint  `json:"employementTypeId"`
}
