package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/cmd/database"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type WareHouseMap struct {
	db database.Database
}

func NewWareHouseMap(db database.Database) *WareHouseMap {
	return &WareHouseMap{db: db}
}

func (r *WareHouseMap) GetAllWareHouse() (w map[int]model.WareHouse, err error) {
	w = r.db.TbWarehouse
	return
}
