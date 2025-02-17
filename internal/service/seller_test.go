package service_test

import (
	"errors"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setupSeller(t *testing.T) *service.SellersService {
	mockSeller := mocks.NewMockISellerRepo(t)
	mockLocality := mocks.NewMockILocalityRepo(t)

	return service.CreateServiceSellers(mockSeller, mockLocality)
}

func setupLocality(mockLocality *mocks.MockILocalityRepo) *service.LocalitiesService {
	return service.CreateServiceLocalities(mockLocality)
}

func TestSellersService_GetAll(t *testing.T) {
	s := setupSeller(t)
	mock := s.Rp.(*mocks.MockISellerRepo)

	t.Run("test service method for get a list of all existing sellers successfully", func(t *testing.T) {
		sl := []model.Seller{{ID: 1, CID: 1, CompanyName: "Enterprise Liberty", Address: "456 Elm St", Telephone: "4443335454", Locality: 1},
			{ID: 2, CID: 2, CompanyName: "Libre Mercado", Address: "123 Montain St Avenue", Telephone: "5554545999", Locality: 2}}

		mock.On("Get").Return(sl, nil).Once()

		sellers, err := s.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, sl, sellers)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for get a list of all existing sellers with internal server error", func(t *testing.T) {
		sl := []model.Seller{}
		errS := errors.New("internal server error")

		mock.On("Get").Return(sl, errS).Once()

		sellers, err := s.GetAll()

		assert.ErrorIs(t, errS, err)
		assert.Equal(t, sl, sellers)
		mock.AssertExpectations(t)
	})
}

func TestSellersService_GetByID(t *testing.T) {
	s := setupSeller(t)
	mock := s.Rp.(*mocks.MockISellerRepo)

	t.Run("test service method for get seller by ID existing successfully", func(t *testing.T) {
		ID := 3
		sl := model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 3}

		mock.On("GetByID", ID).Return(sl, nil).Once()

		seller, err := s.GetByID(ID)

		assert.NoError(t, err)
		assert.Equal(t, sl, seller)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for get seller by ID not found", func(t *testing.T) {
		sl := model.Seller{}
		ID := 999
		errS := customerror.ErrSellerNotFound

		mock.On("GetByID", ID).Return(sl, errS).Once()

		seller, err := s.GetByID(ID)

		assert.ErrorIs(t, errS, err)
		assert.Equal(t, sl, seller)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for get seller by ID with zero id", func(t *testing.T) {
		sl := model.Seller{}
		ID := 0
		errS := customerror.ErrMissingSellerID

		mock.On("GetByID", ID).Return(sl, errS).Once()

		seller, err := s.GetByID(ID)

		assert.ErrorIs(t, errS, err)
		assert.Equal(t, sl, seller)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for get seller by ID with internal server error", func(t *testing.T) {
		sl := model.Seller{}
		ID := 4
		errS := customerror.ErrDefaultSeller

		mock.On("GetByID", ID).Return(sl, errS).Once()

		seller, err := s.GetByID(ID)

		assert.ErrorIs(t, errS, err)
		assert.Equal(t, sl, seller)
		mock.AssertExpectations(t)
	})
}

