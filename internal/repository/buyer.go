package repository

import (
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

type BuyerRepository struct {
	dbBuyer database.Database
}

// Delete implements interfaces.IBuyerRepo.
func (br BuyerRepository) Delete(id int) error {
	panic("unimplemented")
}

// Get implements interfaces.IBuyerRepo.
func (br *BuyerRepository) Get() (map[int]model.Buyer, error) {

	if len(br.dbBuyer.TbBuyer) == 0 {
		return nil, fmt.Errorf("no buyers found")
	}
	return br.dbBuyer.TbBuyer, nil

}

// GetById implements interfaces.IBuyerRepo.
func (br BuyerRepository) GetById(id int) (model.Buyer, error) {
	panic("unimplemented")
}

// Post implements interfaces.IBuyerRepo.
func (br BuyerRepository) Post(buyer model.Buyer) (model.Buyer, error) {
	panic("unimplemented")
}

// Update implements interfaces.IBuyerRepo.
func (br BuyerRepository) Update(id int, buyer model.Buyer) (model.Buyer, error) {
	panic("unimplemented")
}

func NewBuyerRepository(db database.Database) *BuyerRepository {
	return &BuyerRepository{dbBuyer: db}
}
