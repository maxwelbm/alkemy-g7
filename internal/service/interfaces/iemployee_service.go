package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IEmployeeService interface {
	GetEmployees() (map[int]model.Employee, error)
	GetEmployeeById(id int) (model.Employee, error)
	UpdateEmployee(id int, employee model.Employee) (model.Employee, error)
	InsertEmployee(employee model.Employee) (model.Employee, error)
}
