package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
	"github.com/stretchr/testify/assert"
)

func TestGetEmployeesHandler(t *testing.T) {
	findAll := func() ([]model.Employee, error) {
		return []model.Employee{
			{Id: 1, CardNumberId: "1", FirstName: "John", LastName: "Cena", WarehouseId: 1},
			{Id: 2, CardNumberId: "2", FirstName: "Martha", LastName: "Piana", WarehouseId: 2}}, nil
	}
	employeeHd := EmployeeHandler{
		sv: &StubMockService{FuncGetEmployees: findAll},
	}

	t.Run("should return a list of employees", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		employeeHd.GetEmployeesHandler(res, req)
		expected := `{"data":[{"id":1,"card_number_id":"1","first_name":"John","last_name":"Cena","warehouse_id":1},{"id":2,"card_number_id":"2","first_name":"Martha","last_name":"Piana","warehouse_id":2}]}`
		assert.Equal(t, res.Code, http.StatusOK)
		assert.Equal(t, res.Body.String(), expected)
	})

	findAll = func() ([]model.Employee, error) {
		return []model.Employee{}, errors.New("something went wrong")
	}

	employeeHd = EmployeeHandler{
		sv: &StubMockService{FuncGetEmployees: findAll},
	}
	t.Run("should return 500 internal error in case of unexpected error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		employeeHd.GetEmployeesHandler(res, req)
		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.Equal(t, res.Body.String(), expected)
	})

	findAll = func() ([]model.Employee, error) {
		return []model.Employee{}, customError.EmployeeErrNotFound
	}

	employeeHd = EmployeeHandler{
		sv: &StubMockService{FuncGetEmployees: findAll},
	}

	t.Run("should return error in case of expected error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		employeeHd.GetEmployeesHandler(res, req)

		expected := `{"message":"employee not found"}`
		assert.Equal(t, res.Code, http.StatusNotFound)
		assert.Equal(t, res.Body.String(), expected)
	})
}

