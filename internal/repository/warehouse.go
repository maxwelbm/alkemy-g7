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

func (r *WareHouseMap) DeleteByIdWareHouse(id int) error {
	_, ok := r.db.TbWarehouses[id]

	if !ok {
		return &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	delete(r.db.TbWarehouses, id)
	return nil
}

func (r *WareHouseMap) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {
	warehouse.Id = (len(r.db.TbWarehouses) + 1)

	for _, value := range r.db.TbWarehouses {
		if value.WareHouseCode == warehouse.WareHouseCode {
			return model.WareHouse{}, &custom_error.CustomError{Object: warehouse.WareHouseCode, Err: custom_error.AlreadyExists}
		}
	}
	r.db.TbWarehouses[warehouse.Id] = warehouse
	w = r.db.TbWarehouses[warehouse.Id]
	return
}

func (r *WareHouseMap) Update(id int, warehouse model.WareHouse) (w model.WareHouse, err error) {
	panic("unimplemented")
}
