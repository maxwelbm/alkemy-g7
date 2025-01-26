package service_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
	"github.com/stretchr/testify/assert"
)

func setupSeller() *service.SellersService {
	mockSeller := new(repository.SellerMockRepository)
	mockLocality := new(repository.LocalityMockRepository)

	return service.CreateServiceSellers(mockSeller, mockLocality)
}

func setupLocality(mockLocality *repository.LocalityMockRepository) *service.LocalitiesService {
	return service.CreateServiceLocalities(mockLocality)
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
				assert.EqualError(t, test.err, err.Error())
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
			seller:      model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central park Avenue", Telephone: "999444555", Locality: 3},
			id:          3,
			existErr:    false,
			err:         nil,
			call:        true,
		},
		{
			description: "get seller by id not found",
			seller:      model.Seller{},
			id:          999,
			existErr:    true,
			err:         customError.ErrSellerNotFound,
			call:        true,
		},
		{
			description: "get seller by id with zero id",
			seller:      model.Seller{},
			id:          0,
			existErr:    true,
			err:         customError.ErrMissingSellerID,
			call:        true,
		},
		{
			description: "get seller by id with internal server error",
			seller:      model.Seller{},
			id:          4,
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
				assert.EqualError(t, test.err, err.Error())
			case false:
				assert.NoError(t, err)
			}

			if test.call {
				mock.AssertExpectations(t)
			}
		})
	}
}

func TestServiceCreateSeller(t *testing.T) {
	tests := []struct {
		description string
		arg         model.Seller
		seller      model.Seller
		id          int
		locality    model.Locality
		validations map[string]bool
		errSeller   error
		errLocality error
	}{
		{
			description: "create seller with success",
			arg:         model.Seller{CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5},
			seller:      model.Seller{ID: 5, CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5},
			id:          5,
			locality:    model.Locality{ID: 5, Locality: "Brooklyn", Province: "New York", Country: "EUA"},
			validations: map[string]bool{
				"existErrSeller":   false,
				"existErrLocality": false,
				"callSeller":       true,
				"callLocality":     true,
			},
			errSeller:   nil,
			errLocality: nil,
		},
		{
			description: "create seller with empty attributes values",
			arg:         model.Seller{CID: 0, CompanyName: "", Address: "", Telephone: "", Locality: 0},
			seller:      model.Seller{},
			id:          0,
			locality:    model.Locality{},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": true,
				"callSeller":       false,
				"callLocality":     false,
			},
			errSeller:   customError.ErrNullSellerAttribute,
			errLocality: nil,
		},
		{
			description: "create seller with attribute cid already existing",
			arg:         model.Seller{CID: 1, CompanyName: "Midgard Sellers", Address: "3 New Time Park", Telephone: "99989898778", Locality: 7},
			seller:      model.Seller{},
			id:          7,
			locality:    model.Locality{ID: 7, Locality: "Manhattan", Province: "New York", Country: "EUA"},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": true,
				"callSeller":       true,
				"callLocality":     true,
			},
			errSeller:   customError.ErrCIDSellerAlreadyExist,
			errLocality: nil,
		},
		{
			description: "create seller with attribute locality id not found",
			arg:         model.Seller{CID: 8, CompanyName: "Rupture Clivers", Address: "1200 New Time Park", Telephone: "7776657987", Locality: 9999},
			seller:      model.Seller{},
			id:          9999,
			locality:    model.Locality{},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": true,
				"callSeller":       false,
				"callLocality":     true,
			},
			errSeller:   customError.ErrLocalityNotFound,
			errLocality: customError.ErrLocalityNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			service := setupSeller()
			mockSeller := service.Rp.(*repository.SellerMockRepository)
			mockLocality := service.Rpl.(*repository.LocalityMockRepository)
			mockSeller.On("Post", &test.arg).Return(test.seller, test.errSeller)
			mockLocality.On("GetById", test.id).Return(test.locality, test.errLocality)

			seller, err := service.CreateSeller(&test.arg)

			assert.Equal(t, test.seller, seller)

			for key, value := range test.validations {
				if strings.Contains(key, "exist") {
					switch value {
					case true:
						assert.Error(t, err)
					case false:
						assert.NoError(t, err)
					}
				}

				if strings.Contains(key, "callSeller") {
					switch value {
					case true:
						mockSeller.AssertExpectations(t)
					}
				}

				if strings.Contains(key, "callLocality") {
					switch value {
					case true:
						mockLocality.AssertExpectations(t)
					}
				}
			}

			if test.errLocality != nil {
				assert.EqualError(t, test.errLocality, err.Error())
			}

			if test.errSeller != nil {
				assert.EqualError(t, test.errSeller, err.Error())
			}
		})
	}
}

