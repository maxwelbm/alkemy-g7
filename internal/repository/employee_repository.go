package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type EmployeeRepository struct {
	db map[int]model.Employee
}

func CreateEmployeeRepository(db map[int]model.Employee) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (e *EmployeeRepository) Get() (map[int]model.Employee, error) {
	return e.db, nil
}

func (e *EmployeeRepository) GetById(id int) (model.Employee, error) {
	employee, ok := e.db[id]

	if !ok {
		return model.Employee{}, custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	return employee, nil
}

func (e *EmployeeRepository) Update(id int, employee model.Employee) (model.Employee, error) {
	e.db[id] = employee
	return employee, nil
}
