package repository

import (
	"database/sql"
	"errors"

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
		return model.Employee{}, custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	} else if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *EmployeeRepository) Post(employee model.Employee) (model.Employee, error) {
	_, err := e.getEmployeeByCardNumber(employee.CardNumberId)
	if err == nil {
		return model.Employee{}, custom_error.CustomError{Object: employee, Err: errors.New("duplicated CardNumberId")}
	}

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
		return custom_error.CustomError{Object: id, Err: err}
	}
	return nil
}

func (e *EmployeeRepository) getEmployeeByCardNumber(cardNumberId string) (model.Employee, error) {
	var employee model.Employee
	row := e.db.QueryRow("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE card_number_id = ?", cardNumberId)
	err := row.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId)
	if err == sql.ErrNoRows {
		return model.Employee{}, custom_error.CustomError{Object: cardNumberId, Err: custom_error.Conflict}
	} else if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}
