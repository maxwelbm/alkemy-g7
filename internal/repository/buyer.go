package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)


type BuyerRepository struct {
	dbBuyer map[int]model.Buyer
}


func NewBuyerRepository(db map[int]model.Buyer) *BuyerRepository{
	return &BuyerRepository{dbBuyer: db}
}


func (br *BuyerRepository) Get() (map[int]model.Buyer, error){
	return br.dbBuyer, nil
}