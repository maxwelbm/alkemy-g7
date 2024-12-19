package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type WareHouseDefault struct {
	rp interfaces.IWarehouseRepo
}

func (s *WareHouseDefault) DeleteByIdWareHouse(id int) error {
	return s.rp.DeleteByIdWareHouse(id)
}

func (s *WareHouseDefault) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {
	return s.rp.PostWareHouse(warehouse)
}

func (s *WareHouseDefault) Update(id int, warehouse model.WareHouse) (w model.WareHouse, err error) {
	panic("unimplemented")
}

func NewWareHoureService(rp interfaces.IWarehouseRepo) *WareHouseDefault {
	return &WareHouseDefault{rp: rp}
}

func (s *WareHouseDefault) GetAllWareHouse() (w map[int]model.WareHouse, err error) {
	w, err = s.rp.GetAllWareHouse()
	return
}

func (s *WareHouseDefault) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	w, err = s.rp.GetByIdWareHouse(id)
	return
}
