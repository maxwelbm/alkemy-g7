package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type PurchaseOrderServiceMock struct {
	mock.Mock
}

func (p *PurchaseOrderServiceMock) GetPurchaseOrderByID(id int) (purchaseOrder model.PurchaseOrder, err error) {
	args := p.Called(id)
	purchaseOrder = args.Get(0).(model.PurchaseOrder)
	err = args.Error(1)

	return
}

func (p *PurchaseOrderServiceMock) CreatePurchaseOrder(newPurchaseOrder model.PurchaseOrder) (purchaseOrder model.PurchaseOrder, err error) {
	args := p.Called(newPurchaseOrder)

	purchaseOrder, _ = args.Get(0).(model.PurchaseOrder)
	err = args.Error(1)

	return
}
