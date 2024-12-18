package database

import (
	"encoding/json"
	"os"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type Database struct {
	TbBuyer    map[int]model.Buyer
	TbSellers  map[int]model.Seller
}

func CreateDatabase() *Database {
	db := &Database{
		TbBuyer: make(map[int]model.Buyer),
		TbSellers: make(map[int]model.Seller),
	}

	//db.LoadJsonBuyer("/workspaces/alkemy-g7/cmd/database/docs/buyers.json")
	db.LoadJsonSeller("/Users/julidsilva/Documents/bootcamp/linguagem_go/sprint01/alkemy-g7/cmd/database/docs/sellers.json")

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

func (db *Database) LoadJsonSeller(filepath string) (string, error) {
	var sellers []model.Seller = make([]model.Seller, 0)

	Load(filepath, &sellers)

	for _, seller := range sellers {
		db.TbSellers[seller.ID] = seller
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