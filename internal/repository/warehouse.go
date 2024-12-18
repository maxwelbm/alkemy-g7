package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
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
