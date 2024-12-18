package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IEmployeeRepo interface {
	Get() (map[int]model.Employee, error)
}
