package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type WareHouseDefault struct {
	rp interfaces.IWarehouseRepo
}

func NewWareHoureService(rp interfaces.IWarehouseRepo) *WareHouseDefault {
	return &WareHouseDefault{rp: rp}
}

func (s *WareHouseDefault) DeleteByIdWareHouse(id int) error {
	return s.rp.DeleteByIdWareHouse(id)
}

func (s *WareHouseDefault) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {
	return s.rp.PostWareHouse(warehouse)
}

func (s *WareHouseDefault) UpdateWareHouse(id int, warehouse model.WareHouse) (w model.WareHouse, err error) {
	wareHouseById, err := s.rp.GetByIdWareHouse(id)

	if err != nil {
		return model.WareHouse{}, &custom_error.CustomError{Object: "empty body", Err: custom_error.NotFound}
	}

	if warehouse.Id == 0 && warehouse.Address == "" && warehouse.WareHouseCode == "" && warehouse.Telephone == "" && warehouse.MinimunCapacity == 0 && warehouse.MinimunTemperature == 0 {
		return model.WareHouse{}, custom_error.CustomError{Object: warehouse, Err: custom_error.NotFound}
	}

	if warehouse.Address != "" {
		wareHouseById.Address = warehouse.Address
	}
	if warehouse.WareHouseCode != "" {
		wareHouseById.WareHouseCode = warehouse.WareHouseCode
	}
	if warehouse.Telephone != "" {
		wareHouseById.Telephone = warehouse.Telephone
	}
	if warehouse.MinimunCapacity != 0 {
		wareHouseById.MinimunCapacity = warehouse.MinimunCapacity
	}
	if warehouse.MinimunTemperature != 0 {
		wareHouseById.MinimunTemperature = warehouse.MinimunTemperature
	}

	return s.rp.UpdateWareHouse(id, wareHouseById)
}

func (s *WareHouseDefault) GetAllWareHouse() (w map[int]model.WareHouse, err error) {
	w, err = s.rp.GetAllWareHouse()
	return
}

func (s *WareHouseDefault) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	w, err = s.rp.GetByIdWareHouse(id)
	return
}
