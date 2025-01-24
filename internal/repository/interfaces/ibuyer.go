package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IBuyerRepo interface {
	Get() (buyers []model.Buyer, err error)
	GetByID(id int) (buyer model.Buyer, err error)
	Post(newBuyer model.Buyer) (id int64, err error)
	Update(id int, newBuyer model.Buyer) (err error)
	Delete(id int) (err error)
	CountPurchaseOrderByBuyerID(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error)
	CountPurchaseOrderBuyers() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error)
}
