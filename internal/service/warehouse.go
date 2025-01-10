package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type WareHouseDefault struct {
	rp interfaces.IWarehouseRepo
}

func NewWareHoureService(rp interfaces.IWarehouseRepo) *WareHouseDefault {
	return &WareHouseDefault{rp: rp}
}

func (wp *WareHouseDefault) DeleteByIdWareHouse(id int) error {
	panic("unimplemented")
}

func (wp *WareHouseDefault) GetAllWareHouse() (w []model.WareHouse, err error) {
	w, err = wp.rp.GetAllWareHouse()
	return
}

func (wp *WareHouseDefault) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	w, err = wp.rp.GetByIdWareHouse(id)
	return
}

func (wp *WareHouseDefault) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {

	// if warehouse.warehou_code {
	//  return custom_error.NewWareHouseError(custom_error.Conflict.Error(), "WareHouse", http.StatusConflict)
	// }

	id, err := wp.rp.PostWareHouse(&warehouse)

	if err != nil {
		return
	}

	w, err = wp.GetByIdWareHouse(int(id))

	return
}

func (wp *WareHouseDefault) UpdateWareHouse(id int, warehouse model.WareHouse) (w model.WareHouse, err error) {
	warehouseExisting, err := wp.GetByIdWareHouse(id)

	if err != nil {
		return
	}

	if warehouse.WareHouseCode != "" {
		warehouseExisting.WareHouseCode = warehouse.WareHouseCode
	}

	if warehouse.Address != "" {
		warehouseExisting.Address = warehouse.Address
	}

	if warehouse.Telephone != "" {
		warehouseExisting.Telephone = warehouse.Telephone
	}

	if warehouse.MinimunCapacity != 0 {
		warehouseExisting.MinimunCapacity = warehouse.MinimunCapacity
	}

	if warehouse.MinimunTemperature != 0 {
		warehouseExisting.MinimunTemperature = warehouse.MinimunTemperature
	}

	err = wp.rp.UpdateWareHouse(id, &warehouseExisting)

	if err != nil {
		return
	}

	w, err = wp.GetByIdWareHouse(id)

	return
}
