package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IEmployeeService interface {
	GetEmployees() (map[int]model.Employee, error)
}
