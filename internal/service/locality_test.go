package service_test

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupLocalityServiceTest(t *testing.T) *service.LocalitiesService {
	mock := mocks.NewMockILocalityRepo(t)
	return service.CreateServiceLocalities(mock)
}

func TestLocalitiesService_GetByID(t *testing.T) {
	s := setupLocalityServiceTest(t)
	mock := s.Rp.(*mocks.MockILocalityRepo)

	t.Run("test service method for get localities by ID successfully", func(t *testing.T) {
		ID := 1
		l := model.Locality{ID: 1, Locality: "Tokyo", Province: "Kanto", Country: "Japan"}

		mock.On("GetByID", ID).Return(l, nil).Once()

		locality, err := s.GetByID(ID)

		assert.NoError(t, err)
		assert.Equal(t, l, locality)
		mock.AssertExpectations(t)
	})
}

func TestLocalitiesService_CreateLocality(t *testing.T) {
	s := setupLocalityServiceTest(t)
	mock := s.Rp.(*mocks.MockILocalityRepo)

	t.Run("test service method for create localities successfully", func(t *testing.T) {
		arg := model.Locality{Locality: "Brooklyn", Province: "New York", Country: "EUA"}
		l := model.Locality{ID: 2, Locality: "Brooklyn", Province: "New York", Country: "EUA"}

		mock.On("CreateLocality", &arg).Return(l, nil).Once()

		locality, err := s.CreateLocality(&arg)

		assert.NoError(t, err)
		assert.Equal(t, l, locality)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for create localities with empty attributes", func(t *testing.T) {
		arg := model.Locality{Locality: "", Province: "", Country: ""}
		l := model.Locality{}
		errS := customerror.ErrNullLocalityAttribute

		locality, err := s.CreateLocality(&arg)

		assert.ErrorIs(t, errS, err)
		assert.Equal(t, l, locality)
	})
}

func TestLocalitiesService_GetSellers(t *testing.T) {
	s := setupLocalityServiceTest(t)
	mock := s.Rp.(*mocks.MockILocalityRepo)

	t.Run("test service method for get report sellers by ID successfully", func(t *testing.T) {
		ID := 3
		l := []model.LocalitiesJSONSellers{{ID: "3", Locality: "Phoenix", Sellers: 5}}

		mock.On("GetReportSellersWithID", ID).Return(l, nil).Once()

		report, err := s.GetSellers(ID)

		assert.NoError(t, err)
		assert.Equal(t, l, report)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for get report all sellers successfully", func(t *testing.T) {
		ID := 0
		l := []model.LocalitiesJSONSellers{
			{ID: "4", Locality: "Phoenix", Sellers: 5},
			{ID: "5", Locality: "Kanto", Sellers: 8},
			{ID: "6", Locality: "New York", Sellers: 9},
			{ID: "7", Locality: "Kansas", Sellers: 3},
		}

		mock.On("GetSellers", ID).Return(l, nil).Once()

		report, err := s.GetSellers(ID)

		assert.NoError(t, err)
		assert.Equal(t, l, report)
		mock.AssertExpectations(t)
	})
}

func TestLocalitiesService_GetCarriers(t *testing.T) {
	s := setupLocalityServiceTest(t)
	mock := s.Rp.(*mocks.MockILocalityRepo)

	t.Run("test service method for get report carriers by ID successfully", func(t *testing.T) {
		ID := 3
		l := []model.LocalitiesJSONCarriers{{ID: "3", Locality: "Phoenix", Carriers: 5}}

		mock.On("GetReportCarriersWithID", ID).Return(l, nil).Once()

		report, err := s.GetCarriers(ID)

		assert.NoError(t, err)
		assert.Equal(t, l, report)
		mock.AssertExpectations(t)
	})

	t.Run("test service method for get report all carriers successfully", func(t *testing.T) {
		ID := 0
		l := []model.LocalitiesJSONCarriers{
			{ID: "4", Locality: "Phoenix", Carriers: 5},
			{ID: "5", Locality: "Kanto", Carriers: 8},
			{ID: "6", Locality: "New York", Carriers: 9},
			{ID: "7", Locality: "Kansas", Carriers: 3},
		}

		mock.On("GetCarriers", ID).Return(l, nil).Once()

		report, err := s.GetCarriers(ID)

		assert.NoError(t, err)
		assert.Equal(t, l, report)
		mock.AssertExpectations(t)
	})
}
