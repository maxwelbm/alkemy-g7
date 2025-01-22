package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type WareHouseMockRepo struct {
	mock.Mock
}

// DeleteByIdWareHouse implements interfaces.IWarehouseRepo.
func (w *WareHouseMockRepo) DeleteByIdWareHouse(id int) error {
	panic("unimplemented")
}

// GetAllWareHouse implements interfaces.IWarehouseRepo.
func (mock *WareHouseMockRepo) GetAllWareHouse() (w []model.WareHouse, err error) {
	args := mock.Called()

	w = args.Get(0).([]model.WareHouse)
	err = args.Error(1)

	return
}

// GetByIdWareHouse implements interfaces.IWarehouseRepo.
func (*WareHouseMockRepo) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	panic("unimplemented")
}

// PostWareHouse implements interfaces.IWarehouseRepo.
func (w *WareHouseMockRepo) PostWareHouse(warehouse *model.WareHouse) (id int64, err error) {
	panic("unimplemented")
}

// UpdateWareHouse implements interfaces.IWarehouseRepo.
func (w *WareHouseMockRepo) UpdateWareHouse(id int, warehouse *model.WareHouse) (err error) {
	panic("unimplemented")
}
