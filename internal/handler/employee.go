package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type EmployeeJSON struct {
	ID           int    `json:"id,omitempty"`
	CardNumberID string `json:"card_number_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	WarehouseID  int    `json:"warehouse_id,omitempty"`
}

func (e *EmployeeJSON) toEmployeeEntity() *model.Employee {
	return &model.Employee{
		ID:           e.ID,
		CardNumberID: e.CardNumberID,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		WarehouseID:  e.WarehouseID,
	}
}

func (e *EmployeeJSON) fromEmployeeEntity(employee model.Employee) {
	e.ID = employee.ID
	e.CardNumberID = employee.CardNumberID
	e.FirstName = employee.FirstName
	e.LastName = employee.LastName
	e.WarehouseID = employee.WarehouseID
}

type EmployeeHandler struct {
	sv  interfaces.IEmployeeService
	log logger.Logger
}

func CreateEmployeeHandler(service interfaces.IEmployeeService, log logger.Logger) *EmployeeHandler {
	return &EmployeeHandler{sv: service, log: log}
}

// GetEmployeesHandler retrieves all employees.
// @Summary Retrieve all employees
// @Description Fetch all registered employees from the database
// @Tags Employee
// @Produce json
// @Success 200 {object} handler.EmployeeJSON
// @Failure 404 {object} model.ErrorResponseSwagger "Employee not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to retrieve employee"
// @Router /employees [get]
func (e *EmployeeHandler) GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	e.log.Log("EmployeeHandler", "INFO", "initializing GetEmployeesHandler")

	data, err := e.sv.GetEmployees()

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to retrieve employees: %v", err))

		if err, ok := err.(*customerror.EmployeerErr); ok {
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

		return
	}

	var employeesJSON = make([]EmployeeJSON, 0)

	for _, employee := range data {
		employeesJSON = append(employeesJSON, EmployeeJSON{
			ID:           employee.ID,
			CardNumberID: employee.CardNumberID,
			FirstName:    employee.FirstName,
			LastName:     employee.LastName,
			WarehouseID:  employee.WarehouseID,
		})
	}

	e.log.Log("EmployeeHandler", "INFO", "GetEmployeesHandler finished successfully")
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", employeesJSON))
}

// GetEmployeeByID retrieves a single employee by ID.
// @Summary Retrieve a single employee
// @Description Fetch an employee by their ID
// @Tags Employee
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} handler.EmployeeJSON
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID format"
// @Failure 404 {object} model.ErrorResponseSwagger "Employee not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to retrieve employee"
// @Router /employees/{id} [get]
func (e *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	e.log.Log("EmployeeHandler", "INFO", "initializing GetEmployeeByID")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("invalid ID format: %v", err))
		response.JSON(w, http.StatusBadRequest, nil)

		return
	}

	data, err := e.sv.GetEmployeeByID(id)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to retrieve employee with ID %d: %v", id, err))

		if err, ok := err.(*customerror.EmployeerErr); ok {
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

		return
	}

	employeeJSON := EmployeeJSON{}
	employeeJSON.fromEmployeeEntity(data)

	e.log.Log("EmployeeHandler", "INFO", fmt.Sprintf("GetEmployeeByID finished successfully for employee ID: %d", id))
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", employeeJSON))
}

// InsertEmployee creates a new employee.
// @Summary Create a new employee
// @Description Add a new employee to the database
// @Tags Employee
// @Accept json
// @Produce json
// @Param employee body handler.EmployeeJSON true "Employee details"
// @Success 201 {object} handler.EmployeeJSON
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid request body"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to create employee"
// @Router /employees [post]
func (e *EmployeeHandler) InsertEmployee(w http.ResponseWriter, r *http.Request) {
	e.log.Log("EmployeeHandler", "INFO", "initializing InsertEmployee")

	var newEmployee EmployeeJSON
	err := request.JSON(r, &newEmployee)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to parse request body: %v", err))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("error parsing the request body", nil))

		return
	}

	employee := newEmployee.toEmployeeEntity()

	data, err := e.sv.InsertEmployee(*employee)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to insert employee: %v", err))

		if err, ok := err.(*customerror.EmployeerErr); ok {
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

		return
	}

	var employeeJSON EmployeeJSON

	employeeJSON.fromEmployeeEntity(data)

	e.log.Log("EmployeeHandler", "INFO", "InsertEmployee finished successfully")
	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", employeeJSON))
}

// UpdateEmployee updates an existing employee.
// @Summary Update an employee
// @Description Modify the details of an existing employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body handler.EmployeeJSON true "Updated employee details"
// @Success 200 {object} handler.EmployeeJSON
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid request or ID format"
// @Failure 404 {object} model.ErrorResponseSwagger "Employee not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to update employee"
// @Router /employees/{id} [put]
func (e *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	e.log.Log("EmployeeHandler", "INFO", "initializing UpdateEmployee")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("invalid ID format: %v", err))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("error parsing the id in path param", nil))

		return
	}

	var reqBody EmployeeJSON

	err = request.JSON(r, &reqBody)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to parse request body: %v", err))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("error parsing the request body", nil))

		return
	}

	employee := *reqBody.toEmployeeEntity()

	updatedEmployee, err := e.sv.UpdateEmployee(id, employee)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to update employee with ID %d: %v", id, err))

		if err, ok := err.(*customerror.EmployeerErr); ok {
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
			return
		} else {
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))
			return
		}
	}

	employeeJSON := EmployeeJSON{}
	employeeJSON.fromEmployeeEntity(updatedEmployee)

	e.log.Log("EmployeeHandler", "INFO", fmt.Sprintf("UpdateEmployee finished successfully for employee ID: %d", id))
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", employeeJSON))
}

// DeleteEmployee deletes an employee by ID.
// @Summary Delete an employee
// @Description Remove an employee from the database by their ID
// @Tags Employee
// @Param id path int true "Employee ID"
// @Success 204 "No content"
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID format"
// @Failure 404 {object} model.ErrorResponseSwagger "Employee not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to delete employee"
// @Router /employees/{id} [delete]
func (e *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	e.log.Log("EmployeeHandler", "INFO", "initializing DeleteEmployee")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("invalid ID format: %v", err))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("error parsing the id in path param", nil))

		return
	}

	err = e.sv.DeleteEmployee(id)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to delete employee with ID %d: %v", id, err))

		if err, ok := err.(*customerror.EmployeerErr); ok {
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

		return
	}

	e.log.Log("EmployeeHandler", "INFO", fmt.Sprintf("DeleteEmployee finished successfully for employee ID: %d", id))
	response.JSON(w, http.StatusNoContent, nil)
}

// GetInboundOrdersReports retrieves inbound order reports.
// @Summary Retrieve inbound order reports
// @Description Fetch inbound order reports, optionally filtering by employee ID
// @Tags Employee
// @Produce json
// @Param id query int false "Employee ID (optional)"
// @Success 200 {object} interface{} "Inbound order report(s)"
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID format"
// @Failure 404 {object} model.ErrorResponseSwagger "Reports not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to retrieve reports"
// @Router /employees/reports [get]
func (e *EmployeeHandler) GetInboundOrdersReports(w http.ResponseWriter, r *http.Request) {
	e.log.Log("EmployeeHandler", "INFO", "initializing GetInboundOrdersReports")

	id := r.URL.Query().Get("id")

	if id == "" {
		data, err := e.sv.GetInboundOrdersReports()

		if err != nil {
			e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to retrieve inbound orders reports: %v", err))

			if err, ok := err.(*customerror.EmployeerErr); ok {
				response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

			return
		}

		e.log.Log("EmployeeHandler", "INFO", "GetInboundOrdersReports finished successfully")
		response.JSON(w, http.StatusOK, responses.CreateResponseBody("", data))

		return
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("invalid employee ID format: %v", err))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid employee id", nil))

		return
	}

	data, err := e.sv.GetInboundOrdersReportByEmployee(idInt)

	if err != nil {
		e.log.Log("EmployeeHandler", "ERROR", fmt.Sprintf("failed to retrieve inbound orders report for employee ID %d: %v", idInt, err))

		if err, ok := err.(*customerror.EmployeerErr); ok {
			response.JSON(w, err.StatusCode, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("something went wrong", nil))

		return
	}

	e.log.Log("EmployeeHandler", "INFO", fmt.Sprintf("GetInboundOrdersReports finished successfully for employee ID: %d", idInt))
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", data))
}