func TestSellersService_CreateSeller(t *testing.T) {
	s := setupSeller(t)
	mockSeller := s.Rp.(*mocks.MockISellerRepo)
	mockLocality := s.Rpl.(*mocks.MockILocalityRepo)

	t.Run("test service method for create seller with success", func(t *testing.T) {
		arg := model.Seller{CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5}
		sl := model.Seller{ID: 5, CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5}
		ID := 5
		l := model.Locality{ID: 5, Locality: "Brooklyn", Province: "New York", Country: "EUA"}

		mockSeller.On("Post", &arg).Return(sl, nil).Once()
		mockLocality.On("GetByID", ID).Return(l, nil).Once()

		seller, err := s.CreateSeller(&arg)

		assert.NoError(t, err)
		assert.Equal(t, sl, seller)
		mockSeller.AssertExpectations(t)
		mockLocality.AssertExpectations(t)
	})

	t.Run("test service method for create seller with empty attributes values", func(t *testing.T) {
		arg := model.Seller{}
		sl := model.Seller{}
		errSeller := customerror.ErrNullSellerAttribute

		seller, err := s.CreateSeller(&arg)

		assert.ErrorIs(t, errSeller, err)
		assert.Equal(t, sl, seller)
	})

	t.Run("test service method for create seller with attribute CID already existing", func(t *testing.T) {
		arg := model.Seller{CID: 1, CompanyName: "Midgard Sellers", Address: "3 New Time Park", Telephone: "99989898778", Locality: 7}
		sl := model.Seller{}
		ID := 7
		l := model.Locality{ID: 7, Locality: "Manhattan", Province: "New York", Country: "EUA"}
		errSeller := customerror.ErrCIDSellerAlreadyExist

		mockSeller.On("Post", &arg).Return(sl, errSeller).Once()
		mockLocality.On("GetByID", ID).Return(l, nil).Once()

		seller, err := s.CreateSeller(&arg)

		assert.ErrorIs(t, errSeller, err)
		assert.Equal(t, sl, seller)
		mockSeller.AssertExpectations(t)
		mockLocality.AssertExpectations(t)
	})

	t.Run("test service method for create seller with attribute locality ID not found", func(t *testing.T) {
		arg := model.Seller{CID: 8, CompanyName: "Rupture Clivers", Address: "1200 New Time Park", Telephone: "7776657987", Locality: 9999}
		sl := model.Seller{}
		ID := 9999
		l := model.Locality{}
		errSeller := customerror.ErrLocalityNotFound
		errLocality := customerror.ErrLocalityNotFound

		mockLocality.On("GetByID", ID).Return(l, errLocality).Once()
		seller, err := s.CreateSeller(&arg)

		assert.ErrorIs(t, errSeller, err)
		assert.Equal(t, sl, seller)
		mockLocality.AssertExpectations(t)
	})
}

