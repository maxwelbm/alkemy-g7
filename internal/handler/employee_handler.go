package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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

func (e *EmployeeHandler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ResponseBody{Data: nil})
		return
	}

	data, err := e.sv.GetEmployeeById(idInt)

	if err != nil && errors.Is(err.(custom_error.CustomError).Err, custom_error.NotFound) {
		response.JSON(w, http.StatusNotFound, ResponseBody{Data: nil})
		return
	}

	employeeJSON := EmployeeJSON{
		Id:           data.Id,
		CardNumberId: data.CardNumberId,
		FirstName:    data.FirstName,
		LastName:     data.FirstName,
		WarehouseId:  data.WarehouseId,
	}

	response.JSON(w, http.StatusOK, ResponseBody{Data: employeeJSON})
}