func TestServiceUpdateSeller(t *testing.T) {
	tests := []struct {
		description string
		arg         model.Seller
		seller      model.Seller
		sellerID    int
		localityID  int
		locality    model.Locality
		validations map[string]bool
		errSeller   error
		errLocality error
	}{
		{
			description: "update seller with success",
			arg:         model.Seller{CID: 55, CompanyName: "Cypress Company", Address: "900 Central park", Telephone: "55566777787", Locality: 10},
			seller:      model.Seller{ID: 5, CID: 55, CompanyName: "Cypress Company", Address: "900 Central park", Telephone: "55566777787", Locality: 10},
			sellerID:    5,
			localityID:  10,
			locality:    model.Locality{ID: 10, Locality: "Los Angeles", Province: "California", Country: "EUA"},
			validations: map[string]bool{
				"existErrSeller":   false,
				"existErrLocality": false,
				"callSeller":       true,
				"callLocality":     true,
			},
			errSeller:   nil,
			errLocality: nil,
		},
		{
			description: "update seller with id not found",
			arg:         model.Seller{CID: 65, CompanyName: "Cypress Company", Address: "30 Central park", Telephone: "55566777787", Locality: 9},
			seller:      model.Seller{},
			sellerID:    999,
			localityID:  9,
			locality:    model.Locality{ID: 9, Locality: "Little Rock", Province: "Arkansas", Country: "EUA"},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": false,
				"callSeller":       false,
				"callLocality":     false,
			},
			errSeller:   customError.ErrSellerNotFound,
			errLocality: nil,
		},
		{
			description: "update seller with empty attributes values",
			arg:         model.Seller{CID: 0, CompanyName: "", Address: "", Telephone: "", Locality: 9},
			seller:      model.Seller{},
			sellerID:    2,
			localityID:  9,
			locality:    model.Locality{ID: 9, Locality: "Little Rock", Province: "Arkansas", Country: "EUA"},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": false,
				"callSeller":       false,
				"callLocality":     false,
			},
			errSeller:   customError.ErrNullSellerAttribute,
			errLocality: nil,
		},
		{
			description: "update seller with attribute cid already existing",
			arg:         model.Seller{CID: 1, CompanyName: "Cypress Company", Address: "400 Central park", Telephone: "55566777787", Locality: 17},
			seller:      model.Seller{},
			sellerID:    9,
			localityID:  17,
			locality:    model.Locality{ID: 17, Locality: "Phoenix", Province: "Arizona", Country: "EUA"},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": false,
				"callSeller":       true,
				"callLocality":     false,
			},
			errSeller:   customError.ErrCIDSellerAlreadyExist,
			errLocality: nil,
		},
		{
			description: "update seller with attribute locality id not found",
			arg:         model.Seller{CID: 8, CompanyName: "Rupture Clivers", Address: "1200 New Time Park", Telephone: "7776657987", Locality: 9999},
			seller:      model.Seller{},
			sellerID:    8,
			localityID:  9999,
			locality:    model.Locality{},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": true,
				"callSeller":       false,
				"callLocality":     false,
			},
			errSeller:   customError.ErrLocalityNotFound,
			errLocality: customError.ErrLocalityNotFound,
		},
		{
			description: "update seller with zero id",
			arg:         model.Seller{CID: 55, CompanyName: "Cypress Company", Address: "400 Central park", Telephone: "55566777787", Locality: 30},
			seller:      model.Seller{},
			sellerID:    0,
			localityID:  30,
			locality:    model.Locality{ID: 17, Locality: "Denver", Province: "Colorado", Country: "EUA"},
			validations: map[string]bool{
				"existErrSeller":   true,
				"existErrLocality": false,
				"callSeller":       false,
				"callLocality":     false,
			},
			errSeller:   customError.ErrMissingSellerID,
			errLocality: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			serviceSeller := setupSeller()
			mockSeller := serviceSeller.Rp.(*repository.SellerMockRepository)
			mockLocality := serviceSeller.Rpl.(*repository.LocalityMockRepository)
			serviceLocality := setupLocality(mockLocality)
			mockSeller.On("Patch", test.sellerID, &test.arg).Return(test.seller, test.errSeller)
			mockSeller.On("GetById", test.sellerID).Return(test.seller, test.errSeller)
			mockLocality.On("GetById", test.localityID).Return(test.locality, test.errLocality)

			seller, err := serviceSeller.UpdateSeller(test.sellerID, &test.arg)
			locality, errl := serviceLocality.GetById(test.localityID)

			assert.Equal(t, test.seller, seller)
			assert.Equal(t, test.locality, locality)

			for key, value := range test.validations {
				if strings.Contains(key, "exist") {
					switch value {
					case true:
						assert.Error(t, err)
					case false:
						assert.NoError(t, errl)
					}
				}

				if strings.Contains(key, "callSeller") {
					switch value {
					case true:
						mockSeller.AssertExpectations(t)
					}
				}

				if strings.Contains(key, "callLocality") {
					switch value {
					case true:
						mockLocality.AssertExpectations(t)
					}
				}
			}

			if test.errLocality != nil {
				assert.EqualError(t, test.errLocality, err.Error())
			}

			if test.errSeller != nil {
				assert.EqualError(t, test.errSeller, err.Error())
			}
		})
	}
}
