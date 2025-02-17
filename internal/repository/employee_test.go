package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

var logMock = mocks.MockLog{}

func TestEmployeeRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := CreateEmployeeRepository(db, logMock)

	t.Run("retrieving all employees", func(t *testing.T) {
		employees := []model.Employee{
			{ID: 1, CardNumberID: "12345", FirstName: "John", LastName: "Doe", WarehouseID: 1},
			{ID: 2, CardNumberID: "67890", FirstName: "Jane", LastName: "Smith", WarehouseID: 2},
		}

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
		for _, emp := range employees {
			rows.AddRow(emp.ID, emp.CardNumberID, emp.FirstName, emp.LastName, emp.WarehouseID)
		}

		mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees").
			WillReturnRows(rows)

		result, err := rp.Get()
		assert.NoError(t, err)
		assert.Equal(t, employees, result)
	})
}

func TestEmployeeRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := CreateEmployeeRepository(db, logMock)

	t.Run("retrieving an existing employee by ID", func(t *testing.T) {
		employeeID := 1
		employee := model.Employee{
			ID:           employeeID,
			CardNumberID: "12345",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
			AddRow(employee.ID, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)

		mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?").
			WithArgs(employeeID).
			WillReturnRows(rows)

		result, err := rp.GetByID(employeeID)
		assert.NoError(t, err)
		assert.Equal(t, employee, result)
	})

	t.Run("employee not found", func(t *testing.T) {
		employeeID := 100

		mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id=?").
			WithArgs(employeeID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))

		_, err := rp.GetByID(employeeID)
		assert.Error(t, err)
	})
}

func TestEmployeeRepository_Post(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := CreateEmployeeRepository(db, logMock)

	t.Run("successful addition of an employee", func(t *testing.T) {
		employee := model.Employee{
			CardNumberID: "12345",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}

		mock.ExpectExec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)").
			WithArgs(employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		result, err := rp.Post(employee)
		assert.NoError(t, err)
		assert.Equal(t, 1, result.ID)
	})
}

func TestEmployeeRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := CreateEmployeeRepository(db, logMock)

	t.Run("successful update of an employee's details", func(t *testing.T) {
		employeeID := 1
		employee := model.Employee{
			CardNumberID: "67890",
			FirstName:    "Jane",
			LastName:     "Doe",
			WarehouseID:  2,
		}

		mock.ExpectExec("UPDATE employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?").
			WithArgs(employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID, employeeID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		result, err := rp.Update(employeeID, employee)
		assert.NoError(t, err)
		assert.Equal(t, employeeID, result.ID)
	})
}

func TestEmployeeRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := CreateEmployeeRepository(db, logMock)

	t.Run("successful deletion of an employee", func(t *testing.T) {
		employeeID := 1

		mock.ExpectExec("DELETE FROM employees WHERE id = ?").
			WithArgs(employeeID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := rp.Delete(employeeID)
		assert.NoError(t, err)
	})
}

func TestEmployeeRepository_GetInboundOrdersReportByEmployee(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := CreateEmployeeRepository(db, logMock)

	t.Run("successful retrieval of inbound orders report by employee", func(t *testing.T) {
		employeeID := 1
		report := model.InboundOrdersReportByEmployee{
			ID:                 employeeID,
			CardNumberID:       "12345",
			FirstName:          "John",
			LastName:           "Doe",
			WarehouseID:        1,
			InboundOrdersCount: 5,
		}

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}).
			AddRow(report.ID, report.CardNumberID, report.FirstName, report.LastName, report.WarehouseID, report.InboundOrdersCount)

		mock.ExpectQuery("SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(i.id) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders i ON i.employee_id = e.id WHERE e.id = ? GROUP BY e.id").
			WithArgs(employeeID).
			WillReturnRows(rows)

		result, err := rp.GetInboundOrdersReportByEmployee(employeeID)
		assert.NoError(t, err)
		assert.Equal(t, report, result)
	})

	t.Run("employee not found", func(t *testing.T) {
		employeeID := 100

		mock.ExpectQuery("SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(i.id) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders i ON i.employee_id = e.id WHERE e.id = ? GROUP BY e.id").
			WithArgs(employeeID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))

		_, err := rp.GetInboundOrdersReportByEmployee(employeeID)
		assert.Error(t, err)
	})
}

func TestEmployeeRepository_GetInboundOrdersReports(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := CreateEmployeeRepository(db, logMock)

	t.Run("successful retrieval of all inbound orders reports", func(t *testing.T) {
		reports := []model.InboundOrdersReportByEmployee{
			{ID: 1, CardNumberID: "12345", FirstName: "John", LastName: "Doe", WarehouseID: 1, InboundOrdersCount: 5},
			{ID: 2, CardNumberID: "67890", FirstName: "Jane", LastName: "Smith", WarehouseID: 2, InboundOrdersCount: 3},
		}

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})
		for _, report := range reports {
			rows.AddRow(report.ID, report.CardNumberID, report.FirstName, report.LastName, report.WarehouseID, report.InboundOrdersCount)
		}

		mock.ExpectQuery("SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(i.id) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders i ON i.employee_id = e.id GROUP BY e.id").
			WillReturnRows(rows)

		result, err := rp.GetInboundOrdersReports()
		assert.NoError(t, err)
		assert.Equal(t, reports, result)
	})
}
