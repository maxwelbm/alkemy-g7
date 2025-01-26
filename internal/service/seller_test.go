package service_test

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
	"github.com/stretchr/testify/assert"
)

func setupSeller() *service.SellersService {
	mocksl := new(repository.SellerMockRepository)
	mocklc := new(repository.LocalityMockRepository)

	return service.CreateServiceSellers(mocksl, mocklc)
}

func TestServiceGetAllSeller(t *testing.T) {
	tests := []struct {
		description string
		sellers     []model.Seller
		existErr    bool
		err         error
		call        bool
	}{
		{
			description: "get a list of all existing sellers successfully",
			sellers: []model.Seller{{ID: 1, CID: 1, CompanyName: "Enterprise Liberty", Address: "456 Elm St", Telephone: "4443335454", Locality: 1},
				{ID: 2, CID: 2, CompanyName: "Libre Mercado", Address: "123 Montain St Avenue", Telephone: "5554545999", Locality: 2}},
			existErr: false,
			err:      nil,
			call:     true,
		},
		{
			description: "get a list of all existing sellers with internal server error",
			sellers:     []model.Seller{},
			existErr:    true,
			err:         errors.New("internal server error"),
			call:        true,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			service := setupSeller()
			mock := service.Rp.(*repository.SellerMockRepository)
			mock.On("Get").Return(test.sellers, test.err)

			sellers, err := service.GetAll()

			assert.Equal(t, test.sellers, sellers)
			switch test.existErr {
			case true:
				assert.Error(t, err)
			case false:
				assert.NoError(t, err)
			}

			if test.call {
				mock.AssertExpectations(t)
			}
		})
	}
}

func TestServiceGetByIDSeller(t *testing.T) {
	tests := []struct {
		description string
		seller      model.Seller
		id          int
		existErr    bool
		err         error
		call        bool
	}{
		{
			description: "get seller by id existing successfully",
			seller:      model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Perk Avenue", Telephone: "999444555", Locality: 3},
			id: 3,
			existErr:    false,
			err:         nil,
			call:        true,
		},
		{
			description: "get seller by id not found",
			seller:      model.Seller{},
			id: 999,
			existErr:    true,
			err:         customError.ErrSellerNotFound,
			call:        true,
		},
		{
			description: "get seller by id with zero id",
			seller:      model.Seller{},
			id: 0,
			existErr:    true,
			err:         customError.ErrMissingSellerID,
			call:        true,
		},
		{
			description: "get seller by id with internal server error",
			seller:      model.Seller{},
			id: 4,
			existErr:    true,
			err:         customError.ErrDefaultSeller,
			call:        true,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			service := setupSeller()
			mock := service.Rp.(*repository.SellerMockRepository)
			mock.On("GetById", test.id).Return(test.seller, test.err)

			seller, err := service.GetById(test.id)

			assert.Equal(t, test.seller, seller)
			switch test.existErr {
			case true:
				assert.Error(t, err)
			case false:
				assert.NoError(t, err)
			}

			if test.call {
				mock.AssertExpectations(t)
			}
		})
	}
}
