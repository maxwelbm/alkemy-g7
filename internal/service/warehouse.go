package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type WareHouseDefault struct {
	rp interfaces.IWarehouseRepo
}

// Delete implements interfaces.IWarehouseService.
func (s *WareHouseDefault) Delete(id int) error {
	panic("unimplemented")
}

// Post implements interfaces.IWarehouseService.
func (s *WareHouseDefault) Post(warehouse model.WareHouse) (w model.WareHouse, err error) {
	panic("unimplemented")
}

// Update implements interfaces.IWarehouseService.
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
