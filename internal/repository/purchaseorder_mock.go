package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type PurchaseOrderRepositoryMock struct {
	mock.Mock
}

func (p *PurchaseOrderRepositoryMock) GetByID(id int) (purchaseOrder model.PurchaseOrder, err error) {
	args := p.Called(id)
	purchaseOrder = args.Get(0).(model.PurchaseOrder)
	err = args.Error(1)

	return
}

func (p *PurchaseOrderRepositoryMock) Post(newPurchaseOrder model.PurchaseOrder) (id int64, err error) {
	args := p.Called(newPurchaseOrder)
	id = args.Get(0).(int64)
	err = args.Error(1)

	return
}