func TestGetEmployeeById(t *testing.T) {
	getById := func(id int) (model.Employee, error) {
		return model.Employee{Id: 1, CardNumberId: "1", FirstName: "John", LastName: "Cena", WarehouseId: 1}, nil
	}
	employeeHd := EmployeeHandler{
		sv: &StubMockService{FuncGetEmployeeById: getById},
	}

	r := chi.NewRouter()
	r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeById)

	t.Run("should return the employee requested and 200 ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"data":{"id":1,"card_number_id":"1","first_name":"John","last_name":"Cena","warehouse_id":1}}`
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, res.Body.String(), expected)
	})

	t.Run("should return a not found when employee id not exists", func(t *testing.T) {
		getById := func(id int) (model.Employee, error) {
			return model.Employee{}, customError.EmployeeErrNotFound
		}
		employeeHd := EmployeeHandler{
			sv: &StubMockService{FuncGetEmployeeById: getById},
		}

		r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeById)

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

		r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeById)

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "", res.Body.String())
	})

	t.Run("should return an error in case of unexpected error", func(t *testing.T) {
		getById := func(id int) (model.Employee, error) {
			return model.Employee{}, errors.New("unexpected error")
		}
		employeeHd := EmployeeHandler{
			sv: &StubMockService{FuncGetEmployeeById: getById},
		}

		req := httptest.NewRequest("GET", "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.Get("/api/v1/employees/{id}", employeeHd.GetEmployeeById)

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
	insertEmployee := func(employee model.Employee) (model.Employee, error) {
		employee.Id = 1
		return employee, nil
	}
	employeeHd := EmployeeHandler{
		sv: &StubMockService{FuncInsertEmployee: insertEmployee},
	}
	newEmployee := EmployeeJSON{
		CardNumberId: "#123",
		FirstName:    "Islam",
		LastName:     "Makhachev",
		WarehouseId:  1,
	}
	employeeJSON, err := json.Marshal(newEmployee)
	if err != nil {
		t.Fatalf("failed to marshal new employee: %v", err)
	}

	t.Run("should return 201 created and the new employee created", func(t *testing.T) {
		req := createRequest(string(employeeJSON))
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)

		expected := `{"data":{"id":1,"card_number_id":"#123","first_name":"Islam","last_name":"Makhachev","warehouse_id":1}}`

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return 422 unprocessable entity when the input is missing fields", func(t *testing.T) {
		newEmployee := `
		{
			"first_name": "islam"
		}`

		insertEmployee := func(employee model.Employee) (model.Employee, error) {
			return model.Employee{}, customError.EmployeeErrInvalid
		}

		employeeHd := EmployeeHandler{
			sv: &StubMockService{FuncInsertEmployee: insertEmployee},
		}

		req := createRequest(newEmployee)
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)

		expected := `{"message":"invalid employeee"}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("should return 409 conflict when cardnumberid already exists", func(t *testing.T) {
		insertEmployee := func(employee model.Employee) (model.Employee, error) {
			return model.Employee{}, customError.EmployeeErrDuplicatedCardNumber
		}

		employeeHd := EmployeeHandler{
			sv: &StubMockService{FuncInsertEmployee: insertEmployee},
		}

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
		insertEmployee := func(employee model.Employee) (model.Employee, error) {
			return model.Employee{}, errors.New("unexpected error")
		}

		employeeHd := EmployeeHandler{
			sv: &StubMockService{FuncInsertEmployee: insertEmployee},
		}

		req := createRequest(string(employeeJSON))
		res := httptest.NewRecorder()

		employeeHd.InsertEmployee(res, req)
		expected := `{"message":"something went wrong"}`
		assert.Equal(t, res.Code, http.StatusInternalServerError)
		assert.Equal(t, res.Body.String(), expected)
	})
}

func TestUpdateEmployee(t *testing.T) {
	updateRequest := func(body string) *http.Request {
		req := httptest.NewRequest("PATCH", "/api/v1/employees/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		return req
	}
	updateEmployee := func(id int, employee model.Employee) (model.Employee, error) {
		return model.Employee{Id: id, CardNumberId: "1", FirstName: employee.FirstName, LastName: "Cena", WarehouseId: 1}, nil
	}
	employeeHd := EmployeeHandler{
		sv: &StubMockService{FuncUpdateEmployee: updateEmployee},
	}
	r := chi.NewRouter()
	r.Patch("/api/v1/employees/{id}", employeeHd.UpdateEmployee)

	newEmployee := `
	{
		"first_name": "Miguel"
	}
	`
	t.Run("should return 200 ok and the employee with the new data", func(t *testing.T) {
		req := updateRequest(newEmployee)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"data":{"id":1,"card_number_id":"1","first_name":"Miguel","last_name":"Cena","warehouse_id":1}}`
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, res.Body.String(), expected)

	})

	t.Run("should return 404 not found when employee not found", func(t *testing.T) {
		updateEmployee := func(id int, employee model.Employee) (model.Employee, error) {
			return model.Employee{}, customError.EmployeeErrNotFound
		}
		employeeHd := EmployeeHandler{
			sv: &StubMockService{FuncUpdateEmployee: updateEmployee},
		}
		r.Patch("/api/v1/employees/{id}", employeeHd.UpdateEmployee)

		req := updateRequest(newEmployee)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"message":"employee not found"}`
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, res.Body.String(), expected)
	})
}

func TestDeleteEmployee(t *testing.T) {
	deleteEmployee := func(id int) error {
		return nil
	}
	employeeHd := EmployeeHandler{
		sv: &StubMockService{FuncDeleteEmployee: deleteEmployee},
	}
	r := chi.NewRouter()
	r.Delete("/api/v1/employees/{id}", employeeHd.DeleteEmployee)

	t.Run("should return 204 no content when delete with success", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/api/v1/employees/2", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
		assert.Equal(t, "", res.Body.String())

	})

	t.Run("should return 404 not found when employee id does not exist", func(t *testing.T) {
		deleteEmployee := func(id int) error {
			return customError.EmployeeErrNotFound
		}
		employeeHd := EmployeeHandler{
			sv: &StubMockService{FuncDeleteEmployee: deleteEmployee},
		}
		r.Delete("/api/v1/employees/{id}", employeeHd.DeleteEmployee)

		req := httptest.NewRequest("DELETE", "/api/v1/employees/2", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"message":"employee not found"}`
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expected, res.Body.String())

	})
}

// StubMockService
type StubMockService struct {
	FuncGetEmployees                     func() ([]model.Employee, error)
	FuncGetEmployeeById                  func(id int) (model.Employee, error)
	FuncUpdateEmployee                   func(id int, employee model.Employee) (model.Employee, error)
	FuncInsertEmployee                   func(employee model.Employee) (model.Employee, error)
	FuncDeleteEmployee                   func(id int) error
	FuncGetInboundOrdersReportByEmployee func(employeeId int) (model.InboundOrdersReportByEmployee, error)
	FuncGetInboundOrdersReports          func() ([]model.InboundOrdersReportByEmployee, error)
}

func (s *StubMockService) GetEmployees() ([]model.Employee, error) {
	return s.FuncGetEmployees()
}
func (s *StubMockService) GetEmployeeById(id int) (model.Employee, error) {
	return s.FuncGetEmployeeById(id)
}
func (s *StubMockService) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	return s.FuncUpdateEmployee(id, employee)
}
func (s *StubMockService) InsertEmployee(employee model.Employee) (model.Employee, error) {
	return s.FuncInsertEmployee(employee)
}
func (s *StubMockService) DeleteEmployee(id int) error {
	return s.FuncDeleteEmployee(id)
}
func (s *StubMockService) GetInboundOrdersReportByEmployee(employeeId int) (model.InboundOrdersReportByEmployee, error) {
	return s.FuncGetInboundOrdersReportByEmployee(employeeId)
}
func (s *StubMockService) GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error) {
	return s.FuncGetInboundOrdersReports()
}
