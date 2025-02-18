package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setupCarrierHandler(t *testing.T) *handler.CarrierHandler {
	mockServiceCarrier := mocks.NewMockICarrierService(t)
	hd := handler.NewCarrierHandler(mockServiceCarrier)
	return hd
}
func TestHandlerPostCarriers(t *testing.T) {
	hd := setupCarrierHandler(t)
	mockServiceCarrier := hd.Srv.(*mocks.MockICarrierService)

	r := chi.NewRouter()
	r.Post("/api/v1/carriers", hd.PostCarriers())

	t.Run("PostCarrier create success", func(t *testing.T) {
		carrier := model.Carries{
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}

		mockServiceCarrier.On("PostCarrier", carrier).Return(model.Carries{
			ID:          1,
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}, nil)

		reqBody := []byte(`{
			"cid": "CID001",
			"company_name": "ABC Company",
			"address": "123 Main St",
			"telephone": "1234567890",
			"locality_id": 1
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)

		expectedJson := `{
			"data": {
				"id": 1,
				"cid": "CID001",
				"company_name": "ABC Company",
				"address": "123 Main St",
				"telephone": "1234567890",
				"locality_id": 1
			}
		}`

		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceCarrier.AssertExpectations(t)
	})

	t.Run("PostCarrier required fields not found", func(t *testing.T) {
		hd := setupCarrierHandler(t)

		mockServiceCarrier := hd.Srv.(*mocks.MockICarrierService)

		reqBody := []byte(`{}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedJson := `{"message": "the following fields are required: cid, company_name, address, telephone, locality_id"}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockServiceCarrier.AssertExpectations(t)
	})

	t.Run("Conflict - Duplicate CID", func(t *testing.T) {
		hd := setupCarrierHandler(t)
		mockServiceCarrier := hd.Srv.(*mocks.MockICarrierService)

		r := chi.NewRouter()
		r.Post("/api/v1/carriers", hd.PostCarriers())

		carrier := model.Carries{
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}

		mockServiceCarrier.On("PostCarrier", carrier).Return(model.Carries{}, customerror.NewCarrierError(customerror.ErrConflict.Error(), "cid", http.StatusConflict))

		reqBody := []byte(`{
			"cid": "CID001",
			"company_name": "ABC Company",
			"address": "123 Main St",
			"telephone": "1234567890",
			"locality_id": 1
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		expectedJson := `{"message":"cid, it already exists"}`

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("Not Found - Locality Not Found", func(t *testing.T) {

		carrier := model.Carries{
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  99,
		}

		mockServiceCarrier.On("PostCarrier", carrier).Return(model.Carries{}, customerror.NewCarrierError(customerror.ErrLocalityNotFound.Error(), "locality", http.StatusNotFound))

		reqBody := []byte(`{
			"cid": "CID001",
			"company_name": "ABC Company",
			"address": "123 Main St",
			"telephone": "1234567890",
			"locality_id": 99
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)

		expectedJson := `{"message":"locality, locality not found"}`

		assert.JSONEq(t, expectedJson, response.Body.String())
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		hd := setupCarrierHandler(t)
		mockServiceCarrier := hd.Srv.(*mocks.MockICarrierService)

		r := chi.NewRouter()
		r.Post("/api/v1/carriers", hd.PostCarriers())
		carrier := model.Carries{
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}

		mockServiceCarrier.On("PostCarrier", carrier).Return(model.Carries{}, errors.New("some unexpected error"))

		reqBody := []byte(`{
			"cid": "CID001",
			"company_name": "ABC Company",
			"address": "123 Main St",
			"telephone": "1234567890",
			"locality_id": 1
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", bytes.NewReader(reqBody))

		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		expectedJson := `{"message":"unable to post carrier"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})
}
