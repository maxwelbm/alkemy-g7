package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IPurchaseOrdersService interface {
	GetPurchaseOrderByID(id int) (purchaseOrder model.PurchaseOrder, err error)
	CreatePurchaseOrder(newPurchaseOrder model.PurchaseOrder) (purchaseOrder model.PurchaseOrder, err error)
}
