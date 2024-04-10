package employee

import employeeservice "github.com/manish-sa/india-template/internal/service/employee"

type APIEmployee struct {
	service employeeservice.ServiceEmployee
}

func NewEmployeeApi() *APIEmployee {
	return &APIEmployee{}
}
