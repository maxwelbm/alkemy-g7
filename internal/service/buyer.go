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

func (bs *BuyerService) GetAllBuyer() (map[int]model.Buyer, error) {
	return bs.rp.Get()
}

func (bs *BuyerService) GetBuyerByID(id int) (model.Buyer, error) {
	return bs.rp.GetById(id)
}
