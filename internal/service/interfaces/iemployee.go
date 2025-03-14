package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IEmployeeService interface {
	GetEmployees() ([]model.Employee, error)
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployee(id int, employee model.Employee) (model.Employee, error)
	InsertEmployee(employee model.Employee) (model.Employee, error)
	DeleteEmployee(id int) error
	GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error)
	GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error)
}
