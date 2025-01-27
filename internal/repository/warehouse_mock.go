package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type WareHouseMockRepo struct {
	mock.Mock
}

func (mr *WareHouseMockRepo) DeleteByIDWareHouse(id int) error {
	args := mr.Called(id)

	return args.Error(0)
}

func (mr *WareHouseMockRepo) GetAllWareHouse() (w []model.WareHouse, err error) {
	args := mr.Called()

	w = args.Get(0).([]model.WareHouse)
	err = args.Error(1)

	return
}

func (mr *WareHouseMockRepo) GetByIDWareHouse(id int) (w model.WareHouse, err error) {
	args := mr.Called(id)

	w = args.Get(0).(model.WareHouse)
	err = args.Error(1)

	return
}

func (mr *WareHouseMockRepo) PostWareHouse(warehouse model.WareHouse) (id int64, err error) {
	args := mr.Called(warehouse)

	id = args.Get(0).(int64)
	err = args.Error(1)

	return
}

func (mr *WareHouseMockRepo) UpdateWareHouse(id int, warehouse model.WareHouse) (err error) {
	args := mr.Called(id, warehouse)

	err = args.Error(0)

	return
}
