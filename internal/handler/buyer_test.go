package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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

func TestHandlerGetBuyerById(t *testing.T) {
	t.Run("return buyer by id existing successfully", func(t *testing.T) {
		hd := setup()

		expectedBuyer := model.Buyer{Id: 2, FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("GetBuyerByID", 2).Return(expectedBuyer, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/2", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerById(response, request)

		expectedJson := `{
    "data": 
        {
            "id": 2,
            "card_number_id": "4321",
            "first_name": "Ac",
            "last_name": "Milan"
        }
}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Return buyer not Found", func(t *testing.T) {
		hd := setup()

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("GetBuyerByID", 99).Return(model.Buyer{}, custom_error.NewBuyerError(http.StatusNotFound, custom_error.ErrNotFound.Error(), "Buyer"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/99", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerById(response, request)

		expectedJson := `{"message":"Buyer not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Error return due to invalid ID", func(t *testing.T) {
		hd := setup()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/aaa", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerById(response, request)

		expectedJson := `{"message":"Invalid ID"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

	t.Run("return an error when fetching buyer", func(t *testing.T) {
		hd := setup()

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("GetBuyerByID", 2).Return(model.Buyer{}, errors.New("Unmapped error"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/2", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerById(response, request)

		expectedJson := `{"message":"Unable to search for buyer"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

}

func TestHandlerCreateBuyer(t *testing.T) {
	t.Run("Buyer created successfully", func(t *testing.T) {
		hd := setup()

		createdBuyer := model.Buyer{Id: 1, FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("CreateBuyer", model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}).
			Return(createdBuyer, nil)

		body := []byte(`{           
           
            "card_number_id": "4321",
            "first_name": "Ac",
            "last_name": "Milan"
}`)
		request := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerCreateBuyer(response, request)

		expectedJson := `{
    "data": 
        {
            "id": 1,
            "card_number_id": "4321",
            "first_name": "Ac",
            "last_name": "Milan"
        }
}`

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Buyer does not have the required fields", func(t *testing.T) {
		hd := setup()
		body := []byte(`{           
           
		"card_number_id": "",
		"first_name": "",
		"last_name": ""
}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerCreateBuyer(response, request)

		expectedJson := `{
    "message": "field(s) card_number_id, first_name, last_name cannot be empty"
}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("Return error card_number already exists", func(t *testing.T) {
		hd := setup()

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("CreateBuyer", model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}).
			Return(model.Buyer{}, custom_error.NewBuyerError(http.StatusConflict, custom_error.ErrConflict.Error(), "card_number_id"))

		body := []byte(`{           
           
            "card_number_id": "4321",
            "first_name": "Ac",
            "last_name": "Milan"
}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerCreateBuyer(response, request)

		expectedJson := `{
     "message": "card_number_id it already exists"
}`

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Return error Json Syntax", func(t *testing.T) {
		hd := setup()
		body := []byte(`{           
		"last_name": "",
		FieldInexistingInBuyer
}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerCreateBuyer(response, request)

		expectedJson := `{
    "message": "JSON syntax error. Please verify your input."
}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("return an error when created buyer", func(t *testing.T) {
		hd := setup()

		mockSvc := hd.Svc.(*service.MockBuyerService)
		mockSvc.On("CreateBuyer", model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberId: "4321"}).Return(model.Buyer{}, errors.New("Unmapped error"))

		body := []byte(`{           
           
		"card_number_id": "4321",
		"first_name": "Ac",
		"last_name": "Milan"
}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerCreateBuyer(response, request)

		expectedJson := `{"message":"Unable to create buyer"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

}
