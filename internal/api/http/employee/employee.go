package employee

import (
	"context"

	"github.com/manish-sa/india-template/internal/api/http/oapi"
	employeeservice "github.com/manish-sa/india-template/internal/service/employee"
)

type APIEmployee struct {
	EmpService employeeservice.ServiceEmployee
}

func NewEmployee(empService employeeservice.ServiceEmployee) APIEmployee {
	return APIEmployee{
		EmpService: empService,
	}
}

func (api *APIEmployee) SaveEmployee(
	ctx context.Context, request oapi.SaveEmployeeRequestObject,
) (oapi.SaveEmployeeResponseObject, error) {
	return oapi.SaveEmployee201JSONResponse{Id: nil, Name: "Manish"}, nil
}
