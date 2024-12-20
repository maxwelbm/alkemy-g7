package service

import (
	"errors"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type EmployeeService struct {
	rp    interfaces.IEmployeeRepo
	wrSrv interfaces.IWarehouseRepo
}

func CreateEmployeeService(rp interfaces.IEmployeeRepo, wrSrv interfaces.IWarehouseRepo) *EmployeeService {
	return &EmployeeService{rp: rp, wrSrv: wrSrv}
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

func (e *EmployeeService) InsertEmployee(employee model.Employee) (model.Employee, error) {
	if !employee.IsValidEmployee() {
		return model.Employee{}, custom_error.CustomError{Object: employee, Err: custom_error.InvalidErr}
	}

	_, err := e.wrSrv.GetByIdWareHouse(employee.WarehouseId)

	if err != nil {
		return model.Employee{}, custom_error.CustomError{Object: employee, Err: errors.New("WarehouseID dont exist")}
	}

	employee.Id = e.generateNewId()
	return e.rp.Post(employee)
}

func (e *EmployeeService) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	existingEmployee, err := e.GetEmployeeById(id)

	if err != nil {
		return model.Employee{}, err
	}

	updateEmployeeFields(&existingEmployee, employee)

	return e.rp.Update(id, existingEmployee)
}

func (e *EmployeeService) DeleteEmployee(id int) error {
	_, err := e.rp.GetById(id)

	if err != nil {
		return err
	}

	return e.rp.Delete(id)
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

func (e *EmployeeService) generateNewId() int {
	lastId := 0
	data, _ := e.rp.Get()

	for _, employee := range data {
		if employee.Id > lastId {
			lastId = employee.Id
		}
	}

	return lastId + 1
}
