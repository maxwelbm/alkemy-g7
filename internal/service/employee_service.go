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

func (e *EmployeeService) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	existingEmployee, err := e.GetEmployeeById(id)

	if err != nil {
		return model.Employee{}, err
	}

	updateEmployeeFields(&existingEmployee, employee)

	return e.UpdateEmployee(id, existingEmployee)
}

func updateEmployeeFields(existing *model.Employee, updates model.Employee) {
	if updates.CardNumberId != "" {
		existing.CardNumberId = updates.CardNumberId
	}
	if updates.FirstName != "" {
		existing.FirstName = updates.FirstName
	}
	if updates.LastName != "" {
		existing.LastName = updates.LastName
	}
	if updates.WarehouseId != 0 {
		existing.WarehouseId = updates.WarehouseId
	}
}
