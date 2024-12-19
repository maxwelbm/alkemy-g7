package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

type WareHouseMap struct {
	db database.Database
}

func NewWareHouseRepository(db database.Database) *WareHouseMap {
	return &WareHouseMap{db: db}
}

func (r *WareHouseMap) GetAllWareHouse() (w map[int]model.WareHouse, err error) {
	w = r.db.TbWarehouses
	return
}

func (r *WareHouseMap) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	w, ok := r.db.TbWarehouses[id]

	if !ok {
		return model.WareHouse{}, &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}
	return
}

func (r *WareHouseMap) Delete(id int) error {
	panic("unimplemented")
}

func (r *WareHouseMap) Post(warehouse model.WareHouse) (model.WareHouse, error) {
	panic("unimplemented")
}

func (r *WareHouseMap) Update(id int, warehouse model.WareHouse) (model.WareHouse, error) {
	panic("unimplemented")
}