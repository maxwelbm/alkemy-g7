package service

import (
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type WareHouseDefault struct {
	Rp  interfaces.IWarehouseRepo
	log logger.Logger
}

func NewWareHouseService(rp interfaces.IWarehouseRepo, log logger.Logger) *WareHouseDefault {
	return &WareHouseDefault{Rp: rp, log: log}
}

func (wp *WareHouseDefault) DeleteByIDWareHouse(id int) error {
	wp.log.Log("WareHouseService", "INFO", "initializing DeleteByIDWareHouse function")

	_, err := wp.GetByIDWareHouse(id)

	if err != nil {
		wp.log.Log("WareHouseService", "ERROR", fmt.Sprintf("Error: %v", err))

		return err
	}

	err = wp.Rp.DeleteByIDWareHouse(id)

	if err != nil {
		wp.log.Log("WareHouseService", "ERROR", fmt.Sprintf("Error: %v", err))

		return err
	}
	wp.log.Log("WareHouseService", "INFO", "DeleteByIDWareHouse completed successfully")

	return nil
}

func (wp *WareHouseDefault) GetAllWareHouse() (w []model.WareHouse, err error) {
	wp.log.Log("WareHouseService", "INFO", "initializing GetAllWareHouse function")

	w, err = wp.Rp.GetAllWareHouse()

	wp.log.Log("WareHouseService", "INFO", "GetAllWareHouse completed successfully")
	return
}

func (wp *WareHouseDefault) GetByIDWareHouse(id int) (w model.WareHouse, err error) {
	wp.log.Log("WareHouseService", "INFO", "initializing GetByIDWareHouse function")

	w, err = wp.Rp.GetByIDWareHouse(id)

	wp.log.Log("WareHouseService", "INFO", "GetByIDWareHouse completed successfully")
	return
}

func (wp *WareHouseDefault) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {
	wp.log.Log("WareHouseService", "INFO", "initializing PostWareHouse function")

	id, err := wp.Rp.PostWareHouse(warehouse)

	if err != nil {
		wp.log.Log("WareHouseService", "ERROR", fmt.Sprintf("Error: %v", err))

		return w, err
	}

	w, err = wp.GetByIDWareHouse(int(id))

	wp.log.Log("WareHouseService", "INFO", "PostWareHouse completed successfully")
	return w, err
}

func (wp *WareHouseDefault) UpdateWareHouse(id int, warehouse model.WareHouse) (w model.WareHouse, err error) {
	wp.log.Log("WareHouseService", "INFO", "initializing UpdateWareHouse function")

	warehouseExisting, err := wp.GetByIDWareHouse(id)

	if err != nil {
		wp.log.Log("WareHouseService", "ERROR", fmt.Sprintf("Error: %v", err))

		return w, err
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
		wp.log.Log("WareHouseService", "ERROR", fmt.Sprintf("Error: %v", err))
		return w, err
	}

	w, err = wp.GetByIDWareHouse(id)

	wp.log.Log("WareHouseService", "INFO", "UpdateWareHouse completed successfully")
	return w, err
}
