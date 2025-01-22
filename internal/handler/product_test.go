package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
	"github.com/stretchr/testify/assert"
)

type StubMockProductService struct {
	FuncGetAllProducts func() ([]model.Product, error)
	FuncGetProductById func(id int) (model.Product, error)
	FuncCreateProduct  func(product model.Product) (model.Product, error)
	FuncUpdateProduct  func(id int, product model.Product) (model.Product, error)
	FuncDeleteProduct  func(id int) error
}

func (s *StubMockProductService) GetAllProducts() ([]model.Product, error) {
	if s.FuncGetAllProducts != nil {
		return s.FuncGetAllProducts()
	}
	return nil, nil
}

func (s *StubMockProductService) GetProductById(id int) (model.Product, error) {
	if s.FuncGetProductById != nil {
		return s.FuncGetProductById(id)
	}
	return model.Product{}, nil
}

func (s *StubMockProductService) CreateProduct(product model.Product) (model.Product, error) {
	if s.FuncCreateProduct != nil {
		return s.FuncCreateProduct(product)
	}
	return model.Product{}, nil
}

func (s *StubMockProductService) UpdateProduct(id int, product model.Product) (model.Product, error) {
	if s.FuncUpdateProduct != nil {
		return s.FuncUpdateProduct(id, product)
	}
	return model.Product{}, nil
}

func (s *StubMockProductService) DeleteProduct(id int) error {
	if s.FuncDeleteProduct != nil {
		return s.FuncDeleteProduct(id)
	}
	return nil
}

func TestGetProductsHandler(t *testing.T) {
	findAll := func() ([]model.Product, error) {
		return []model.Product{
			{ID: 1, ProductCode: "P001", Description: "Product 1", Width: 10, Height: 20, Length: 0, NetWeight: 0, ExpirationRate: 0, RecommendedFreezingTemperature: 0, FreezingRate: 0, ProductTypeID: 0, SellerID: 0},
			{ID: 2, ProductCode: "P002", Description: "Product 2", Width: 15, Height: 25, Length: 0, NetWeight: 0, ExpirationRate: 0, RecommendedFreezingTemperature: 0, FreezingRate: 0, ProductTypeID: 0, SellerID: 0},
		}, nil
	}

	productHd := ProductHandler{
		ProductService: &StubMockProductService{FuncGetAllProducts: findAll},
	}

	t.Run("GetAllProducts - should return list of products", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/products", nil)
		res := httptest.NewRecorder()

		productHd.GetAllProducts(res, req)

		expected := `{"data":[{"id":1,"product_code":"P001","description":"Product 1","width":10,"height":20,"length":0,"net_weight":0,"expiration_rate":0,"recommended_freezing_temperature":0,"freezing_rate":0,"product_type_id":0,"seller_id":0},{"id":2,"product_code":"P002","description":"Product 2","width":15,"height":25,"length":0,"net_weight":0,"expiration_rate":0,"recommended_freezing_temperature":0,"freezing_rate":0,"product_type_id":0,"seller_id":0}]}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})
}

func TestGetProductById(t *testing.T) {
	getById := func(id int) (model.Product, error) {
		if id == 1 {
			return model.Product{
				ID:                             1,
				ProductCode:                    "P001",
				Description:                    "Product 1",
				Width:                          10,
				Height:                         20,
				Length:                         0,
				NetWeight:                      0,
				ExpirationRate:                 0,
				RecommendedFreezingTemperature: 0,
				FreezingRate:                   0,
				ProductTypeID:                  0,
				SellerID:                       0,
			}, nil
		}
		return model.Product{}, custom_error.HandleError("product", custom_error.ErrorNotFound, "")
	}

	productHd := ProductHandler{
		ProductService: &StubMockProductService{FuncGetProductById: getById},
	}

	r := chi.NewRouter()
	r.Get("/api/v1/products/{id}", productHd.GetProductById)

	testCases := []struct {
		name               string
		id                 string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:               "should return the product requested and 200 ok",
			id:                 "1",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"data":{"id":1,"product_code":"P001","description":"Product 1","width":10,"height":20,"length":0,"net_weight":0,"expiration_rate":0,"recommended_freezing_temperature":0,"freezing_rate":0,"product_type_id":0,"seller_id":0}}`,
		},
		{
			name:               "should return 404 not found for non-existent product",
			id:                 "2",
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   `{"message":"product not found"}`,
		},
		{
			name:               "should return 400 bad request for invalid id",
			id:                 "abc",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"message":"invalid id"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/products/"+tc.id, nil)
			res := httptest.NewRecorder()

			r.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			assert.JSONEq(t, tc.expectedResponse, res.Body.String())
		})
	}
}

