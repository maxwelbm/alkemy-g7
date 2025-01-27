package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type BuyerService struct {
	Rp interfaces.IBuyerRepo
}

func NewBuyerService(rp interfaces.IBuyerRepo) *BuyerService {
	return &BuyerService{Rp: rp}
}

func (bs *BuyerService) GetAllBuyer() (buyers []model.Buyer, err error) {
	return bs.Rp.Get()
}

func (bs *BuyerService) GetBuyerByID(id int) (buyer model.Buyer, err error) {
	return bs.Rp.GetByID(id)
}

func (bs *BuyerService) DeleteBuyerByID(id int) (err error) {
	_, err = bs.GetBuyerByID(id)

	if err != nil {
		return
	}

	return bs.Rp.Delete(id)
}

func (bs *BuyerService) CreateBuyer(newBuyer model.Buyer) (buyer model.Buyer, err error) {
	id, err := bs.Rp.Post(newBuyer)

	if err != nil {
		return
	}

	buyer, err = bs.GetBuyerByID(int(id))

	return
}

func (bs *BuyerService) UpdateBuyer(id int, newBuyer model.Buyer) (buyer model.Buyer, err error) {
	existingBuyer, err := bs.GetBuyerByID(id)

	if err != nil {
		return
	}

	if newBuyer.CardNumberID != "" {
		existingBuyer.CardNumberID = newBuyer.CardNumberID
	}

	if newBuyer.FirstName != "" {
		existingBuyer.FirstName = newBuyer.FirstName
	}

	if newBuyer.LastName != "" {
		existingBuyer.LastName = newBuyer.LastName
	}

	err = bs.Rp.Update(existingBuyer.ID, existingBuyer)

	if err != nil {
		return
	}

	buyer, err = bs.GetBuyerByID(id)

	return
}

func (bs *BuyerService) CountPurchaseOrderByBuyerID(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error) {
	countBuyerPurchaseOrder, err = bs.Rp.CountPurchaseOrderByBuyerID(id)
	return
}

func (bs *BuyerService) CountPurchaseOrderBuyer() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error) {
	countBuyerPurchaseOrder, err = bs.Rp.CountPurchaseOrderBuyers()
	return
}
