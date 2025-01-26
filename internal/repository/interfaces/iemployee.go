package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IEmployeeRepo interface {
	Get() ([]model.Employee, error)
	GetByID(id int) (model.Employee, error)
	Update(id int, employee model.Employee) (model.Employee, error)
	Post(employee model.Employee) (model.Employee, error)
	Delete(id int) error
	GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error)
	GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error)
}
