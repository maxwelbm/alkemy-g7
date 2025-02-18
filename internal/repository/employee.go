package repository

import (
	"database/sql"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type EmployeeRepository struct {
	db  *sql.DB
	log logger.Logger
}

func CreateEmployeeRepository(db *sql.DB, log logger.Logger) *EmployeeRepository {
	return &EmployeeRepository{db: db, log: log}
}

func (e *EmployeeRepository) Get() ([]model.Employee, error) {
	e.log.Log("EmployeeRepository", "INFO", "initializing Get function")

	rows, err := e.db.Query("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees")

	if err != nil {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to query employees: %v", err))
		return nil, err
	}

	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {
		var employee model.Employee

		err := rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
		if err != nil {
			e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to scan employee row: %v", err))
			return nil, err
		}

		employees = append(employees, employee)
	}

	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("Get function finished successfully, retrieved %d employees", len(employees)))

	return employees, nil
}

func (e *EmployeeRepository) GetByID(id int) (model.Employee, error) {
	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("initializing GetByID function for employee ID: %d", id))

	var employee model.Employee

	row := e.db.QueryRow("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?", id)

	err := row.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
	if err == sql.ErrNoRows {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("employee not found with ID: %d", id))
		return model.Employee{}, customerror.EmployeeErrNotFound
	} else if err != nil {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to scan employee row: %v", err))
		return model.Employee{}, err
	}

	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("GetByID function finished successfully for employee ID: %d", id))

	return employee, nil
}

func (e *EmployeeRepository) Post(employee model.Employee) (model.Employee, error) {
	e.log.Log("EmployeeRepository", "INFO", "initializing Post function")

	result, err := e.db.Exec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)",
		employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)

	if err != nil {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to insert employee: %v", err))
		return model.Employee{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to retrieve last insert ID: %v", err))
		return model.Employee{}, err
	}

	employee.ID = int(id)
	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("Post function finished successfully, created employee with ID: %d", employee.ID))

	return employee, nil
}

func (e *EmployeeRepository) Update(id int, employee model.Employee) (model.Employee, error) {
	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("initializing Update function for employee ID: %d", id))

	_, err := e.db.Exec("UPDATE employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?",
		employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID, id)
	if err != nil {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to update employee with ID: %d: %v", id, err))
		return model.Employee{}, err
	}

	employee.ID = id
	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("Update function finished successfully for employee ID: %d", id))

	return employee, nil
}

func (e *EmployeeRepository) Delete(id int) error {
	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("initializing Delete function for employee ID: %d", id))

	_, err := e.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to delete employee with ID: %d: %v", id, err))
		return err
	}

	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("Delete function finished successfully for employee ID: %d", id))

	return nil
}

func (e *EmployeeRepository) GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error) {
	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("initializing GetInboundOrdersReportByEmployee function for employee ID: %d", employeeID))

	query := `
		SELECT
			e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(i.id) as inbound_orders_count
		FROM	
			employees e
		LEFT JOIN 
			inbound_orders i 
			ON i.employee_id = e.id
		WHERE e.id = ?
		GROUP BY e.id `

	rows := e.db.QueryRow(query, employeeID)

	var inboundReport model.InboundOrdersReportByEmployee

	err := rows.Scan(&inboundReport.ID, &inboundReport.CardNumberID, &inboundReport.FirstName, &inboundReport.LastName, &inboundReport.WarehouseID, &inboundReport.InboundOrdersCount)

	if err != nil {
		if err == sql.ErrNoRows {
			e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("no inbound orders found for employee ID: %d", employeeID))

			err = customerror.EmployeeErrNotFoundInboundOrders
		} else {
			e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to scan inbound orders report: %v", err))
		}

		return model.InboundOrdersReportByEmployee{}, err
	}

	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("GetInboundOrdersReportByEmployee function finished successfully for employee ID: %d", employeeID))

	return inboundReport, nil
}

func (e *EmployeeRepository) GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error) {
	e.log.Log("EmployeeRepository", "INFO", "initializing GetInboundOrdersReports function")

	query := `
		SELECT
			e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(i.id) as inbound_orders_count
		FROM	
			employees e
		LEFT JOIN 
			inbound_orders i 
			ON i.employee_id = e.id
		GROUP BY e.id `

	rows, err := e.db.Query(query)
	if err != nil {
		e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to query inbound orders reports: %v", err))
		return nil, err
	}

	defer rows.Close()

	var inboundReports []model.InboundOrdersReportByEmployee

	for rows.Next() {
		var inboundReport model.InboundOrdersReportByEmployee

		err = rows.Scan(&inboundReport.ID, &inboundReport.CardNumberID, &inboundReport.FirstName, &inboundReport.LastName, &inboundReport.WarehouseID, &inboundReport.InboundOrdersCount)
		if err != nil {
			e.log.Log("EmployeeRepository", "ERROR", fmt.Sprintf("failed to scan inbound orders report row: %v", err))
			return nil, err
		}

		inboundReports = append(inboundReports, inboundReport)
	}

	e.log.Log("EmployeeRepository", "INFO", fmt.Sprintf("GetInboundOrdersReports function finished successfully, retrieved %d reports", len(inboundReports)))

	return inboundReports, nil
}
