package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
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

func (e *EmployeeJSON) toEmployeeEntity() *model.Employee {
	return &model.Employee{
		Id:           e.Id,
		CardNumberId: e.CardNumberId,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		WarehouseId:  e.WarehouseId,
	}
}

func (e *EmployeeJSON) fromEmployeeEntity(employee model.Employee) {
	e.Id = employee.Id
	e.CardNumberId = employee.CardNumberId
	e.FirstName = employee.FirstName
	e.LastName = employee.LastName
	e.WarehouseId = employee.WarehouseId
}

type ResponseBody struct {
	Data any `json:"data"`
}

type ResponseBodyError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
		response.JSON(w, http.StatusBadRequest, nil)
		return
	}

	data, err := e.sv.GetEmployeeById(idInt)

	if err != nil {
		if errors.Is(err.(custom_error.CustomError).Err, custom_error.NotFound) {
			response.JSON(w, http.StatusNotFound, ResponseBodyError{Status: "error", Message: err.Error()})
			return
		}

		response.JSON(w, http.StatusInternalServerError, ResponseBodyError{Status: "error", Message: err.Error()})
		return
	}

	employeeJSON := EmployeeJSON{}
	employeeJSON.fromEmployeeEntity(data)

	response.JSON(w, http.StatusOK, ResponseBody{Data: employeeJSON})
}

func (e *EmployeeHandler) InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee EmployeeJSON
	err := request.JSON(r, &newEmployee)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ResponseBodyError{Status: "error", Message: "error parsing the request body"})
		return
	}

	employee := newEmployee.toEmployeeEntity()

	data, err := e.sv.InsertEmployee(*employee)

	if err != nil {
		if errors.Is(err.(custom_error.CustomError).Err, custom_error.InvalidErr) {
			response.JSON(w, http.StatusUnprocessableEntity, ResponseBodyError{Status: "error", Message: err.Error()})
			return
		}
		response.JSON(w, http.StatusInternalServerError, ResponseBodyError{Status: "error", Message: err.Error()})
		return
	}

	var employeeJSON EmployeeJSON
	employeeJSON.fromEmployeeEntity(data)

	response.JSON(w, http.StatusCreated, ResponseBody{Data: employeeJSON})
}

func (e *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ResponseBodyError{Status: "error", Message: "error parsing the id in path param"})
		return
	}

	var reqBody EmployeeJSON

	err = request.JSON(r, &reqBody)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, ResponseBodyError{Status: "error", Message: "error parsing the request body"})
		return
	}

	employee := *reqBody.toEmployeeEntity()

	updatedEmployee, err := e.sv.UpdateEmployee(idInt, employee)

	if err != nil {
		if errors.Is(err.(custom_error.CustomError).Err, custom_error.NotFound) {
			response.JSON(w, http.StatusNotFound, ResponseBodyError{Status: "error", Message: err.Error()})
			return
		}
		response.JSON(w, http.StatusInternalServerError, ResponseBodyError{Status: "error", Message: err.Error()})
		return
	}

	employeeJSON := EmployeeJSON{}
	employeeJSON.fromEmployeeEntity(updatedEmployee)

	response.JSON(w, http.StatusOK, ResponseBody{Data: employeeJSON})
}
