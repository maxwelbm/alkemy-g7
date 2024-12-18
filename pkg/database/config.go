package database

import (
	"encoding/json"
	"os"

	// "github.com/maxwelbm/alkemy-g7.git/internal/handler"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type Database struct {
	TbBuyer      map[int]model.Buyer
	TbProducts   map[int]model.Product
	TbSections   map[int]model.Section
	TbEmployees  map[int]model.Employee
	TbSellers    map[int]model.Seller
	TbWarehouses map[int]model.WareHouse
}

func CreateDatabase() *Database {
	db := &Database{
		TbBuyer:      make(map[int]model.Buyer),
		TbProducts:   make(map[int]model.Product),
		TbSections:   make(map[int]model.Section),
		TbEmployees:  make(map[int]model.Employee),
		TbSellers:    make(map[int]model.Seller),
		TbWarehouses: make(map[int]model.WareHouse),
	}

	db.LoadJsonSellers("pkg/database/docs/sellers.json")
	db.LoadJsonSections("pkg/database/docs/sections.json")
	db.LoadJsonProducts("pkg/database/docs/products.json")
	db.LoadJsonBuyer("pkg/database/docs/buyers.json")
	db.LoadJsonEmployee("pkg/database/docs/employees.json")
	db.LoadJsonWarehouse("pkg/database/docs/warehouse.json")

	return db
}

func (db *Database) LoadJsonProducts(filepath string) (string, error) {
	var products []model.Product = make([]model.Product, 0)

	Load(filepath, &products)

	for _, b := range products {
		db.TbProducts[b.ID] = b
	}

	return "Succes", nil
}

func (db *Database) LoadJsonBuyer(filepath string) (string, error) {
	var buyers []model.Buyer = make([]model.Buyer, 0)

	Load(filepath, &buyers)

	for _, b := range buyers {
		db.TbBuyer[b.Id] = b
	}

	return "Succes", nil
}

func (db *Database) LoadJsonSections(filepath string) (string, error) {
	var sections []model.Section = make([]model.Section, 0)

	Load(filepath, &sections)

	for _, section := range sections {
		db.TbSections[section.ID] = section
	}

	return "Succes", nil
}

func (db *Database) LoadJsonSellers(filepath string) (string, error) {
	var sellers []model.Seller = make([]model.Seller, 0)
	Load(filepath, &sellers)
	for _, seller := range sellers {
		db.TbSellers[seller.ID] = seller
	}
	return "Succes", nil
}

func (db *Database) LoadJsonEmployee(filepath string) (string, error) {
	var employees []model.Employee = make([]model.Employee, 0)

	Load(filepath, &employees)

	db.TbEmployees = make(map[int]model.Employee)

	for _, employee := range employees {
		db.TbEmployees[employee.Id] = employee

	}

	return "Success", nil
}

func (db *Database) LoadJsonWarehouse(filepath string) (string, error) {
	var warehouse []model.WareHouse = make([]model.WareHouse, 0)
	Load(filepath, &warehouse)
	for _, b := range warehouse {
		db.TbWarehouses[b.Id] = b
	}
	return "Succes", nil
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
