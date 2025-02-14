package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEmployeesHandler(t *testing.T) {
	srv := mocks.NewMockIEmployeeService(t)
	employeeHd := EmployeeHandler{
		sv: srv,
	}

	t.Run("should return a list of employees", func(t *testing.T) {
		srv.On("GetEmployees", mock.Anything).Return([]model.Employee{
			{ID: 1, CardNumberID: "1", FirstName: "John", LastName: "Cena", WarehouseID: 1},
			{ID: 2, CardNumberID: "2", FirstName: "Martha", LastName: "Piana", WarehouseID: 2}}, nil).Once()

		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		employeeHd.GetEmployeesHandler(res, req)
		expected := `{"data":[{"id":1,"card_number_id":"1","first_name":"John","last_name":"Cena","warehouse_id":1},{"id":2,"card_number_id":"2","first_name":"Martha","last_name":"Piana","warehouse_id":2}]}`
		assert.Equal(t, res.Code, http.StatusOK)
		assert.JSONEq(t, res.Body.String(), expected)
	})

	t.Run("should return 500 internal error in case of unexpected error", func(t *testing.T) {
		srv.On("GetEmployees", mock.Anything).Return([]model.Employee{}, errors.New("something went wrong")).Once()

		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		employeeHd.GetEmployeesHandler(res, req)
		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.JSONEq(t, res.Body.String(), expected)
	})

	t.Run("should return error in case of expected error", func(t *testing.T) {
		srv.On("GetEmployees", mock.Anything).Return([]model.Employee{}, customerror.EmployeeErrNotFound).Once()

		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		employeeHd.GetEmployeesHandler(res, req)

		expected := `{"message":"employee not found"}`
		assert.Equal(t, res.Code, http.StatusNotFound)
		assert.Equal(t, res.Body.String(), expected)
	})
}

