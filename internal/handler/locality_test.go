package handler_test

import (
	"bytes"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func setupLocality(t *testing.T) *handler.LocalitiesController {
	mock := mocks.NewMockILocalityService(t)
	hd := handler.CreateHandlerLocality(mock)
	return hd
}

var (
	url = "/api/v1/localities/"
)

func TestLocalitiesController_CreateLocality(t *testing.T) {
	t.Run("test handler method for create locality successfully", func(t *testing.T) {
		hd := setupLocality(t)
		mock := hd.Service.(*mocks.MockILocalityService)

		arg := model.Locality{Locality: "Brooklyn", Province: "New York", Country: "EUA"}
		returnService := model.Locality{ID: 1, Locality: "Brooklyn", Province: "New York", Country: "EUA"}
		body := []byte(`{           
						"locality_name": "Brooklyn",
						"province_name": "New York",
						"country_name": "EUA"
					}`)
		res := `{
					"data": {
						"id": 1,
						"locality_name": "Brooklyn",
						"province_name": "New York",
						"country_name": "EUA"
					}
				}`
		statusCode := http.StatusCreated
		mock.On("CreateLocality", &arg).Return(returnService, nil)

		request := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateLocality(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for create locality with bad request", func(t *testing.T) {
		hd := setupLocality(t)
		body := []byte(`{           
						"locality_name": 999,
						"province_name": 999,
						"country_name": 999
					}`)
		res := `{"message":"invalid JSON format in the request body"}`
		statusCode := http.StatusBadRequest

		request := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateLocality(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})

	t.Run("test handler method for create locality with empty attributes values", func(t *testing.T) {
		hd := setupLocality(t)
		mock := hd.Service.(*mocks.MockILocalityService)

		arg := model.Locality{Locality: "", Province: "", Country: ""}
		returnService := model.Locality{}
		body := []byte(`{           
						"locality_name": "",
						"province_name": "",
						"country_name": ""
					}`)
		res := `{"message":"invalid request body, received empty or null value"}`
		statusCode := http.StatusUnprocessableEntity
		errS := customerror.ErrNullLocalityAttribute

		mock.On("CreateLocality", &arg).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateLocality(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})
}

func TestLocalitiesController_GetByID(t *testing.T) {
	hd := setupLocality(t)
	mock := hd.Service.(*mocks.MockILocalityService)

	r := chi.NewRouter()
	r.Get("/api/v1/localities/{id}", hd.GetByID)

	t.Run("test handler method for get locality by ID successfully", func(t *testing.T) {
		returnService := model.Locality{ID: 3, Locality: "Phoenix", Province: "Arizona", Country: "EUA"}
		ID := 3
		res := `{
					"data": {
						"id": 3,
						"locality_name": "Phoenix",
						"province_name": "Arizona",
						"country_name": "EUA"
					}
				}`
		statusCode := http.StatusOK

		mock.On("GetByID", ID).Return(returnService, nil).Once()

		request := httptest.NewRequest(http.MethodGet, url+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for get locality with zero id", func(t *testing.T) {
		ID := 0
		res := `{"message":"missing 'id' parameter in the request"}`
		statusCode := http.StatusBadRequest

		request := httptest.NewRequest(http.MethodGet, url+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})

	t.Run("test handler method for get locality with id not found", func(t *testing.T) {
		returnService := model.Locality{}
		ID := 999
		res := `{"message":"locality not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrLocalityNotFound

		mock.On("GetByID", ID).Return(returnService, errS).Once()

		request := httptest.NewRequest(http.MethodGet, url+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})
}

func TestLocalitiesController_GetSellers(t *testing.T) {
	hd := setupLocality(t)
	mock := hd.Service.(*mocks.MockILocalityService)

	r := chi.NewRouter()
	r.Get("/api/v1/localities/reportSellers", hd.GetSellers)

	t.Run("test handler method for get report seller by ID successfully", func(t *testing.T) {
		returnService := []model.LocalitiesJSONSellers{
			{ID: "5", Locality: "Phoenix", Sellers: 15},
		}
		ID := 5
		res := `{
					"data":[
						{
							"locality_id":"5",
							"locality_name":"Phoenix",
							"sellers_count": 15
						}
					]
				}`
		statusCode := http.StatusOK

		mock.On("GetSellers", ID).Return(returnService, nil).Once()

		url := "/api/v1/localities/reportSellers?id="
		request := httptest.NewRequest(http.MethodGet, url+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for get report seller with zero id", func(t *testing.T) {
		ID := 0
		res := `{"message":"invalid value for request path parameter"}`
		statusCode := http.StatusBadRequest

		url := "/api/v1/localities/reportSellers?id="
		request := httptest.NewRequest(http.MethodGet, url+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for report sellers when service returns error", func(t *testing.T) {
		ID := 5
		res := `{"message":"unmapped locality handler error"}`
		statusCode := http.StatusInternalServerError

		mock.On("GetSellers", ID).Return(nil, errors.New("service error")).Once()

		url := "/api/v1/localities/reportSellers?id=" + strconv.Itoa(ID)
		request := httptest.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for report sellers with locality not found", func(t *testing.T) {
		res := `{"message":"locality not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrLocalityNotFound

		url := "/api/v1/localities/reportSellers?id=999"

		mock.On("GetSellers", 999).Return(nil, errS).Once()

		request := httptest.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})

	t.Run("test handler method for report sellers missing id parameter", func(t *testing.T) {
		res := `{"message":"missing 'id' parameter in the request"}`
		statusCode := http.StatusBadRequest

		url := "/api/v1/localities/reportSellers?id="

		request := httptest.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})
}

func TestLocalitiesController_GetCarriers(t *testing.T) {
	hd := setupLocality(t)
	mock := hd.Service.(*mocks.MockILocalityService)

	r := chi.NewRouter()
	r.Get("/api/v1/localities/reportCarriers", hd.GetCarriers)

	t.Run("test handler method for get carrier by ID successfully", func(t *testing.T) {
		returnService := []model.LocalitiesJSONCarriers{
			{ID: "5", Locality: "Phoenix", Carriers: 15},
		}
		ID := 5

		res := `{
			"data": [
				{
					"locality_id": "5",
					"locality_name": "Phoenix",
					"carriers_count": 15
				}
			]
		}`
		statusCode := http.StatusOK

		mock.On("GetCarriers", ID).Return(returnService, nil).Once()

		url := "/api/v1/localities/reportCarriers?id="
		request := httptest.NewRequest(http.MethodGet, url+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for get report carrier with zero id", func(t *testing.T) {
		ID := 0
		res := `{"message":"invalid value for request path parameter"}`
		statusCode := http.StatusBadRequest

		url := "/api/v1/localities/reportCarriers?id="
		request := httptest.NewRequest(http.MethodGet, url+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for report carriers when service returns error", func(t *testing.T) {
		ID := 5
		res := `{"message":"unmapped locality handler error"}`
		statusCode := http.StatusInternalServerError

		mock.On("GetCarriers", ID).Return(nil, errors.New("service error")).Once()

		url := "/api/v1/localities/reportCarriers?id=" + strconv.Itoa(ID)
		request := httptest.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for report carriers with locality not found", func(t *testing.T) {
		res := `{"message":"locality not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrLocalityNotFound

		url := "/api/v1/localities/reportCarriers?id=999"

		mock.On("GetCarriers", 999).Return(nil, errS).Once()

		request := httptest.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})

	t.Run("test handler method for report carriers missing id parameter", func(t *testing.T) {
		res := `{"message":"missing 'id' parameter in the request"}`
		statusCode := http.StatusBadRequest

		url := "/api/v1/localities/reportCarriers?id="

		request := httptest.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})
}
