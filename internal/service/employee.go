package service

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type EmployeeService struct {
	rp    interfaces.IEmployeeRepo
	wrSrv interfaces.IWarehouseRepo
	log   logger.Logger
}

func CreateEmployeeService(rp interfaces.IEmployeeRepo, wrSrv interfaces.IWarehouseRepo, log logger.Logger) *EmployeeService {
	return &EmployeeService{rp: rp, wrSrv: wrSrv, log: log}
}

func (e *EmployeeService) GetEmployees() ([]model.Employee, error) {
	e.log.Log("EmployeeService", "INFO", "Fetching all employees")
	data, err := e.rp.Get()

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Failed to fetch employees: %v", err))
		return nil, err
	}

	e.log.Log("EmployeeService", "INFO", "Successfully fetched employees")

	return data, nil
}

func (e *EmployeeService) GetEmployeeByID(id int) (model.Employee, error) {
	e.log.Log("EmployeeService", "INFO", fmt.Sprintf("Fetching employee with ID %d", id))
	data, err := e.rp.GetByID(id)

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Failed to fetch employee %d: %v", id, err))
	}

	return data, err
}

func (e *EmployeeService) InsertEmployee(employee model.Employee) (model.Employee, error) {
	e.log.Log("EmployeeService", "INFO", "Inserting new employee")

	if !employee.IsValidEmployee() {
		return model.Employee{}, customerror.EmployeeErrInvalid
	}

	_, err := e.wrSrv.GetByIDWareHouse(employee.WarehouseID)

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", "Invalid warehouse ID")
		return model.Employee{}, customerror.EmployeeErrInvalidWarehouseID
	}

	employee, err = e.rp.Post(employee)

	if err != nil {
		if mySQLErr, ok := err.(*mysql.MySQLError); ok && mySQLErr.Number == 1062 {
			e.log.Log("EmployeeService", "ERROR", "Duplicate card number")
			return model.Employee{}, customerror.EmployeeErrDuplicatedCardNumber
		}

		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Failed to insert employee: %v", err))

		return model.Employee{}, err
	}

	e.log.Log("EmployeeService", "INFO", "Employee inserted successfully")

	return employee, nil
}

func (e *EmployeeService) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	e.log.Log("EmployeeService", "INFO", fmt.Sprintf("Updating employee with ID %d", id))

	if employee.IsEmptyEmployee() {
		return model.Employee{}, customerror.EmployeeErrInvalid
	}

	if employee.WarehouseID != 0 {
		_, err := e.wrSrv.GetByIDWareHouse(employee.WarehouseID)

		if err != nil {
			e.log.Log("EmployeeService", "ERROR", "Invalid warehouse ID for update")
			return model.Employee{}, customerror.EmployeeErrInvalidWarehouseID
		}
	}

	existingEmployee, err := e.rp.GetByID(id)

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Employee %d not found: %v", id, err))
		return model.Employee{}, err
	}

	updateEmployeeFields(&existingEmployee, employee)
	updatedEmployee, err := e.rp.Update(id, existingEmployee)

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Failed to update employee %d: %v", id, err))

		return model.Employee{}, err
	}

	e.log.Log("EmployeeService", "INFO", fmt.Sprintf("Employee %d updated successfully", id))

	return updatedEmployee, nil
}

func (e *EmployeeService) DeleteEmployee(id int) error {
	e.log.Log("EmployeeService", "INFO", fmt.Sprintf("Deleting employee with ID %d", id))
	_, err := e.rp.GetByID(id)

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Failed to find employee %d: %v", id, err))
		return err
	}

	err = e.rp.Delete(id)
	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Failed to delete employee %d: %v", id, err))
		return err
	}

	e.log.Log("EmployeeService", "INFO", fmt.Sprintf("Employee %d deleted successfully", id))

	return nil
}

func (e *EmployeeService) GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error) {
	e.log.Log("EmployeeService", "INFO", fmt.Sprintf("Fetching inbound orders report for employee %d", employeeID))

	if employeeID <= 0 {
		return model.InboundOrdersReportByEmployee{}, customerror.EmployeeErrInvalid
	}

	_, err := e.rp.GetByID(employeeID)

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Employee %d not found: %v", employeeID, err))
		return model.InboundOrdersReportByEmployee{}, err
	}

	data, err := e.rp.GetInboundOrdersReportByEmployee(employeeID)

	if err != nil {
		e.log.Log("EmployeeService", "ERROR", fmt.Sprintf("Failed to fetch report for employee %d: %v", employeeID, err))
		return model.InboundOrdersReportByEmployee{}, err
	}

	e.log.Log("EmployeeService", "INFO", fmt.Sprintf("Fetched report for employee %d successfully", employeeID))

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
