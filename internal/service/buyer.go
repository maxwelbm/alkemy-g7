package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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

func (bs *BuyerService) DeleteBuyerByID(id int) error {
	return bs.rp.Delete(id)
}

func (bs *BuyerService) CreateBuyer(newBuyer model.Buyer) (model.Buyer, error) {
	return bs.rp.Post(newBuyer)
}

func (bs *BuyerService) UpdateBuyer(id int, newBuyer model.Buyer) (model.Buyer, error) {

	existingBuyer, err := bs.GetBuyerByID(id)

	if newBuyer.CardNumberId == "" && newBuyer.FirstName == "" && newBuyer.LastName == "" {
		return model.Buyer{}, custom_error.EmptyFields
	}

	if err == nil {

		if newBuyer.CardNumberId != "" {
			existingBuyer.CardNumberId = newBuyer.CardNumberId
		}
		if newBuyer.FirstName != "" {
			existingBuyer.FirstName = newBuyer.FirstName
		}
		if newBuyer.LastName != "" {
			existingBuyer.LastName = newBuyer.LastName
		}

		err := bs.rp.Update(id, existingBuyer)
		if err != nil {
			return model.Buyer{}, err
		}

		return model.Buyer{}, nil
	} else {

		return model.Buyer{}, err
	}
}
