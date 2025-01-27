package service

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type EmployeeService struct {
	rp    interfaces.IEmployeeRepo
	wrSrv interfaces.IWarehouseRepo
}

func CreateEmployeeService(rp interfaces.IEmployeeRepo, wrSrv interfaces.IWarehouseRepo) *EmployeeService {
	return &EmployeeService{rp: rp, wrSrv: wrSrv}
}

func (e *EmployeeService) GetEmployees() ([]model.Employee, error) {
	data, err := e.rp.Get()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

func (e *EmployeeService) GetEmployeeByID(id int) (model.Employee, error) {
	return e.rp.GetByID(id)
}

func (e *EmployeeService) InsertEmployee(employee model.Employee) (model.Employee, error) {
	if !employee.IsValidEmployee() {
		return model.Employee{}, customerror.EmployeeErrInvalid
	}

	_, err := e.wrSrv.GetByIDWareHouse(employee.WarehouseID)

	if err != nil {
		return model.Employee{}, customerror.EmployeeErrInvalidWarehouseID
	}

	employee, err = e.rp.Post(employee)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.EmployeeErrDuplicatedCardNumber
		}

		return model.Employee{}, err
	}

	return employee, nil
}

func (e *EmployeeService) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	if employee.IsEmptyEmployee() {
		return model.Employee{}, customerror.EmployeeErrInvalid
	}

	if employee.WarehouseID != 0 {
		_, err := e.wrSrv.GetByIDWareHouse(employee.WarehouseID)

		if err != nil {
			return model.Employee{}, customerror.EmployeeErrInvalidWarehouseID
		}
	}

	existingEmployee, err := e.rp.GetByID(id)

	if err != nil {
		return model.Employee{}, err
	}

	updateEmployeeFields(&existingEmployee, employee)

	return e.rp.Update(id, existingEmployee)
}

func (e *EmployeeService) DeleteEmployee(id int) error {
	_, err := e.rp.GetByID(id)

	if err != nil {
		return err
	}

	return e.rp.Delete(id)
}

func (e *EmployeeService) GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error) {
	if employeeID <= 0 {
		return model.InboundOrdersReportByEmployee{}, customerror.EmployeeErrInvalid
	}

	_, err := e.rp.GetByID(employeeID)

	if err != nil {
		return model.InboundOrdersReportByEmployee{}, err
	}

	data, err := e.rp.GetInboundOrdersReportByEmployee(employeeID)

	if err != nil {
		return model.InboundOrdersReportByEmployee{}, err
	}

	return data, nil
}

func (e *EmployeeService) GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error) {
	return e.rp.GetInboundOrdersReports()
}

func updateEmployeeFields(existing *model.Employee, updates model.Employee) {
	if updates.CardNumberID != "" {
		existing.CardNumberID = updates.CardNumberID
	}

	if updates.FirstName != "" {
		existing.FirstName = updates.FirstName
	}

	if updates.LastName != "" {
		existing.LastName = updates.LastName
	}

	if updates.WarehouseID != 0 {
		existing.WarehouseID = updates.WarehouseID
	}
}
