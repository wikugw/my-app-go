package repositories

type EmployeeDetail struct {
	ID             uint   `json:"id"`
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
	DepartmentName string `json:"departmentName"`
	EmploymentType string `json:"employmentType"`
	HireDate       string `json:"hireDate"`
}