func TestGetEmployeeById(t *testing.T) {
	srv := mocks.NewMockIEmployeeService(t)
	employeeHd := EmployeeHandler{
		sv: srv,
	}

	r := chi.NewRouter()
	r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeByID)

	t.Run("should return the employee requested and 200 ok", func(t *testing.T) {
		srv.On("GetEmployeeByID", mock.Anything).Return(model.Employee{ID: 1, CardNumberID: "1", FirstName: "John", LastName: "Cena", WarehouseID: 1}, nil).Once()
		req := httptest.NewRequest("GET", "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"data":{"id":1,"card_number_id":"1","first_name":"John","last_name":"Cena","warehouse_id":1}}`
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, res.Body.String(), expected)
	})

	t.Run("should return a not found when employee id not exists", func(t *testing.T) {
		srv.On("GetEmployeeByID", mock.Anything).Return(model.Employee{}, customerror.EmployeeErrNotFound).Once()

		r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeByID)

		req := httptest.NewRequest("GET", "/api/v1/employees/2", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"message":"employee not found"}`
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return an error in case of invalid id type", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/employees/number", nil)
		res := httptest.NewRecorder()

		r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeByID)

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "", res.Body.String())
	})

	t.Run("should return an error in case of unexpected error", func(t *testing.T) {
		srv.On("GetEmployeeByID", mock.Anything).Return(model.Employee{}, errors.New("unexpected error")).Once()

		req := httptest.NewRequest("GET", "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeByID)

		r.ServeHTTP(res, req)

		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.Equal(t, res.Body.String(), expected)
	})
}

func TestInsertEmployee(t *testing.T) {
	createRequest := func(body string) *http.Request {
		req := httptest.NewRequest("POST", "/api/v1/employees", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		return req
	}

	srv := mocks.NewMockIEmployeeService(t)
	employeeHd := EmployeeHandler{
		sv: srv,
	}

	newEmployee := EmployeeJSON{
		CardNumberID: "#123",
		FirstName:    "Islam",
		LastName:     "Makhachev",
		WarehouseID:  1,
	}

	employeeJSON, err := json.Marshal(newEmployee)
	if err != nil {
		t.Fatalf("failed to marshal new employee: %v", err)
	}

	t.Run("should return 201 created and the new employee created", func(t *testing.T) {
		mockEmployee := model.Employee{
			ID:           1,
			CardNumberID: "#123",
			FirstName:    "Islam",
			LastName:     "Makhachev",
			WarehouseID:  1,
		}
		srv.On("InsertEmployee", mock.Anything).Return(mockEmployee, nil).Once()
		req := createRequest(string(employeeJSON))
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)

		expected := `{"data":{"id":1,"card_number_id":"#123","first_name":"Islam","last_name":"Makhachev","warehouse_id":1}}`

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 422 unprocessable entity when the input is missing fields", func(t *testing.T) {
		srv.On("InsertEmployee", mock.Anything).Return(model.Employee{}, customerror.EmployeeErrInvalid).Once()
		newEmployee := `
		{
			"first_name": "islam"
		}`

		req := createRequest(newEmployee)
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)

		expected := `{"message":"invalid employeee"}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return 409 conflict when cardnumberid already exists", func(t *testing.T) {
		srv.On("InsertEmployee", mock.Anything).Return(model.Employee{}, customerror.EmployeeErrDuplicatedCardNumber).Once()

		req := createRequest(string(employeeJSON))
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)

		expected := `{"message":"duplicated card number id"}`

		assert.Equal(t, http.StatusConflict, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return 400 bad request when invalid request body", func(t *testing.T) {
		reqBody := `something`
		req := createRequest(string(reqBody))
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)

		expected := `{"message":"error parsing the request body"}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return 500 internal error in case of unexpected error", func(t *testing.T) {
		srv.On("InsertEmployee", mock.Anything).Return(model.Employee{}, errors.New("unexpected error")).Once()

		req := createRequest(string(employeeJSON))
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)
		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.JSONEq(t, res.Body.String(), expected)
	})
}

func TestUpdateEmployee(t *testing.T) {
	updateRequest := func(body string) *http.Request {
		req := httptest.NewRequest("PATCH", "/api/v1/employees/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		return req
	}
	srv := mocks.NewMockIEmployeeService(t)

	employeeHd := EmployeeHandler{
		sv: srv,
	}
	r := chi.NewRouter()
	r.Patch("/api/v1/employees/{id}", employeeHd.UpdateEmployee)

	newEmployee := `
	{
		"first_name": "Miguel"
	}
	`
	t.Run("should return 200 ok and the employee with the new data", func(t *testing.T) {
		srv.On("UpdateEmployee", mock.Anything, mock.Anything).Return(model.Employee{ID: 1, CardNumberID: "1", FirstName: "Miguel", LastName: "Cena", WarehouseID: 1}, nil).Once()

		req := updateRequest(newEmployee)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"data":{"id":1,"card_number_id":"1","first_name":"Miguel","last_name":"Cena","warehouse_id":1}}`
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, res.Body.String(), expected)

	})

	t.Run("should return 404 not found when employee not found", func(t *testing.T) {
		srv.On("UpdateEmployee", mock.Anything, mock.Anything).Return(model.Employee{}, customerror.EmployeeErrNotFound).Once()

		r.Patch("/api/v1/employees/{id}", employeeHd.UpdateEmployee)

		req := updateRequest(newEmployee)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"message":"employee not found"}`
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, res.Body.String(), expected)
	})

	t.Run("should return an error in case of invalid id type", func(t *testing.T) {
		req := httptest.NewRequest("PATCH", "/api/v1/employees/something", nil)
		res := httptest.NewRecorder()

		r.Patch("/api/v1/employees/{id}", employeeHd.UpdateEmployee)

		r.ServeHTTP(res, req)

		expected := `{"message":"error parsing the id in path param"}`
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return an error in case of invalid requestBody", func(t *testing.T) {
		req := updateRequest("")
		res := httptest.NewRecorder()

		r.Patch("/api/v1/employees/{id}", employeeHd.UpdateEmployee)

		r.ServeHTTP(res, req)

		expected := `{"message":"error parsing the request body"}`
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return 500 internal error in case of unexpected error", func(t *testing.T) {
		srv.On("UpdateEmployee", mock.Anything, mock.Anything).Return(model.Employee{}, errors.New("unexpected error")).Once()

		req := updateRequest(newEmployee)
		res := httptest.NewRecorder()

		r.Patch("/api/v1/employees/{id}", employeeHd.UpdateEmployee)
		r.ServeHTTP(res, req)

		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.Equal(t, res.Body.String(), expected)
	})
}

func TestDeleteEmployee(t *testing.T) {
	srv := mocks.NewMockIEmployeeService(t)

	employeeHd := EmployeeHandler{
		sv: srv,
	}
	r := chi.NewRouter()
	r.Delete("/api/v1/employees/{id}", employeeHd.DeleteEmployee)

	t.Run("should return 204 no content when delete with success", func(t *testing.T) {
		srv.On("DeleteEmployee", mock.Anything).Return(nil).Once()

		req := httptest.NewRequest("DELETE", "/api/v1/employees/2", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
		assert.Equal(t, "", res.Body.String())

	})

	t.Run("should return 404 not found when employee id does not exist", func(t *testing.T) {
		srv.On("DeleteEmployee", mock.Anything).Return(customerror.EmployeeErrNotFound).Once()

		r.Delete("/api/v1/employees/{id}", employeeHd.DeleteEmployee)

		req := httptest.NewRequest("DELETE", "/api/v1/employees/2", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"message":"employee not found"}`
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expected, res.Body.String())

	})
	t.Run("should return an error in case of invalid id type", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/api/v1/employees/something", nil)
		res := httptest.NewRecorder()

		r.Delete("/api/v1/employees/{id}", employeeHd.DeleteEmployee)

		r.ServeHTTP(res, req)

		expected := `{"message":"error parsing the id in path param"}`
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return 500 internal error in case of unexpected error", func(t *testing.T) {
		srv.On("DeleteEmployee", mock.Anything).Return(errors.New("unexpected")).Once()

		req := httptest.NewRequest("DELETE", "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.Delete("/api/v1/employees/{id}", employeeHd.DeleteEmployee)
		r.ServeHTTP(res, req)

		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.Equal(t, res.Body.String(), expected)
	})
}

