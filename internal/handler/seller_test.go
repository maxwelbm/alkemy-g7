package handler_test

import (
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

