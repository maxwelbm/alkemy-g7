package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type PurchaseOrderService struct {
	Rp            interfaces.IPurchaseOrdersRepo
	SvcBuyer      svc.IBuyerservice
	SvcProductRec svc.IProductRecService
}

func (p *PurchaseOrderService) CreatePurchaseOrder(newPurchaseOrder model.PurchaseOrder) (purchaseOrder model.PurchaseOrder, err error) {
	_, err = p.SvcBuyer.GetBuyerByID(newPurchaseOrder.BuyerID)

	if err != nil {
		return
	}

	_, err = p.SvcProductRec.GetProductRecordByID(newPurchaseOrder.ProductRecordID)

	if err != nil {
		return
	}

	id, err := p.Rp.Post(newPurchaseOrder)

	if err != nil {
		return
	}

	purchaseOrder, err = p.Rp.GetById(int(id))

	return
}

func (p *PurchaseOrderService) GetPurchaseOrderByID(id int) (purchaseOrder model.PurchaseOrder, err error) {
	return p.Rp.GetById(id)
}
func NewPurchaseOrderService(rp interfaces.IPurchaseOrdersRepo, svcBuyer svc.IBuyerservice, svcProductRec svc.IProductRecService) *PurchaseOrderService {
	return &PurchaseOrderService{Rp: rp, SvcBuyer: svcBuyer, SvcProductRec: svcProductRec}
}