func TestGetInboundOrdersReports(t *testing.T) {
	createRequest := func(query string) *http.Request {
		req := httptest.NewRequest("GET", "/api/v1/reportInboundOrders?id="+query, nil)
		return req
	}

	srv := mocks.NewMockIEmployeeService(t)
	employeeHd := EmployeeHandler{
		sv: srv,
	}

	t.Run("should return 200 OK and reports when no ID is provided", func(t *testing.T) {
		srv.On("GetInboundOrdersReports").Return([]model.InboundOrdersReportByEmployee{
			{ID: 1, CardNumberID: "#123", FirstName: "Jon", LastName: "Jones", WarehouseID: 1, InboundOrdersCount: 20},
			{ID: 2, CardNumberID: "#456", FirstName: "Islam", LastName: "Makachev", WarehouseID: 6, InboundOrdersCount: 26},
		}, nil).Once()

		req := createRequest("")
		res := httptest.NewRecorder()

		employeeHd.GetInboundOrdersReports(res, req)

		expected := `{
    "data": [
        {
            "id": 1,
            "card_number_id": "#123",
            "first_name": "Jon",
            "last_name": "Jones",
            "warehouse_id": 1,
            "inbound_orders_count": 20
        },
        {
            "id": 2,
            "card_number_id": "#456",
            "first_name": "Islam",
            "last_name": "Makachev",
            "warehouse_id": 6,
            "inbound_orders_count": 26
        }
		]
	}`
		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 400 bad request for invalid ID", func(t *testing.T) {
		req := createRequest("abc")
		res := httptest.NewRecorder()

		employeeHd.GetInboundOrdersReports(res, req)

		expected := `{"message":"invalid employee id"}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return error in case of expected error without ID", func(t *testing.T) {
		srv.On("GetInboundOrdersReports").Return(nil, customerror.EmployeeErrNotFoundInboundOrders).Once()

		req := createRequest("")
		res := httptest.NewRecorder()

		employeeHd.GetInboundOrdersReports(res, req)

		expected := ` {"message":"inboud orders not found"}`

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 500 internal server error when service fails without ID", func(t *testing.T) {
		srv.On("GetInboundOrdersReports", mock.Anything).Return(nil, errors.New("something went wrong")).Once()

		req := createRequest("")
		res := httptest.NewRecorder()

		employeeHd.GetInboundOrdersReports(res, req)

		expected := `{"message":"something went wrong"}`
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 200 OK and reports for valid employee ID", func(t *testing.T) {
		id := 1
		srv.On("GetInboundOrdersReportByEmployee", id).Return(model.InboundOrdersReportByEmployee{ID: 1, CardNumberID: "#123", FirstName: "Jon", LastName: "Jones", WarehouseID: 1, InboundOrdersCount: 20}, nil).Once()

		req := createRequest(strconv.Itoa(id))
		res := httptest.NewRecorder()

		employeeHd.GetInboundOrdersReports(res, req)

		expected := `{"data":{"id":1,"card_number_id":"#123","first_name":"Jon","last_name":"Jones","warehouse_id":1,"inbound_orders_count":20}}`
		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return 500 internal server error when getting report by employee fails", func(t *testing.T) {
		id := 1
		srv.On("GetInboundOrdersReportByEmployee", id).Return(model.InboundOrdersReportByEmployee{}, errors.New("something went wrong")).Once()

		req := createRequest(strconv.Itoa(id))
		res := httptest.NewRecorder()

		employeeHd.GetInboundOrdersReports(res, req)

		expected := `{"message":"something went wrong"}`
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("should return error in case of expected error with ID", func(t *testing.T) {
		srv.On("GetInboundOrdersReportByEmployee", mock.Anything).Return(model.InboundOrdersReportByEmployee{}, customerror.EmployeeErrNotFoundInboundOrders).Once()

		req := createRequest("12")
		res := httptest.NewRecorder()

		employeeHd.GetInboundOrdersReports(res, req)

		expected := ` {"message":"inboud orders not found"}`

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})
}
