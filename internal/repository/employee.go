package repository

import (
	"database/sql"
	"log"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

type EmployeeRepository struct {
	db *sql.DB
}

func CreateEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (e *EmployeeRepository) Get() ([]model.Employee, error) {

	rows, err := e.db.Query("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *EmployeeRepository) GetById(id int) (model.Employee, error) {
	var employee model.Employee
	row := e.db.QueryRow("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?", id)
	err := row.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId)
	if err == sql.ErrNoRows {
		return model.Employee{}, customError.EmployeeErrNotFound
	} else if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *EmployeeRepository) Post(employee model.Employee) (model.Employee, error) {

	result, err := e.db.Exec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)",
		employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)

	if err != nil {
		return model.Employee{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.Employee{}, err
	}
	employee.Id = int(id)
	return employee, nil
}

func (e *EmployeeRepository) Update(id int, employee model.Employee) (model.Employee, error) {
	_, err := e.db.Exec("UPDATE employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?",
		employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId, id)
	if err != nil {
		return model.Employee{}, err
	}
	employee.Id = id
	return employee, nil
}

func (e *EmployeeRepository) Delete(id int) error {
	_, err := e.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmployeeRepository) GetInboundOrdersReportByEmployee(employeeId int) (model.InboundOrdersReportByEmployee, error) {
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

	rows := e.db.QueryRow(query, employeeId)

	var inboundReport model.InboundOrdersReportByEmployee

	err := rows.Scan(&inboundReport.Id, &inboundReport.CardNumberId, &inboundReport.FirstName, &inboundReport.LastName, &inboundReport.WarehouseId, &inboundReport.InboundOrdersCount)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customError.EmployeeErrNotFoundInboundOrders
		}
		return model.InboundOrdersReportByEmployee{}, err
	}

	return inboundReport, nil
}

func (e *EmployeeRepository) GetInboundOrdersReports() ([]model.InboundOrdersReportByEmployee, error) {
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
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var inboundReports []model.InboundOrdersReportByEmployee
	for rows.Next() {
		var inboundReport model.InboundOrdersReportByEmployee
		err = rows.Scan(&inboundReport.Id, &inboundReport.CardNumberId, &inboundReport.FirstName, &inboundReport.LastName, &inboundReport.WarehouseId, &inboundReport.InboundOrdersCount)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		inboundReports = append(inboundReports, inboundReport)
	}

	return inboundReports, nil
}
