package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IPurchaseOrdersRepo interface {
	GetById(id int) (purchaseOrder model.PurchaseOrder, err error)
	Post(newPurchaseOrder model.PurchaseOrder) (id int64, err error)
}
