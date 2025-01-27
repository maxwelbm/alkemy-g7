package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type WareHouseDefault struct {
	Rp interfaces.IWarehouseRepo
}

func NewWareHouseService(rp interfaces.IWarehouseRepo) *WareHouseDefault {
	return &WareHouseDefault{Rp: rp}
}

func (wp *WareHouseDefault) DeleteByIdWareHouse(id int) error {
	_, err := wp.GetByIdWareHouse(id)

	if err != nil {
		return err
	}

	err = wp.Rp.DeleteByIdWareHouse(id)

	if err != nil {
		return err
	}

	return nil
}

func (wp *WareHouseDefault) GetAllWareHouse() (w []model.WareHouse, err error) {
	w, err = wp.Rp.GetAllWareHouse()
	return
}

func (wp *WareHouseDefault) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	w, err = wp.Rp.GetByIdWareHouse(id)
	return
}

func (wp *WareHouseDefault) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {

	id, err := wp.Rp.PostWareHouse(warehouse)

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

	err = wp.Rp.UpdateWareHouse(id, warehouseExisting)

	if err != nil {
		return
	}

	w, err = wp.GetByIdWareHouse(id)

	return
}
