package database

import (
	"encoding/json"
	"os"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type Database struct {
	TbBuyer     map[int]model.Buyer
	TbProducts  map[int]model.Buyer
	TbWarehouse map[int]model.WareHouse
}

func CreateDatabase() *Database {
	db := &Database{
		TbBuyer:     make(map[int]model.Buyer),
		TbWarehouse: make(map[int]model.WareHouse),
	}

	db.LoadJsonBuyer("/workspaces/alkemy-g7/cmd/database/docs/buyers.json")
	db.LoadJsonWarehouse("cmd/database/docs/warehouse.json")

	return db
}

func (db *Database) LoadJsonBuyer(filepath string) (string, error) {
	var buyers []model.Buyer = make([]model.Buyer, 0)

	Load(filepath, &buyers)

	for _, b := range buyers {
		db.TbBuyer[b.Id] = b
	}

	return "Succes", nil
}

func (db *Database) LoadJsonWarehouse(filepath string) (string, error) {
	var warehouse []model.WareHouse = make([]model.WareHouse, 0)

	Load(filepath, &warehouse)

	for _, b := range warehouse {
		db.TbWarehouse[b.Id] = b
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
