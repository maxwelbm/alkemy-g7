package database

import (
	"encoding/json"
	"os"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type Database struct {
	TbBuyer    map[int]model.Buyer
	TbProducts map[int]model.Buyer
}

func CreateDatabase() *Database {
	db := &Database{}

	db.LoadJsonBuyer("cmd/database/docs/buyers.json")

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
