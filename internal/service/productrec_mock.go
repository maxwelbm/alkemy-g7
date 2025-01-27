package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type ProductrecMock struct {
	mock.Mock
}

func (p *ProductrecMock) CreateProductRecords(pr model.ProductRecords) (model.ProductRecords, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductrecMock) GetProductRecordByID(id int) (model.ProductRecords, error) {
	args := p.Called(id)
	return args.Get(0).(model.ProductRecords), args.Error(1)
}

func (p *ProductrecMock) GetProductRecordReport(idProduct int) ([]model.ProductRecordsReport, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductrecMock() *ProductrecMock {
	return &ProductrecMock{}
}
