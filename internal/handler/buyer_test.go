package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/stretchr/testify/assert"
)

func setup() *handler.BuyerHandler {
	mockBuyerService := new(service.MockBuyerService)
	hd := handler.NewBuyerHandler(mockBuyerService)
	return hd
}

func TestHandlerGetAllBuyers(t *testing.T) {
	t.Run("return a list of all existing buyers successfully", func(t *testing.T) {
		hd := setup()

		expectedBuyers := []model.Buyer{{Id: 1, FirstName: "John", LastName: "Doe", CardNumberId: "1234"},
			{Id: 2, FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}}

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("GetAllBuyer").Return(expectedBuyers, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetAllBuyers(response, request)

		expectedJSON := `{
    "data": [
        {
            "id": 1,
            "card_number_id": "1234",
            "first_name": "John",
            "last_name": "Doe"
        },
        {
            "id": 2,
            "card_number_id": "4321",
            "first_name": "Ac",
            "last_name": "Milan"
        }
    ]
}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJSON, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("return an error when fetching buyers", func(t *testing.T) {
		hd := setup()

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("GetAllBuyer").Return([]model.Buyer{}, errors.New("Unmapped error"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetAllBuyers(response, request)

		expectedJson := `{"message":"Unable to list Buyers"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})
}