func TestSellersService_UpdateSeller(t *testing.T) {
	serviceSeller := setupSeller(t)
	mockSeller := serviceSeller.Rp.(*mocks.MockISellerRepo)
	mockLocality := serviceSeller.Rpl.(*mocks.MockILocalityRepo)
	serviceLocality := setupLocality(mockLocality)

	t.Run("test service method for update seller with success", func(t *testing.T) {
		arg := model.Seller{CID: 55, CompanyName: "Cypress Company", Address: "900 Central Park", Telephone: "55566777787", Locality: 10}
		sl := model.Seller{ID: 5, CID: 55, CompanyName: "Cypress Company", Address: "900 Central Park", Telephone: "55566777787", Locality: 10}
		sellerID := 5
		localityID := 10
		l := model.Locality{ID: 10, Locality: "Los Angeles", Province: "California", Country: "EUA"}

		mockSeller.On("Patch", sellerID, &arg).Return(sl, nil)
		mockSeller.On("GetByID", sellerID).Return(sl, nil)
		mockLocality.On("GetByID", localityID).Return(l, nil)

		seller, err := serviceSeller.UpdateSeller(sellerID, &arg)
		locality, errL := serviceLocality.GetByID(localityID)

		assert.NoError(t, err)
		assert.NoError(t, errL)
		assert.Equal(t, sl, seller)
		assert.Equal(t, l, locality)
		mockSeller.AssertExpectations(t)
		mockLocality.AssertExpectations(t)
	})

	t.Run("test service method for update seller with ID not found", func(t *testing.T) {
		arg := model.Seller{CID: 65, CompanyName: "Cypress Company", Address: "30 Central Park", Telephone: "55566777787", Locality: 9}
		sl := model.Seller{}
		sellerID := 999
		localityID := 9
		l := model.Locality{ID: 9, Locality: "Little Rock", Province: "Arkansas", Country: "EUA"}
		errSeller := customerror.ErrSellerNotFound

		mockSeller.On("Patch", sellerID, &arg).Return(sl, errSeller)
		mockSeller.On("GetByID", sellerID).Return(sl, errSeller)
		mockLocality.On("GetByID", localityID).Return(l, nil)

		seller, err := serviceSeller.UpdateSeller(sellerID, &arg)
		locality, errL := serviceLocality.GetByID(localityID)

		assert.ErrorIs(t, errSeller, err)
		assert.NoError(t, errL)
		assert.Equal(t, sl, seller)
		assert.Equal(t, l, locality)
		mockSeller.AssertExpectations(t)
		mockLocality.AssertExpectations(t)
	})

	t.Run("test service method for update seller with empty attributes values", func(t *testing.T) {
		arg := model.Seller{CID: 0, CompanyName: "", Address: "", Telephone: "", Locality: 0}
		sl := model.Seller{}
		sellerID := 2
		errSeller := customerror.ErrNullSellerAttribute

		mockSeller.On("GetByID", sellerID).Return(sl, errSeller)

		seller, err := serviceSeller.UpdateSeller(sellerID, &arg)

		assert.ErrorIs(t, errSeller, err)
		assert.Equal(t, sl, seller)
		mockSeller.AssertExpectations(t)
	})

	t.Run("test service method for update seller with attribute CID already existing", func(t *testing.T) {
		arg := model.Seller{CID: 1, CompanyName: "Cypress Company", Address: "400 Central Park", Telephone: "55566777787", Locality: 17}
		sl := model.Seller{}
		sellerID := 9
		localityID := 17
		l := model.Locality{ID: 17, Locality: "Phoenix", Province: "Arizona", Country: "EUA"}
		errSeller := customerror.ErrCIDSellerAlreadyExist

		mockSeller.On("Patch", sellerID, &arg).Return(sl, errSeller)
		mockSeller.On("GetByID", sellerID).Return(sl, errSeller)
		mockLocality.On("GetByID", localityID).Return(l, nil)

		seller, err := serviceSeller.UpdateSeller(sellerID, &arg)
		locality, errL := serviceLocality.GetByID(localityID)

		assert.ErrorIs(t, errSeller, err)
		assert.NoError(t, errL)
		assert.Equal(t, sl, seller)
		assert.Equal(t, l, locality)
		mockSeller.AssertExpectations(t)
		mockLocality.AssertExpectations(t)
	})

	t.Run("test service method for update seller with attribute locality ID not found", func(t *testing.T) {
		arg := model.Seller{CID: 8, CompanyName: "Rupture Clivers", Address: "1200 New Time Park", Telephone: "7776657987", Locality: 9999}
		sl := model.Seller{}
		sellerID := 8
		localityID := 9999
		l := model.Locality{}
		errSeller := customerror.ErrLocalityNotFound
		errLocality := customerror.ErrLocalityNotFound

		mockLocality.On("GetByID", localityID).Return(l, errLocality)

		seller, err := serviceSeller.UpdateSeller(sellerID, &arg)
		locality, errL := serviceLocality.GetByID(localityID)

		assert.ErrorIs(t, errSeller, err)
		assert.ErrorIs(t, errLocality, errL)
		assert.Equal(t, sl, seller)
		assert.Equal(t, l, locality)
		mockLocality.AssertExpectations(t)
	})

	t.Run("test service method for update seller with update seller with zero id", func(t *testing.T) {
		arg := model.Seller{CID: 55, CompanyName: "Cypress Company", Address: "400 Central Park", Telephone: "55566777787", Locality: 30}
		sl := model.Seller{}
		sellerID := 0
		localityID := 30
		l := model.Locality{ID: 17, Locality: "Denver", Province: "Colorado", Country: "EUA"}
		errSeller := customerror.ErrMissingSellerID

		mockSeller.On("Patch", sellerID, &arg).Return(sl, errSeller)
		mockSeller.On("GetByID", sellerID).Return(sl, errSeller)
		mockLocality.On("GetByID", localityID).Return(l, nil)

		seller, err := serviceSeller.UpdateSeller(sellerID, &arg)
		locality, errL := serviceLocality.GetByID(localityID)

		assert.ErrorIs(t, errSeller, err)
		assert.NoError(t, errL)
		assert.Equal(t, sl, seller)
		assert.Equal(t, l, locality)
	})
}

func TestSellersService_DeleteSeller(t *testing.T) {
	s := setupSeller(t)
	mock := s.Rp.(*mocks.MockISellerRepo)

	t.Run("test service method for delete seller with success", func(t *testing.T) {
		ID := 3
		mock.On("Delete", ID).Return(nil)

		err := s.DeleteSeller(ID)

		assert.NoError(t, err)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for delete seller by ID not found", func(t *testing.T) {
		ID := 999
		errS := customerror.ErrSellerNotFound
		mock.On("Delete", ID).Return(errS)

		err := s.DeleteSeller(ID)

		assert.ErrorIs(t, errS, err)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for delete seller by zero id", func(t *testing.T) {
		ID := 0
		errS := customerror.ErrMissingSellerID
		mock.On("Delete", ID).Return(errS)

		err := s.DeleteSeller(ID)

		assert.ErrorIs(t, errS, err)
		mock.AssertExpectations(t)
	})
}
