package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
)



type BuyerService struct {
	rp repository.BuyerRepository
}


func NewBuyerService(rp repository.BuyerRepository) *BuyerService{
	return &BuyerService{rp}
}


func (bs *BuyerService) GetAllBuyer() (map[int]model.Buyer, error){
	return bs.rp.Get()
}