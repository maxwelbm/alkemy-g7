package service_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

var logMockCarrier = mocks.MockLog{}

func TestCarrierDefault_GetByID(t *testing.T) {
	t.Run("Sucess GetByIdCarries", func(t *testing.T) {

		expectedCarries := model.Carries{
			ID:          1,
			CID:         "1231231",
			CompanyName: "Chico",
			Address:     "Rua Chicola",
			Telephone:   "DD DO CHICO",
			LocalityID:  1,
		}

		mockRepo := mocks.NewMockICarriersRepo(t)
		mockLocality := mocks.NewMockILocalityRepo(t)
		service := service.NewCarrierService(mockRepo, mockLocality, logMockCarrier)

		mockRepo.On("GetByID", 1).Return(expectedCarries, nil)

		actual, err := service.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedCarries, actual)
	})

	t.Run("Error GetByIDCarries", func(t *testing.T) {

		mockRepo := mocks.NewMockICarriersRepo(t)
		mockLocality := mocks.NewMockILocalityRepo(t)
		service := service.NewCarrierService(mockRepo, mockLocality, logMockCarrier)
		expectedError := customerror.NewCarrierError(customerror.ErrNotFound.Error(), "carrier", http.StatusNotFound)

		mockRepo.On("GetByID", 1).Return(model.Carries{}, expectedError)
		actual, err := service.GetByID(1)

		assert.ErrorIs(t, expectedError, err)

		assert.Empty(t, actual)
	})
}

func TestCarrierDefault_PostCarrier(t *testing.T) {
	t.Run("Sucess PostCarries", func(t *testing.T) {
		expectedCarries := model.Carries{
			ID:          1,
			CID:         "1231231",
			CompanyName: "Chico",
			Address:     "Rua Chicola",
			Telephone:   "DD DO CHICO",
			LocalityID:  1,
		}

		expectedLocality := model.Locality{
			ID:       1,
			Locality: "Chiquinho",
			Province: "Chicola",
			Country:  "ChicoWorld",
		}

		mockRepo := mocks.NewMockICarriersRepo(t)
		mockLocality := mocks.NewMockILocalityService(t)
		service := service.NewCarrierService(mockRepo, mockLocality, logMockCarrier)

		mockLocality = service.SvcLocality.(*mocks.MockILocalityService)
		mockLocality.On("GetByID", expectedCarries.LocalityID).Return(expectedLocality, nil)

		mockRepo.On("PostCarrier", expectedCarries).Return(int64(1), nil)

		mockRepo.On("GetByID", expectedCarries.ID).Return(expectedCarries, nil)

		actual, err := service.PostCarrier(expectedCarries)

		assert.NoError(t, err)
		assert.Equal(t, expectedCarries, actual)
	})

	t.Run("Error GetLocality", func(t *testing.T) {
		expectedCarries := model.Carries{
			ID:          1,
			CID:         "1231231",
			CompanyName: "Chico",
			Address:     "Rua Chicola",
			Telephone:   "DD DO CHICO",
			LocalityID:  1,
		}

		mockRepo := mocks.NewMockICarriersRepo(t)
		mockLocality := mocks.NewMockILocalityService(t)
		service := service.NewCarrierService(mockRepo, mockLocality, logMockCarrier)

		mockLocality.On("GetByID", expectedCarries.LocalityID).Return(model.Locality{}, errors.New("locality not found"))

		actual, err := service.PostCarrier(expectedCarries)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "locality not found")
		assert.Empty(t, actual)
	})

	t.Run("Error PostCarrier", func(t *testing.T) {
		expectedCarries := model.Carries{
			ID:          1,
			CID:         "1231231",
			CompanyName: "Chico",
			Address:     "Rua Chicola",
			Telephone:   "DD DO CHICO",
			LocalityID:  1,
		}

		expectedLocality := model.Locality{
			ID:       1,
			Locality: "Chiquinho",
			Province: "Chicola",
			Country:  "ChicoWorld",
		}

		mockRepo := mocks.NewMockICarriersRepo(t)
		mockLocality := mocks.NewMockILocalityService(t)
		service := service.NewCarrierService(mockRepo, mockLocality, logMockCarrier)

		mockLocality.On("GetByID", expectedCarries.LocalityID).Return(expectedLocality, nil)
		mockRepo.On("PostCarrier", expectedCarries).Return(int64(0), errors.New("failed to post carrier"))

		actual, err := service.PostCarrier(expectedCarries)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to post carrier")
		assert.Empty(t, actual)
	})

	t.Run("Error GetByID", func(t *testing.T) {
		expectedCarries := model.Carries{
			ID:          1,
			CID:         "1231231",
			CompanyName: "Chico",
			Address:     "Rua Chicola",
			Telephone:   "DD DO CHICO",
			LocalityID:  1,
		}

		expectedLocality := model.Locality{
			ID:       1,
			Locality: "Chiquinho",
			Province: "Chicola",
			Country:  "ChicoWorld",
		}

		mockRepo := mocks.NewMockICarriersRepo(t)
		mockLocality := mocks.NewMockILocalityService(t)
		service := service.NewCarrierService(mockRepo, mockLocality, logMockCarrier)

		mockLocality.On("GetByID", expectedCarries.LocalityID).Return(expectedLocality, nil)
		mockRepo.On("PostCarrier", expectedCarries).Return(int64(1), nil)
		mockRepo.On("GetByID", expectedCarries.ID).Return(model.Carries{}, errors.New("carrier not found"))

		actual, err := service.PostCarrier(expectedCarries)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "carrier not found")
		assert.Empty(t, actual)
	})

}
