package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	// "github.com/maxwelbm/alkemy-g7.git/internal/handler"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

	db.TbWarehouses = make(map[int]model.WareHouse)

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

type Db struct {
	Connection *sql.DB
}

func NewConnectionDb(db *mysql.Config) (*Db, error) {

	conn, err := sql.Open("mysql", db.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &Db{Connection: conn}, nil
}

func (Db *Db) Close() error {
	return Db.Connection.Close()
}

func GetDbConfig() (*mysql.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading environment variables")
	}

	dbHost := os.Getenv("DB_HOST")

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbNet := os.Getenv("DB_NET")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbNet == "" {

		return nil, errors.New("missing required environment configuration for the database")
	}

	return &mysql.Config{
		User:      dbUser,
		Passwd:    dbPassword,
		Net:       dbNet,
		Addr:      dbHost,
		ParseTime: true,
		DBName:    dbName,
	}, nil

}
