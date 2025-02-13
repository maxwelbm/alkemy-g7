package handler_test

import (
	"bytes"
	"errors"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) *handler.BuyerHandler {
	mockBuyerService := mocks.NewMockIBuyerservice(t)
	hd := handler.NewBuyerHandler(mockBuyerService)
	return hd
}

func TestHandlerGetBuyerById(t *testing.T) {
	t.Run("return buyer by id existing successfully", func(t *testing.T) {
		hd := setup(t)

		expectedBuyer := model.Buyer{ID: 2, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("GetBuyerByID", 2).Return(expectedBuyer, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/2", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerByID(response, request)

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
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("GetBuyerByID", 99).Return(model.Buyer{}, customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/99", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerByID(response, request)

		expectedJson := `{"message":"Buyer not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Error return due to invalid ID", func(t *testing.T) {
		hd := setup(t)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/aaa", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerByID(response, request)

		expectedJson := `{"message":"Invalid ID"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

	t.Run("return an error when fetching buyer", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("GetBuyerByID", 2).Return(model.Buyer{}, errors.New("Unmapped error"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/2", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetBuyerByID(response, request)

		expectedJson := `{"message":"Unable to search for buyer"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

}

func TestHandlerCreateBuyer(t *testing.T) {
	t.Run("Buyer created successfully", func(t *testing.T) {
		hd := setup(t)

		createdBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CreateBuyer", model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}).
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
		hd := setup(t)
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
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CreateBuyer", model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}).
			Return(model.Buyer{}, customerror.NewBuyerError(http.StatusConflict, customerror.ErrConflict.Error(), "card_number_id"))

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
		hd := setup(t)
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
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CreateBuyer", model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}).Return(model.Buyer{}, errors.New("Unmapped error"))

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

func TestHandlerUpdateBuyer(t *testing.T) {
	t.Run("Buyer Updated successfully", func(t *testing.T) {
		hd := setup(t)

		UpdatedBuyer := model.Buyer{ID: 1, FirstName: "Abilio", LastName: "Milan", CardNumberID: "4321"}
		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("UpdateBuyer", 1, model.Buyer{FirstName: "Abilio"}).Return(UpdatedBuyer, nil)

		body := []byte(`{           
           
			"first_name": "Abilio"
			
	}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/1", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerUpdateBuyer(response, request)

		expectedJson := `{
    "data": 
        {
            "id": 1,
            "card_number_id": "4321",
            "first_name": "Abilio",
            "last_name": "Milan"
        }
}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Buyer not Found", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("UpdateBuyer", 99, model.Buyer{FirstName: "Jonas"}).
			Return(model.Buyer{}, customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer"))

		body := []byte(`{           
           
			"first_name": "Jonas"
			
	}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/99", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerUpdateBuyer(response, request)

		expectedJson := `{"message":"Buyer not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Return error card_number already exists", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("UpdateBuyer", 1, model.Buyer{CardNumberID: "1234"}).
			Return(model.Buyer{}, customerror.NewBuyerError(http.StatusConflict, customerror.ErrConflict.Error(), "card_number_id"))

		body := []byte(`{           
           
            "card_number_id": "1234"
}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/1", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerUpdateBuyer(response, request)

		expectedJson := `{
     "message": "card_number_id it already exists"
}`

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Buyer does not have the required fields", func(t *testing.T) {
		hd := setup(t)

		body := []byte(`{           
           
		"card_number_id": "",
		"first_name": "",
		"last_name": ""
}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/1", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerUpdateBuyer(response, request)

		expectedJson := `{
    "message": "at least one field must be filled in"
}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("Return error Json Syntax", func(t *testing.T) {
		hd := setup(t)
		body := []byte(`{           
		"last_name": "",
		FieldInexistingInBuyer
}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/1", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerUpdateBuyer(response, request)

		expectedJson := `{
    "message": "JSON syntax error. Please verify your input."
}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("return an error when updated buyer", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("UpdateBuyer", 1, model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}).Return(model.Buyer{}, errors.New("Unmapped error"))

		body := []byte(`{           
           
		"card_number_id": "4321",
		"first_name": "Ac",
		"last_name": "Milan"
}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/1", bytes.NewReader(body))
		response := httptest.NewRecorder()

		hd.HandlerUpdateBuyer(response, request)

		expectedJson := `{"message":"Unable to update buyer"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

	t.Run("Error return due to invalid ID", func(t *testing.T) {
		hd := setup(t)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/aaa", nil)
		response := httptest.NewRecorder()

		hd.HandlerUpdateBuyer(response, request)

		expectedJson := `{"message":"Invalid ID"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

}

func TestHandlerDeleteBuyerById(t *testing.T) {
	t.Run("Deleted Buyer successfly", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("DeleteBuyerByID", 1).Return(nil)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/buyers/1", nil)
		response := httptest.NewRecorder()

		hd.HandlerDeleteBuyerByID(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
		mockSvc.AssertExpectations(t)

	})

	t.Run("Buyer not Found", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("DeleteBuyerByID", 99).Return(customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer"))

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/buyers/99", nil)
		response := httptest.NewRecorder()

		hd.HandlerDeleteBuyerByID(response, request)

		expectedJson := `{"message":"Buyer not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("There are dependencies with the buyer", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("DeleteBuyerByID", 1).Return(customerror.NewBuyerError(http.StatusConflict, customerror.ErrDependencies.Error(), "Buyer"))

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/buyers/1", nil)
		response := httptest.NewRecorder()

		hd.HandlerDeleteBuyerByID(response, request)

		expectedJson := `{"message":"Buyer cannot be deleted because there are dependencies"}`

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

	t.Run("Invalid ID parameter", func(t *testing.T) {
		hd := setup(t)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/buyers/aaaa", nil)
		response := httptest.NewRecorder()

		hd.HandlerDeleteBuyerByID(response, request)

		expectedJson := `{"message":"Invalid ID"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())

	})

	t.Run("return an error when deleted buyer", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("DeleteBuyerByID", 1).Return(errors.New("Unmapped error"))

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/buyers/1", nil)
		response := httptest.NewRecorder()

		hd.HandlerDeleteBuyerByID(response, request)

		expectedJson := `{"message":"Unable to delete buyer"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})

}

func TestHandlerGetBuyers(t *testing.T) {

	t.Run("return a list of all existing buyers successfully", func(t *testing.T) {

		hd := setup(t)

		expectedBuyers := []model.Buyer{{ID: 1, FirstName: "John", LastName: "Doe", CardNumberID: "1234"},
			{ID: 2, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}}

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", nil)
		response := httptest.NewRecorder()
		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("GetAllBuyer").Return(expectedBuyers, nil)

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
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("GetAllBuyer").Return([]model.Buyer{}, errors.New("Unmapped error"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", nil)
		response := httptest.NewRecorder()

		hd.HandlerGetAllBuyers(response, request)

		expectedJson := `{"message":"Unable to list Buyers"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)

	})
}

func TestHandlerCountPurchaseOrderBuyer(t *testing.T) {
	t.Run("return a count of buyers successfully", func(t *testing.T) {
		hd := setup(t)

		countBuyers := []model.BuyerPurchaseOrder{model.BuyerPurchaseOrder{
			ID:                  1,
			CardNumberID:        "cn001",
			FirstName:           "Jhon",
			LastName:            "DOe",
			PurchaseOrdersCount: 12,
		}}

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CountPurchaseOrderBuyer").Return(countBuyers, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders", nil)
		response := httptest.NewRecorder()

		hd.HandlerCountPurchaseOrderBuyer(response, request)

		expectedJson := `{
    "data": [
        {
            "id": 1,
            "card_number_id": "cn001",
            "first_name": "Jhon",
            "last_name": "DOe",
            "purchase_orders_count": 12
        }
    ]
}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

	t.Run("return a count of buyer successfully", func(t *testing.T) {
		hd := setup(t)

		countBuyer := model.BuyerPurchaseOrder{
			ID:                  1,
			CardNumberID:        "cn001",
			FirstName:           "Jhon",
			LastName:            "DOe",
			PurchaseOrdersCount: 12,
		}

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CountPurchaseOrderByBuyerID", countBuyer.ID).Return(countBuyer, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?id=1", nil)
		response := httptest.NewRecorder()

		hd.HandlerCountPurchaseOrderBuyer(response, request)

		expectedJson := `{
    "data":
        {
            "id": 1,
            "card_number_id": "cn001",
            "first_name": "Jhon",
            "last_name": "DOe",
            "purchase_orders_count": 12
        }
}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

	t.Run("Error invalid ID", func(t *testing.T) {
		hd := setup(t)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?id=aaa", nil)
		response := httptest.NewRecorder()

		hd.HandlerCountPurchaseOrderBuyer(response, request)

		expectedJson := `{
    "message":"Invalid ID"
}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("Buyer Not Found", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CountPurchaseOrderByBuyerID", 99).Return(model.BuyerPurchaseOrder{}, customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?id=99", nil)
		response := httptest.NewRecorder()

		hd.HandlerCountPurchaseOrderBuyer(response, request)

		expectedJson := `{
    "message":"Buyer not found"
}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

	t.Run("Unmapped Error in CountPurchaseOrders", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CountPurchaseOrderBuyer").Return([]model.BuyerPurchaseOrder{}, errors.New("Unmapped error"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders", nil)
		response := httptest.NewRecorder()

		hd.HandlerCountPurchaseOrderBuyer(response, request)

		expectedJson := `{
    "message":"Unable to count buyer Purchase orders"
}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})

	t.Run("Unmapped Error in CountPurchaseOrderByID", func(t *testing.T) {
		hd := setup(t)

		mockSvc := hd.Svc.(*mocks.MockIBuyerservice)
		mockSvc.On("CountPurchaseOrderByBuyerID", 1).Return(model.BuyerPurchaseOrder{}, errors.New("Unmapped error"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?id=1", nil)
		response := httptest.NewRecorder()

		hd.HandlerCountPurchaseOrderBuyer(response, request)

		expectedJson := `{
    "message":"Unable to count buyer Purchase orders"
}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockSvc.AssertExpectations(t)
	})
}
