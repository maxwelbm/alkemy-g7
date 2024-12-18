package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type EmployeeService struct {
	rp interfaces.IEmployeeRepo
}

func CreateEmployeeService(rp interfaces.IEmployeeRepo) *EmployeeService {
	return &EmployeeService{rp: rp}
}

func (e *EmployeeService) GetEmployees() (map[int]model.Employee, error) {
	data, err := e.rp.Get()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *EmployeeService) GetEmployeeById(id int) (model.Employee, error) {
	return e.rp.GetById(id)
}
