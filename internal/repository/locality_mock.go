package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type LocalityMockRepository struct {
	mock.Mock
}

func (rp *LocalityMockRepository) GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error) {
	args := rp.Called(id)
	report = args.Get(0).([]model.LocalitiesJSONCarriers)
	err = args.Error(1)

	return
}

func (rp *LocalityMockRepository) GetReportCarriersWithId(id int) (locality []model.LocalitiesJSONCarriers, err error) {
	args := rp.Called(id)
	locality = args.Get(0).([]model.LocalitiesJSONCarriers)
	err = args.Error(1)

	return
}

func (rp *LocalityMockRepository) GetSellers(id int) (report []model.LocalitiesJSONSellers, err error) {
	args := rp.Called(id)
	report = args.Get(0).([]model.LocalitiesJSONSellers)
	err = args.Error(1)

	return
}

func (rp *LocalityMockRepository) GetReportSellersWithId(id int) (locality []model.LocalitiesJSONSellers, err error) {
	args := rp.Called(id)
	locality = args.Get(0).([]model.LocalitiesJSONSellers)
	err = args.Error(1)

	return
}

func (rp *LocalityMockRepository) Get() (localities []model.Locality, err error) {
	args := rp.Called()
	localities = args.Get(0).([]model.Locality)
	err = args.Error(1)

	return
}

func (rp *LocalityMockRepository) GetById(id int) (l model.Locality, err error) {
	args := rp.Called(id)
	l = args.Get(0).(model.Locality)
	err = args.Error(1)

	return
}

func (rp *LocalityMockRepository) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	args := rp.Called(locality)
	l = args.Get(0).(model.Locality)
	err = args.Error(1)

	return
}
