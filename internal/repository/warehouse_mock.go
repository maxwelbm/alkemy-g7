package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type WareHouseMockRepo struct {
	mock.Mock
}

func (mock *WareHouseMockRepo) DeleteByIdWareHouse(id int) error {
	args := mock.Called(id)

	return args.Error(0)
}

func (mock *WareHouseMockRepo) GetAllWareHouse() (w []model.WareHouse, err error) {
	args := mock.Called()

	w = args.Get(0).([]model.WareHouse)
	err = args.Error(1)

	return
}

func (mock *WareHouseMockRepo) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	args := mock.Called(id)

	w = args.Get(0).(model.WareHouse)
	err = args.Error(1)

	return
}

func (w *WareHouseMockRepo) PostWareHouse(warehouse *model.WareHouse) (id int64, err error) {
	panic("unimplemented")
}

func (w *WareHouseMockRepo) UpdateWareHouse(id int, warehouse *model.WareHouse) (err error) {
	panic("unimplemented")
}
