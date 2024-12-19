package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IEmployeeRepo interface {
	Get() (map[int]model.Employee, error)
	GetById(id int) (model.Employee, error)
	Update(id int, employee model.Employee) (model.Employee, error)
	Post(employee model.Employee) (model.Employee, error)
}
