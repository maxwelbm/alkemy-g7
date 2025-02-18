package service

import (
	"fmt"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type PurchaseOrderService struct {
	Rp            interfaces.IPurchaseOrdersRepo
	SvcBuyer      svc.IBuyerservice
	SvcProductRec svc.IProductRecService
	log           logger.Logger
}

func (p *PurchaseOrderService) CreatePurchaseOrder(newPurchaseOrder model.PurchaseOrder) (purchaseOrder model.PurchaseOrder, err error) {
	p.log.Log("PurchaseOrderService", "INFO", fmt.Sprintf("initializing CreatePurchaseOrder function with parameter: %v", newPurchaseOrder))

	p.log.Log("PurchaseOrderService", "INFO", fmt.Sprintf("Searching Buyer with parameter ID: %d", newPurchaseOrder.BuyerID))
	_, err = p.SvcBuyer.GetBuyerByID(newPurchaseOrder.BuyerID)

	if err != nil {
		p.log.Log("PurchaseOrderService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	p.log.Log("PurchaseOrderService", "INFO", "Buyer found")
	p.log.Log("PurchaseOrderService", "INFO", fmt.Sprintf("Searching ProductRecord with parameter ID: %d", newPurchaseOrder.ProductRecordID))
	_, err = p.SvcProductRec.GetProductRecordByID(newPurchaseOrder.ProductRecordID)

	if err != nil {
		p.log.Log("PurchaseOrderService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	p.log.Log("PurchaseOrderService", "INFO", "Product Record found")

	id, err := p.Rp.Post(newPurchaseOrder)

	if err != nil {
		p.log.Log("PurchaseOrderService", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	p.log.Log("PurchaseOrderService", "INFO", fmt.Sprintf("Purchase Order created with ID: %d", id))
	purchaseOrder, err = p.Rp.GetByID(int(id))
	p.log.Log("PurchaseOrderService", "INFO", fmt.Sprintf("Return Purchase Order created with ID: %d PUrchase: %v", id, purchaseOrder))

	return
}

func (p *PurchaseOrderService) GetPurchaseOrderByID(id int) (purchaseOrder model.PurchaseOrder, err error) {
	p.log.Log("PurchaseOrderService", "INFO", fmt.Sprintf("initializing GetPurchaseOrderByID function with parameter: %v", id))
	return p.Rp.GetByID(id)
}
func NewPurchaseOrderService(rp interfaces.IPurchaseOrdersRepo, svcBuyer svc.IBuyerservice, svcProductRec svc.IProductRecService, log logger.Logger) *PurchaseOrderService {
	return &PurchaseOrderService{Rp: rp, SvcBuyer: svcBuyer, SvcProductRec: svcProductRec, log: log}
}
