package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type EmployeeJSON struct {
	Id           int    `json:"id,omitempty"`
	CardNumberId string `json:"card_number_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	WarehouseId  int    `json:"warehouse_id,omitempty"`
}

type ResponseBody struct {
	Data any `json:"data"`
}

type EmployeeHandler struct {
	sv interfaces.IEmployeeService
}

func CreateEmployeeHandler(service interfaces.IEmployeeService) *EmployeeHandler {
	return &EmployeeHandler{sv: service}
}
func (e *EmployeeHandler) GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := e.sv.GetEmployees()

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ResponseBody{Data: nil})
		return
	}

	var employeesJSON = make([]EmployeeJSON, 0)

	for _, employee := range data {
		employeesJSON = append(employeesJSON, EmployeeJSON{
			Id:           employee.Id,
			CardNumberId: employee.CardNumberId,
			FirstName:    employee.FirstName,
			LastName:     employee.LastName,
			WarehouseId:  employee.WarehouseId,
		})
	}

	response.JSON(w, http.StatusOK, ResponseBody{Data: employeesJSON})

}
