package service_test

import (
	"net/http"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

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
		service := service.NewCarrierService(mockRepo, mockLocality)

		mockRepo.On("GetByID", 1).Return(expectedCarries, nil)

		actual, err := service.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedCarries, actual)
	})

	t.Run("Error GetByIDCarries", func(t *testing.T) {

		mockRepo := mocks.NewMockICarriersRepo(t)
		mockLocality := mocks.NewMockILocalityRepo(t)
		service := service.NewCarrierService(mockRepo, mockLocality)
		expectedError := customerror.NewCarrierError(customerror.ErrNotFound.Error(), "carrier", http.StatusNotFound)

		mockRepo.On("GetByID", 1).Return(model.Carries{}, expectedError)
		actual, err := service.GetByID(1)

		assert.ErrorIs(t, expectedError, err)

		assert.Empty(t, actual)
	})
}
