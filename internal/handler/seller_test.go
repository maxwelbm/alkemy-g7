package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
	"github.com/stretchr/testify/assert"
)

func setupSeller() *handler.SellersController {
	mock := new(service.SellerMockService)
	hd := handler.CreateHandlerSellers(mock)
	return hd
}

var (
	endpoint string = "/api/v1/sellers/"
)

func TestHandlerGetAllSeller(t *testing.T) {
	tests := []struct {
		description   string
		returnService []model.Seller
		response      string
		statusCode    int
		err           error
	}{
		{
			description: "get all sellers with success",
			returnService: []model.Seller{{ID: 1, CID: 1, CompanyName: "Enterprise Liberty", Address: "456 Elm St", Telephone: "4443335454", Locality: 1},
				{ID: 2, CID: 2, CompanyName: "Libre Mercado", Address: "123 Montain St Avenue", Telephone: "5554545999", Locality: 2}},
			response: `{
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
                        }`,
			statusCode: http.StatusOK,
			err:        nil,
		},
		{
			description:   "get all sellers with internal server error",
			returnService: []model.Seller{},
			response:      `{"message":"unmapped seller handler error"}`,
			statusCode:    http.StatusInternalServerError,
			err:           errors.New("internal server error"),
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			hd := setupSeller()
			mock := hd.Service.(*service.SellerMockService)
			mock.On("GetAll").Return(test.returnService, test.err)

			request := httptest.NewRequest(http.MethodGet, endpoint, nil)
			response := httptest.NewRecorder()
			hd.GetAllSellers(response, request)

			assert.Equal(t, test.statusCode, response.Code)
			assert.JSONEq(t, test.response, response.Body.String())
			mock.AssertExpectations(t)
		})
	}
}

func TestHandlerGetByIDSeller(t *testing.T) {
	hd := setupSeller()
	mock := hd.Service.(*service.SellerMockService)

	r := chi.NewRouter()
	r.Get("/api/v1/sellers/{id}", hd.GetById)

	tests := []struct {
		description   string
		returnService model.Seller
		id            int
		response      string
		statusCode    int
		err           error
	}{
		{
			description:   "get seller by id with success",
			returnService: model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Perk Avenue", Telephone: "999444555", Locality: 3},
			id:            3,
			response: `{
                            "data": {
                                "id": 3,
                                "cid": 3,
                                "company_name": "Enterprise Science",
                                "address": "1200 Central Perk Avenue",
                                "telephone": "999444555",
                                "locality_id": 3
                            }
                        }`,
			statusCode: http.StatusOK,
			err:        nil,
		},
		{
			description:   "get seller by id not found",
			returnService: model.Seller{},
			id:            999,
			response:      `{"message":"seller not found"}`,
			statusCode:    http.StatusNotFound,
			err:           customError.ErrSellerNotFound,
		},
		{
			description:   "get seller by id with internal server error",
			returnService: model.Seller{},
			id:            4,
			response:      `{"message":"internal server error"}`,
			statusCode:    http.StatusInternalServerError,
			err:           customError.ErrDefaultSeller,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mock.On("GetById", test.id).Return(test.returnService, test.err)

			request := httptest.NewRequest(http.MethodGet, endpoint+strconv.Itoa(test.id), nil)
			response := httptest.NewRecorder()
			r.ServeHTTP(response, request)

			assert.Equal(t, test.statusCode, response.Code)
			assert.JSONEq(t, test.response, response.Body.String())
			mock.AssertExpectations(t)
		})
	}
}

func TestHandlerCreateSeller(t *testing.T) {
	tests := []struct {
		description   string
		arg           model.Seller
		returnService model.Seller
		body          []byte
		response      string
		statusCode    int
		err           error
		call          bool
	}{
		{
			description:   "create seller with success",
			arg:           model.Seller{CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5},
			returnService: model.Seller{ID: 5, CID: 5, CompanyName: "Enterprise Cypress", Address: "702 St Mark", Telephone: "33344455566", Locality: 5},
			body: []byte(`{           
							"cid": 5,
							"company_name": "Enterprise Cypress",
							"address": "702 St Mark",
							"telephone": "33344455566",
							"locality_id": 5
						}`),
			response: `{
                        "data": {
                            "id": 5,
                            "cid": 5,
                            "company_name": "Enterprise Cypress",
                            "address": "702 St Mark",
                            "telephone": "33344455566",
                            "locality_id": 5
                        }
                    }`,
			statusCode: http.StatusCreated,
			err:        nil,
			call:       true,
		},
		{
			description:   "create seller with bad request",
			arg:           model.Seller{},
			returnService: model.Seller{},
			body: []byte(`{           
							"cid": "cid",
							"company_name": 9999,
							"address": 9999,
							"telephone": 9999,
							"locality_id": "locality"
						}`),
			response:   `{"message":"invalid JSON format in the request body"}`,
			statusCode: http.StatusBadRequest,
			err:        customError.ErrInvalidSellerJSONFormat,
			call:       false,
		},
		{
			description:   "create seller with empty attributes values",
			arg:           model.Seller{CID: 0, CompanyName: "", Address: "", Telephone: "", Locality: 0},
			returnService: model.Seller{},
			body: []byte(`{           
							"cid": 0,
							"company_name": "",
							"address": "",
							"telephone": "",
							"locality_id": 0
						}`),
			response:   `{"message":"invalid request body, received empty or null value"}`,
			statusCode: http.StatusUnprocessableEntity,
			err:        customError.ErrNullSellerAttribute,
			call:       false,
		},
		{
			description:   "create seller with no required attributes",
			arg:           model.Seller{},
			returnService: model.Seller{},
			body: []byte(`{}`),
			response:   `{"message":"invalid request body, received empty or null value"}`,
			statusCode: http.StatusUnprocessableEntity,
			err:        customError.ErrNullSellerAttribute,
			call:       false,
		},
		{
			description:   "create seller with attribute cid already existing",
			arg:           model.Seller{CID: 1, CompanyName: "Midgard Sellers", Address: "3 New Time Park", Telephone: "99989898778", Locality: 7},
			returnService: model.Seller{},
			body: []byte(`{           
							"cid": 1,
							"company_name": "Midgard Sellers",
							"address": "3 New Time Park",
							"telephone": "99989898778",
							"locality_id": 7
						}`),
			response:   `{"message":"seller's CID already exists"}`,
			statusCode: http.StatusConflict,
			err:        customError.ErrCIDSellerAlreadyExist,
			call:       true,
		},
		{
			description:   "create seller with attribute locality id inexisting",
			arg:           model.Seller{CID: 8, CompanyName: "Rupture Clivers", Address: "1200 New Time Park", Telephone: "7776657987", Locality: 9999},
			returnService: model.Seller{},
			body: []byte(`{           
							"cid": 8,
							"company_name": "Rupture Clivers",
							"address": "1200 New Time Park",
							"telephone": "7776657987",
							"locality_id": 9999
						}`),
			response:   `{"message":"locality not found"}`,
			statusCode: http.StatusNotFound,
			err:        customError.ErrLocalityNotFound,
			call:       true,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			hd := setupSeller()
			mock := hd.Service.(*service.SellerMockService)
			mock.On("CreateSeller", &test.arg).Return(test.returnService, test.err)

			request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(test.body))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			hd.CreateSellers(response, request)

			assert.Equal(t, test.statusCode, response.Code)
			assert.JSONEq(t, test.response, response.Body.String())
			if test.call {
				mock.AssertExpectations(t)
			}
		})
	}
}