func TestInsertProduct(t *testing.T) {
	createRequest := func(body string) *http.Request {
		req := httptest.NewRequest("POST", "/api/v1/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		return req
	}

	insertProduct := func(product model.Product) (model.Product, error) {
		product.ID = 1
		return product, nil
	}

	productHd := ProductHandler{
		ProductService: &StubMockProductService{FuncCreateProduct: insertProduct},
	}

	testCases := []struct {
		name               string
		productRequestBody string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "should return 201 created and the new product created",
			productRequestBody: `{
				"product_code": "P003",
				"description": "New Product",
				"width": 1,
				"height": 1,
				"length": 1,
				"net_weight": 1,
				"expiration_rate": 1,
				"recommended_freezing_temperature":1,
				"freezing_rate": 1,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   `{"data":{"id":1,"product_code":"P003","description":"New Product","width":1,"height":1,"length":1,"net_weight":1,"expiration_rate":1,"recommended_freezing_temperature":1,"freezing_rate":1,"product_type_id":1,"seller_id":1}}`,
		},
		{
			name: "should return 400 bad request for invalid type fields",
			productRequestBody: `{
				"product_code": 0
			}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedResponse:   `{"message":"invalid json syntax"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := createRequest(tc.productRequestBody)
			res := httptest.NewRecorder()

			productHd.CreateProduct(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			assert.JSONEq(t, tc.expectedResponse, res.Body.String())
		})
	}
}

func TestUpdateProduct(t *testing.T) {
    createRequest := func(id string, body string) *http.Request {
        req := httptest.NewRequest("PATCH", "/api/v1/products/"+id, strings.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        return req
    }

    updateProduct := func(id int, product model.Product) (model.Product, error) {
        if id != 1 {
            return model.Product{}, custom_error.HandleError("product", custom_error.ErrorNotFound, "")
        }
        product.ID = 1
        return product, nil
    }

    productHd := ProductHandler{
        ProductService: &StubMockProductService{FuncUpdateProduct: updateProduct},
    }

    r := chi.NewRouter()
    r.Patch("/api/v1/products/{id}", productHd.UpdateProduct) // Registrando o handler de atualização

    testCases := []struct {
        name               string
        id                 string
        productRequestBody string
        expectedStatusCode int
        expectedResponse   string
    }{
        {
            name: "should return 200 ok and update the product",
            id:   "1",
            productRequestBody: `{
                "product_code": "P003",
                "description": "New Product",
                "width": 1,
                "height": 1,
                "length": 1,
                "net_weight": 1,
                "expiration_rate": 1,
                "recommended_freezing_temperature": 1,
                "freezing_rate": 1,
                "product_type_id": 1,
                "seller_id": 1
            }`,
            expectedStatusCode: http.StatusOK,
            expectedResponse: `{"data":{"id":1,"product_code":"P003","description":"New Product","width":1,"height":1,"length":1,"net_weight":1,"expiration_rate":1,"recommended_freezing_temperature":1,"freezing_rate":1,"product_type_id":1,"seller_id":1}}`,
        },
		{
            name: "should return 200 ok and update the product",
            id:   "2",
            productRequestBody: `{
                "product_code": "P003",
                "description": "New Product",
                "width": 1,
                "height": 1,
                "length": 1,
                "net_weight": 1,
                "expiration_rate": 1,
                "recommended_freezing_temperature": 1,
                "freezing_rate": 1,
                "product_type_id": 1,
                "seller_id": 1
            }`,
            expectedStatusCode: http.StatusNotFound,
            expectedResponse: `{"message":"product not found"}`,
        },
		{
            name: "should return 200 ok and update the product",
            id:   "1",
            productRequestBody: `{
                "product_code": 1,
            }`,
            expectedStatusCode: http.StatusUnprocessableEntity,
            expectedResponse: `{"message":"invalid json syntax"}`,
        },
		{
            name: "should return 200 ok and update the product",
            id:   "ad",
            productRequestBody: `{
                "product_code": 1,
            }`,
            expectedStatusCode: http.StatusBadRequest,
            expectedResponse: `{"message":"invalid id"}`,
        },
        
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            req := createRequest(tc.id, tc.productRequestBody)
            res := httptest.NewRecorder()

            r.ServeHTTP(res, req) // Agora usando o router para servir a requisição

            assert.Equal(t, tc.expectedStatusCode, res.Code)
            assert.JSONEq(t, tc.expectedResponse, res.Body.String())
        })
    }
}

func TestDeleteProduct(t *testing.T) {
    // Função para simular a deleção do produto
    deleteProduct := func(id int) error {
        if id != 1 {
            return custom_error.HandleError("product", custom_error.ErrorNotFound, "")
        }
        return nil
    }

    // Configura o handler com um Stub do serviço
    productHd := ProductHandler{
        ProductService: &StubMockProductService{FuncDeleteProduct: deleteProduct},
    }

    // Configuração do roteador com chi
    r := chi.NewRouter()
    r.Delete("/api/v1/products/{id}", productHd.DeleteProductById)

    testCases := []struct {
        name               string
        id                 string
        expectedStatusCode int
        expectedResponse   string
    }{
        {
            name:               "should return 204 no content when product is deleted",
            id:                 "1",
            expectedStatusCode: http.StatusNoContent,
            expectedResponse:   `{"message":"product deleted"}`,
        },
        {
            name:               "should return 404 not found when product does not exist",
            id:                 "2",
            expectedStatusCode: http.StatusNotFound,
            expectedResponse:   `{"message":"product not found"}`,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            req := httptest.NewRequest("DELETE", "/api/v1/products/"+tc.id, nil)
            res := httptest.NewRecorder()

            // Simula a requisição ao roteador
            r.ServeHTTP(res, req)

            assert.Equal(t, tc.expectedStatusCode, res.Code)
            assert.JSONEq(t, tc.expectedResponse, res.Body.String())
        })
    }
}
