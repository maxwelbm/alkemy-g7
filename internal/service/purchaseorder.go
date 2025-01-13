package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)

type PurchaseOrderService struct {
	rp            interfaces.IPurchaseOrdersRepo
	svcBuyer      svc.IBuyerservice
	svcProductRec svc.IProductRecService
}

// CreatePurchaseOrder implements interfaces.IPurchaseOrdersService.
func (p *PurchaseOrderService) CreatePurchaseOrder(newPurchaseOrder model.PurchaseOrder) (purchaseOrder model.PurchaseOrder, err error) {
	_, err = p.svcBuyer.GetBuyerByID(newPurchaseOrder.BuyerId)
	if err != nil {
		return
	}

	_, err = p.svcProductRec.GetProductRecordById(newPurchaseOrder.ProductRecordId)
	if err != nil {
		return
	}

	id, err := p.rp.Post(newPurchaseOrder)
	if err != nil {
		return
	}

	purchaseOrder, err = p.rp.GetById(int(id))

	return

}

func (p *PurchaseOrderService) GetPurchaseOrderByID(id int) (purchaseOrder model.PurchaseOrder, err error) {
	return p.rp.GetById(id)
}
func NewPurchaseOrderService(rp interfaces.IPurchaseOrdersRepo, svcBuyer svc.IBuyerservice, svcProductRec svc.IProductRecRepo) *PurchaseOrderService {
	return &PurchaseOrderService{rp: rp, svcBuyer: svcBuyer, svcProductRec: svcProductRec}
}
