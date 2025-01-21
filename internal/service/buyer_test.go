package service_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
	"github.com/stretchr/testify/assert"
)

func setup() *service.BuyerService {
	mockRepo := new(repository.MockBuyerRepo)

	return service.NewBuyerService(mockRepo)
}

func TestGetAllBuyer(t *testing.T) {
	t.Run("return a list of all existing buyers successfully", func(t *testing.T) {
		svc := setup()

		expectedBuyers := []model.Buyer{{Id: 1, FirstName: "John", LastName: "Doe", CardNumberId: "1234"},
			{Id: 2, FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}}

		mockRepo := svc.Rp.(*repository.MockBuyerRepo)
		mockRepo.On("Get").Return(expectedBuyers, nil)

		buyers, err := svc.GetAllBuyer()

		assert.Equal(t, expectedBuyers, buyers)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("return an error when fetching buyers", func(t *testing.T) {
		svc := setup()

		mockRepo := svc.Rp.(*repository.MockBuyerRepo)
		mockRepo.On("Get").Return([]model.Buyer{}, errors.New("Unmapped error"))

		buyers, err := svc.GetAllBuyer()

		assert.Equal(t, []model.Buyer{}, buyers)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

}

func TestGetBuyerByID(t *testing.T) {
	t.Run("return buyer by id existing successfully", func(t *testing.T) {
		svc := setup()

		expectedBuyer := model.Buyer{Id: 2, FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}

		mockRepo := svc.Rp.(*repository.MockBuyerRepo)
		mockRepo.On("GetById", 2).Return(expectedBuyer, nil)

		buyer, err := svc.GetBuyerByID(2)

		assert.Equal(t, expectedBuyer, buyer)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Return buyer not Found", func(t *testing.T) {
		svc := setup()

		expectedError := custom_error.NewBuyerError(http.StatusNotFound, custom_error.ErrNotFound.Error(), "Buyer")

		mockRepo := svc.Rp.(*repository.MockBuyerRepo)
		mockRepo.On("GetById", 99).Return(model.Buyer{}, expectedError)

		buyer, err := svc.GetBuyerByID(99)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.ErrorIs(t, err, expectedError)
		mockRepo.AssertExpectations(t)

	})

	t.Run("return an error when fetching buyer", func(t *testing.T) {
		svc := setup()

		mockRepo := svc.Rp.(*repository.MockBuyerRepo)
		mockRepo.On("GetById", 2).Return(model.Buyer{}, errors.New("unmapped error"))

		buyer, err := svc.GetBuyerByID(2)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
