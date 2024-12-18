package database

import (
	"encoding/json"
	"os"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type Database struct {
	TbBuyer     map[int]model.Buyer
	TbProducts  map[int]model.Buyer
	TbEmployees map[int]model.Employee
}

func CreateDatabase() *Database {
	db := &Database{}

	// db.LoadJsonBuyer("cmd/database/docs/buyers.json")
	db.LoadJsonEmployee("pkg/database/docs/employees.json")
	return db
}

func (db *Database) LoadJsonBuyer(filepath string) (string, error) {
	var buyers []model.Buyer = make([]model.Buyer, 0)

	Load(filepath, buyers)

	for _, b := range buyers {
		db.TbBuyer[b.Id] = b
	}

	return "Succes", nil
}

func (db *Database) LoadJsonEmployee(filepath string) (string, error) {
	var employees []handler.EmployeeJSON = make([]handler.EmployeeJSON, 0)

	Load(filepath, &employees)

	db.TbEmployees = make(map[int]model.Employee)

	for _, employee := range employees {
		db.TbEmployees[employee.Id] = model.Employee{
			Id:           employee.Id,
			CardNumberId: employee.CardNumberId,
			FirstName:    employee.FirstName,
			LastName:     employee.FirstName,
			WarehouseId:  employee.WarehouseId,
		}
	}

	return "Success", nil
}

func Load(filepath string, entity any) {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&entity)

	if err != nil {
		panic(err)
	}
}
