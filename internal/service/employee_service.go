package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type EmployeeService struct {
	rp    interfaces.IEmployeeRepo
	wrSrv interfaces.IWarehouseRepo
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

func (e *EmployeeService) InsertEmployee(employee model.Employee) (model.Employee, error) {
	if !employee.IsValidEmployee() {
		return model.Employee{}, custom_error.CustomError{Object: employee, Err: custom_error.InvalidErr}
	}

	// _, err = e.wrSrv.GetById(employee.WarehouseId)

	//@todo validate error throwed by warehouseService
	// if err != nil {
	// 	return model.Employee{}, custom_error.CustomError{Object: employee, Err: custom_error.InvalidErr}
	// }

	employee.Id = e.generateNewId()
	return e.rp.Post(employee)
}

func (e *EmployeeService) DeleteEmployee(id int) error {
	_, err := e.rp.GetById(id)

	if err != nil {
		return err
	}

	return e.rp.Delete(id)
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
