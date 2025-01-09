package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
)

type BuyerService struct {
	rp interfaces.IBuyerRepo
}

func NewBuyerService(rp interfaces.IBuyerRepo) *BuyerService {
	return &BuyerService{rp: rp}
}

func (bs *BuyerService) GetAllBuyer() (buyers []model.Buyer, err error) {
	return bs.rp.Get()
}

func (bs *BuyerService) GetBuyerByID(id int) (buyer model.Buyer, err error) {
	return bs.rp.GetById(id)
}

func (bs *BuyerService) DeleteBuyerByID(id int) (err error) {

	_, err = bs.GetBuyerByID(id)
	if err != nil {
		return
	}

	return bs.rp.Delete(id)
}

func (bs *BuyerService) CreateBuyer(newBuyer model.Buyer) (buyer model.Buyer, err error) {

	id, err := bs.rp.Post(newBuyer)

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

	if newBuyer.CardNumberId != "" {
		existingBuyer.CardNumberId = newBuyer.CardNumberId
	}
	if newBuyer.FirstName != "" {
		existingBuyer.FirstName = newBuyer.FirstName
	}
	if newBuyer.LastName != "" {
		existingBuyer.LastName = newBuyer.LastName
	}

	err = bs.rp.Update(existingBuyer.Id, existingBuyer)
	if err != nil {
		return
	}

	buyer, err = bs.GetBuyerByID(id)

	return

}
