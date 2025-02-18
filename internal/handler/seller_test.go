package handler_test

import (
	"bytes"
	"errors"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

var logMock = mocks.MockLog{}

func setupSeller(t *testing.T) *handler.SellersController {
	mock := mocks.NewMockISellerService(t)
	hd := handler.CreateHandlerSellers(mock, logMock)
	return hd
}

var (
	endpoint = "/api/v1/sellers/"
)

func TestSellersController_GetAllSellers(t *testing.T) {
	t.Run("test handler method for get a list of all existing sellers successfully", func(t *testing.T) {
		hd := setupSeller(t)
		mock := hd.Service.(*mocks.MockISellerService)

		returnService := []model.Seller{{ID: 1, CID: 1, CompanyName: "Enterprise Liberty", Address: "456 Elm St", Telephone: "4443335454", Locality: 1},
			{ID: 2, CID: 2, CompanyName: "Libre Mercado", Address: "123 Montain St Avenue", Telephone: "5554545999", Locality: 2}}
		res := `{
					"data": [{
						"id": 1,
						"cid": 1,
						"company_name": "Enterprise Liberty",
						"address": "456 Elm St",
						"telephone": "4443335454",
						"locality_id": 1
					},
					{
						"id": 2,
						"cid": 2,
						"company_name": "Libre Mercado",
						"address": "123 Montain St Avenue",
						"telephone": "5554545999",
						"locality_id": 2
					}]
				}`
		statusCode := http.StatusOK

		mock.On("GetAll").Return(returnService, nil)

		request := httptest.NewRequest(http.MethodGet, endpoint, nil)
		response := httptest.NewRecorder()
		hd.GetAllSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for get all sellers with internal server error", func(t *testing.T) {
		hd := setupSeller(t)
		mock := hd.Service.(*mocks.MockISellerService)

		returnService := []model.Seller{}
		res := `{"message":"unmapped seller handler error"}`
		statusCode := http.StatusInternalServerError
		er := errors.New("internal server error")

		mock.On("GetAll").Return(returnService, er)

		request := httptest.NewRequest(http.MethodGet, endpoint, nil)
		response := httptest.NewRecorder()
		hd.GetAllSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})
}

func TestSellersController_GetByID(t *testing.T) {
	hd := setupSeller(t)
	mock := hd.Service.(*mocks.MockISellerService)

	r := chi.NewRouter()
	r.Get("/api/v1/sellers/{id}", hd.GetByID)

	t.Run("test handler service for get seller by ID with success", func(t *testing.T) {
		returnService := model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 3}
		ID := 3
		res := `{
					"data": {
						"id": 3,
						"cid": 3,
						"company_name": "Enterprise Science",
						"address": "1200 Central Park Avenue",
						"telephone": "999444555",
						"locality_id": 3
					}
				}`
		statusCode := http.StatusOK

		mock.On("GetByID", ID).Return(returnService, nil)

		request := httptest.NewRequest(http.MethodGet, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler service for get seller with ID not found", func(t *testing.T) {
		returnService := model.Seller{}
		ID := 999
		res := `{"message":"seller not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrSellerNotFound

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodGet, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler service for get seller with internal server error", func(t *testing.T) {
		returnService := model.Seller{}
		ID := 4
		res := `{"message":"internal server error"}`
		statusCode := http.StatusInternalServerError
		errS := customerror.ErrDefaultSeller

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodGet, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler service for get seller with zero id", func(t *testing.T) {
		ID := 0
		res := `{"message":"missing 'id' parameter in the request"}`
		statusCode := http.StatusBadRequest

		request := httptest.NewRequest(http.MethodGet, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})
}

func TestSellersController_CreateSellers(t *testing.T) {
	t.Run("test handler method for create seller with success", func(t *testing.T) {
		hd := setupSeller(t)
		mock := hd.Service.(*mocks.MockISellerService)

		arg := model.Seller{CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5}
		returnService := model.Seller{ID: 5, CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5}
		body := []byte(`{           
						"cid": 5,
						"company_name": "Enterprise Cypress",
						"address": "702 St Mark",
						"telephone": "33344455566",
						"locality_id": 5
					}`)
		res := `{
                        "data": {
                            "id": 5,
                            "cid": 5,
                            "company_name": "Enterprise Cypress",
                            "address": "702 St Mark",
                            "telephone": "33344455566",
                            "locality_id": 5
                        }
                    }`
		statusCode := http.StatusCreated
		mock.On("CreateSeller", &arg).Return(returnService, nil)

		request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for create seller with bad request", func(t *testing.T) {
		hd := setupSeller(t)
		body := []byte(`{           
						"cid": "cid",
						"company_name": 9999,
						"address": 9999,
						"telephone": 9999,
						"locality_id": "locality"
					}`)
		res := `{"message":"invalid JSON format in the request body"}`
		statusCode := http.StatusBadRequest

		request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})

	t.Run("test handler method for create seller with empty attributes values", func(t *testing.T) {
		hd := setupSeller(t)
		mock := hd.Service.(*mocks.MockISellerService)

		arg := model.Seller{CID: 0, CompanyName: "", Address: "", Telephone: "", Locality: 0}
		returnService := model.Seller{}
		body := []byte(`{           
							"cid": 0,
							"company_name": "",
							"address": "",
							"telephone": "",
							"locality_id": 0
						}`)
		res := `{"message":"invalid request body, received empty or null value"}`
		statusCode := http.StatusUnprocessableEntity
		errS := customerror.ErrNullSellerAttribute

		mock.On("CreateSeller", &arg).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for create seller without required attributes", func(t *testing.T) {
		hd := setupSeller(t)
		mock := hd.Service.(*mocks.MockISellerService)

		arg := model.Seller{}
		returnService := model.Seller{}
		body := []byte(`{}`)
		res := `{"message":"invalid request body, received empty or null value"}`
		statusCode := http.StatusUnprocessableEntity
		errS := customerror.ErrNullSellerAttribute

		mock.On("CreateSeller", &arg).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for create seller with CID attribute already existing", func(t *testing.T) {
		hd := setupSeller(t)
		mock := hd.Service.(*mocks.MockISellerService)

		arg := model.Seller{CID: 1, CompanyName: "Midgard Sellers", Address: "3 New Time Park", Telephone: "99989898778", Locality: 7}
		returnService := model.Seller{}
		body := []byte(`{           
							"cid": 1,
							"company_name": "Midgard Sellers",
							"address": "3 New Time Park",
							"telephone": "99989898778",
							"locality_id": 7
						}`)
		res := `{"message":"seller's CID already exists"}`
		statusCode := http.StatusConflict
		errS := customerror.ErrCIDSellerAlreadyExist

		mock.On("CreateSeller", &arg).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for create seller with locality attribute not found", func(t *testing.T) {
		hd := setupSeller(t)
		mock := hd.Service.(*mocks.MockISellerService)

		arg := model.Seller{CID: 8, CompanyName: "Rupture Clivers", Address: "1200 New Time Park", Telephone: "7776657987", Locality: 9999}
		returnService := model.Seller{}
		body := []byte(`{           
							"cid": 8,
							"company_name": "Rupture Clivers",
							"address": "1200 New Time Park",
							"telephone": "7776657987",
							"locality_id": 9999
						}`)
		res := `{"message":"locality not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrLocalityNotFound

		mock.On("CreateSeller", &arg).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateSellers(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})
}

func TestSellersController_UpdateSellers(t *testing.T) {
	hd := setupSeller(t)
	mock := hd.Service.(*mocks.MockISellerService)

	r := chi.NewRouter()
	r.Patch("/api/v1/sellers/{id}", hd.UpdateSellers)

	t.Run("test handler method for update seller with success", func(t *testing.T) {
		arg := model.Seller{CID: 55, CompanyName: "Cypress Company", Address: "900 Central Park", Telephone: "55566777787", Locality: 10}
		ID := 5
		returnService := model.Seller{ID: 5, CID: 55, CompanyName: "Cypress Company", Address: "900 Central Park", Telephone: "55566777787", Locality: 10}
		body := []byte(`{           
							"cid": 55,
							"company_name": "Cypress Company",
							"address": "900 Central Park",
							"telephone": "55566777787",
							"locality_id": 10
						}`)
		res := `{
                        "data": {
                            "id": 5,
                            "cid": 55,
                            "company_name": "Cypress Company",
                            "address": "900 Central Park",
                            "telephone": "55566777787",
                            "locality_id": 10
                        }
                    }`
		statusCode := http.StatusOK

		mock.On("UpdateSeller", ID, &arg).Return(returnService, nil)
		mock.On("GetByID", ID).Return(returnService, nil)

		request := httptest.NewRequest(http.MethodPatch, endpoint+strconv.Itoa(ID), bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for update seller with ID not found", func(t *testing.T) {
		ID := 999
		returnService := model.Seller{}
		body := []byte(`{           
							"cid": 65,
							"company_name": "Cypress Company",
							"address": "30 Central Park",
							"telephone": "55566777787",
							"locality_id": 20
						}`)
		res := `{"message":"seller not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrSellerNotFound

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPatch, endpoint+strconv.Itoa(ID), bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for update seller with bad request", func(t *testing.T) {
		ID := 4
		returnService := model.Seller{}
		body := []byte(`{           
							"cid": "cid",
							"company_name": 9999,
							"address": 9999,
							"telephone": 9999,
							"locality_id": "locality"
						}`)
		res := `{"message":"invalid JSON format in the request body"}`
		statusCode := http.StatusBadRequest
		errS := customerror.ErrInvalidSellerJSONFormat

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPatch, endpoint+strconv.Itoa(ID), bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for update seller with empty attributes values", func(t *testing.T) {
		ID := 2
		returnService := model.Seller{}
		body := []byte(`{           
							"cid": 0,
							"company_name": "",
							"address": "",
							"telephone": "",
							"locality_id": 0
						}`)
		res := `{"message":"invalid request body, received empty or null value"}`
		statusCode := http.StatusUnprocessableEntity
		errS := customerror.ErrNullSellerAttribute

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPatch, endpoint+strconv.Itoa(ID), bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for update seller with CID attribute already existing", func(t *testing.T) {
		ID := 9
		returnService := model.Seller{}
		body := []byte(`{           
							"cid": 1,
							"company_name": "Midgard Sellers",
							"address": "3 New Time Park",
							"telephone": "99989898778",
							"locality_id": 17
						}`)
		res := `{"message":"seller's CID already exists"}`
		statusCode := http.StatusConflict
		errS := customerror.ErrCIDSellerAlreadyExist

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPatch, endpoint+strconv.Itoa(ID), bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for update seller with locality attribute not found", func(t *testing.T) {
		returnService := model.Seller{}
		ID := 7
		body := []byte(`{           
							"cid": 8,
							"company_name": "Rupture Clivers",
							"address": "1200 New Time Park",
							"telephone": "7776657987",
							"locality_id": 9999
						}`)
		res := `{"message":"locality not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrLocalityNotFound

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodPatch, endpoint+strconv.Itoa(ID), bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for update seller with zero id", func(t *testing.T) {
		ID := 0
		body := []byte(`{           
							"cid": 55,
							"company_name": "Cypress Company",
							"address": "400 Central Park",
							"telephone": "55566777787",
							"locality_id": 30
						}`)
		res := `{"message":"missing 'id' parameter in the request"}`
		statusCode := http.StatusBadRequest

		request := httptest.NewRequest(http.MethodPatch, endpoint+strconv.Itoa(ID), bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})
}

func TestSellersController_DeleteSellers(t *testing.T) {
	hd := setupSeller(t)
	mock := hd.Service.(*mocks.MockISellerService)

	r := chi.NewRouter()
	r.Delete("/api/v1/sellers/{id}", hd.DeleteSellers)

	t.Run("test handler method for delete seller with success", func(t *testing.T) {
		ID := 3
		returnService := model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 3}
		res := `{}`
		statusCode := http.StatusNoContent

		mock.On("DeleteSeller", ID).Return(nil)
		mock.On("GetByID", ID).Return(returnService, nil)

		request := httptest.NewRequest(http.MethodDelete, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for delete seller with ID not found", func(t *testing.T) {
		ID := 999
		returnService := model.Seller{}
		res := `{"message":"seller not found"}`
		statusCode := http.StatusNotFound
		errS := customerror.ErrSellerNotFound

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodDelete, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for delete seller with internal server error", func(t *testing.T) {
		ID := 4
		returnService := model.Seller{}
		res := `{"message":"internal server error"}`
		statusCode := http.StatusInternalServerError
		errS := customerror.ErrDefaultSeller

		mock.On("GetByID", ID).Return(returnService, errS)

		request := httptest.NewRequest(http.MethodDelete, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
		mock.AssertExpectations(t)
	})

	t.Run("test handler method for delete seller with zero id", func(t *testing.T) {
		ID := 0
		res := `{"message":"missing 'id' parameter in the request"}`
		statusCode := http.StatusBadRequest

		request := httptest.NewRequest(http.MethodDelete, endpoint+strconv.Itoa(ID), nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		assert.Equal(t, statusCode, response.Code)
		assert.JSONEq(t, res, response.Body.String())
	})
}
