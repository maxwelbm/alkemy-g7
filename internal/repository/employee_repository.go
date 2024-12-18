package repository

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type EmployeeRepository struct {
	db map[int]model.Employee
}

func CreateEmployeeRepository(db map[int]model.Employee) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (e *EmployeeRepository) Get() (map[int]model.Employee, error) {
	return e.db, nil
}
