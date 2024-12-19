package repository

import (
	"errors"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

type BuyerRepository struct {
	dbBuyer database.Database
}

func (br BuyerRepository) Delete(id int) error {
	buyer, err := br.GetById(id)

	if err != nil && errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
		return &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	delete(br.dbBuyer.TbBuyer, buyer.Id)

	return nil
}

func (br *BuyerRepository) Get() (map[int]model.Buyer, error) {

	if len(br.dbBuyer.TbBuyer) == 0 {
		return nil, fmt.Errorf("no buyers found")
	}
	return br.dbBuyer.TbBuyer, nil

}

func (br *BuyerRepository) GetById(id int) (model.Buyer, error) {
	buyer, ok := br.dbBuyer.TbBuyer[id]

	if !ok {
		return model.Buyer{}, &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	return buyer, nil
}

func (br *BuyerRepository) Post(newBuyer model.Buyer) (model.Buyer, error) {
	BuyerExists := isCardNumberIdExists(newBuyer.CardNumberId, br)

	if BuyerExists {
		return model.Buyer{}, &custom_error.CustomError{Object: newBuyer, Err: custom_error.Conflict}
	}
	buyer := model.Buyer{
		Id:           len(br.dbBuyer.TbBuyer) + 1,
		CardNumberId: newBuyer.CardNumberId,
		FirstName:    newBuyer.FirstName,
		LastName:     newBuyer.LastName,
	}

	br.dbBuyer.TbBuyer[newBuyer.Id] = buyer

	return br.GetById(newBuyer.Id)

}

func (br BuyerRepository) Update(id int, newBuyer model.Buyer) (model.Buyer, error) {

    buyer, err := br.GetById(id)
    if err != nil && errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
        return model.Buyer{}, &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
    }

    if existingBuyer, exists := br.dbBuyer.TbBuyer[buyer.Id]; exists {
        
        if newBuyer.CardNumberId != "" {
            existingBuyer.CardNumberId = newBuyer.CardNumberId
        }
        if newBuyer.FirstName != "" {
            existingBuyer.FirstName = newBuyer.FirstName
        }
        if newBuyer.LastName != "" {
            existingBuyer.LastName = newBuyer.LastName
        }

       
        br.dbBuyer.TbBuyer[buyer.Id] = existingBuyer
    } else {
        return model.Buyer{}, &custom_error.CustomError{Object: buyer.Id, Err: custom_error.NotFound}
    }

    return br.GetById(id)
}


func isCardNumberIdExists(CardNumberId string, br *BuyerRepository) bool {

	for _, b := range br.dbBuyer.TbBuyer {
		if b.CardNumberId == CardNumberId {
			return true
		}
	}

	return false
}

func NewBuyerRepository(db database.Database) *BuyerRepository {
	return &BuyerRepository{dbBuyer: db}
}
