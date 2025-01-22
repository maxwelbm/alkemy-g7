package repository

import (
	"database/sql"
	"log"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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

		err := rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (e *EmployeeRepository) GetByID(id int) (model.Employee, error) {
	var employee model.Employee

	row := e.db.QueryRow("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?", id)

	err := row.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
	if err == sql.ErrNoRows {
		return model.Employee{}, custom_error.EmployeeErrNotFound
	} else if err != nil {
		return model.Employee{}, err
	}

	return employee, nil
}

func (e *EmployeeRepository) Post(employee model.Employee) (model.Employee, error) {
	result, err := e.db.Exec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)",
		employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)

	if err != nil {
		return model.Employee{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.Employee{}, err
	}

	employee.ID = int(id)

	return employee, nil
}

func (e *EmployeeRepository) Update(id int, employee model.Employee) (model.Employee, error) {
	_, err := e.db.Exec("UPDATE employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?",
		employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID, id)
	if err != nil {
		return model.Employee{}, err
	}

	employee.ID = id

	return employee, nil
}

func (e *EmployeeRepository) Delete(id int) error {
	_, err := e.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepository) GetInboundOrdersReportByEmployee(employeeID int) (model.InboundOrdersReportByEmployee, error) {
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
			err = custom_error.EmployeeErrNotFoundInboundOrders
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

		err = rows.Scan(&inboundReport.ID, &inboundReport.CardNumberID, &inboundReport.FirstName, &inboundReport.LastName, &inboundReport.WarehouseID, &inboundReport.InboundOrdersCount)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		inboundReports = append(inboundReports, inboundReport)
	}

	return inboundReports, nil
}
