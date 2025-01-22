package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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
			return model.Employee{}, custom_error.EmployeeErrNotFound
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
