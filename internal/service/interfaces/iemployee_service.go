package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IEmployeeService interface {
	GetEmployees() (map[int]model.Employee, error)
	GetEmployeeById(id int) (model.Employee, error)
}
