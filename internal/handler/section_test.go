package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setupSectionService(t *testing.T) *handler.SectionController {
	mockSectionService := mocks.NewMockISectionService(t)
	hd := handler.CreateHandlerSections(mockSectionService, logMock)
	return hd
}

func TestHandlerGet(t *testing.T) {
	t.Run("return a list of all existing sections successfully", func(t *testing.T) {
		hd := setupSectionService(t)

		expectedSections := []model.Section{{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}, {ID: 2, SectionNumber: "S02", CurrentTemperature: 15.0, MinimumTemperature: 10.0, CurrentCapacity: 20, MinimumCapacity: 10, MaximumCapacity: 30, WarehouseID: 2, ProductTypeID: 2}}

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Get").Return(expectedSections, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
		response := httptest.NewRecorder()

		hd.GetAll(response, request)

		expectedJSON := `{
		"data":[
		{
			"id": 1,
			"section_number": "S01",
			"current_temperature": 10.0,
			"minimum_temperature": 5.0,
			"current_capacity": 10,
			"minimum_capacity": 5,
			"maximum_capacity": 20,
			"warehouse_id": 1,
			"product_type_id": 1
		},
		{
			"id": 2,
			"section_number": "S02",
			"current_temperature": 15.0,
			"minimum_temperature": 10.0,
			"current_capacity": 20,
			"minimum_capacity": 10,
			"maximum_capacity": 30,
			"warehouse_id": 2,
			"product_type_id": 2
		}
		]
		}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return an error when fetching sections", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Get").Return([]model.Section{}, errors.New("unable to list sections"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
		response := httptest.NewRecorder()

		hd.GetAll(response, request)

		expectedJSON := `{"message":"unable to list sections"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJSON, response.Body.String())
	})
}

func TestHandlerGetSectionByID(t *testing.T) {
	t.Run("return section by id if it exist", func(t *testing.T) {
		hd := setupSectionService(t)

		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("GetByID", expectedSection.ID).Return(expectedSection, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()

		hd.GetByID(response, request)

		expectedSectionJSON := `{"data":{
			"id": 1,
			"section_number": "S01",
			"current_temperature": 10.0,
			"minimum_temperature": 5.0,
			"current_capacity": 10,
			"minimum_capacity": 5,
			"maximum_capacity": 20,
			"warehouse_id": 1,
			"product_type_id": 1
		}, "message": "success"
		}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return section not found", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("GetByID", 100).Return(model.Section{}, customerror.HandleError("section", customerror.ErrorNotFound, ""))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/100", nil)
		response := httptest.NewRecorder()

		hd.GetByID(response, request)

		expectedSectionJSON := `{"message": "section not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid id then return error", func(t *testing.T) {
		hd := setupSectionService(t)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/aaa", nil)
		response := httptest.NewRecorder()

		hd.GetByID(response, request)

		expectedJSON := `{"message": "invalid id param"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJSON, response.Body.String())
	})

	t.Run("return unable to search for section", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("GetByID", 100).Return(model.Section{}, errors.New("unable to search for section"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/100", nil)
		response := httptest.NewRecorder()

		hd.GetByID(response, request)

		expectedSectionJSON := `{"message": "unable to search for section"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})
}

func TestHandlerCreateSection(t *testing.T) {
	t.Run("given a valid section then create successfully", func(t *testing.T) {
		hd := setupSectionService(t)

		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Post", &model.Section{SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}).Return(expectedSection, nil)

		reqBody := []byte(`{
			"section_number": "S01",
			"current_temperature": 10.0,
			"minimum_temperature": 5.0,
			"current_capacity": 10,
			"minimum_capacity": 5,
			"maximum_capacity": 20,
			"warehouse_id": 1,
			"product_type_id": 1
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections/", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Post(response, request)

		expectedSectionJSON := `{"data":{
			"id": 1,
			"section_number": "S01",
			"current_temperature": 10.0,
			"minimum_temperature": 5.0,
			"current_capacity": 10,
			"minimum_capacity": 5,
			"maximum_capacity": 20,
			"warehouse_id": 1,
			"product_type_id": 1
		}
		}`

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid section field then return a error", func(t *testing.T) {
		hd := setupSectionService(t)
		reqBody := []byte(`{
			"section_number": "",
			"current_temperature": 0.0,
			"minimum_temperature": 0.0
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections/", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Post(response, request)

		expectedSectionJSON := `{"message": "request body cannot be empty"}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
	})

	t.Run("given a valid section that already exist then return error", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Post", &model.Section{SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}).Return(model.Section{}, customerror.HandleError("section", customerror.ErrorConflict, ""))

		reqBody := []byte(`{
			"section_number": "S01",
			"current_temperature": 10.0,
			"minimum_temperature": 5.0,
			"current_capacity": 10,
			"minimum_capacity": 5,
			"maximum_capacity": 20,
			"warehouse_id": 1,
			"product_type_id": 1
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections/", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Post(response, request)

		expectedSectionJSON := `{"message": "section it already exists"}`

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid json then return error", func(t *testing.T) {
		hd := setupSectionService(t)

		reqBody := []byte(`{
			"section_number": "S01",
			"current_temperature": 10.0,
			InvalidField
		}`)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections/", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Post(response, request)

		expectedSectionJSON := `{"message": "invalid request body"}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
	})

	t.Run("return unable to create section", func(t *testing.T) {
		hd := setupSectionService(t)

		expSection := model.Section{SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		reqBody := []byte(`{
			"section_number": "S01",
			"current_temperature": 10.0,
			"minimum_temperature": 5.0,
			"current_capacity": 10,
			"minimum_capacity": 5,
			"maximum_capacity": 20,
			"warehouse_id": 1,
			"product_type_id": 1
		}`)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Post", &expSection).Return(model.Section{}, errors.New("unable to create section"))

		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections/", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Post(response, request)

		expectedSectionJSON := `{"message": "unable to create section"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})
}

func TestHandlerUpdateSection(t *testing.T) {
	t.Run("given a valid section to update then update it with no error", func(t *testing.T) {
		hd := setupSectionService(t)

		updatedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 12.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Update", 1, &model.Section{ID: 1, CurrentTemperature: 14.0}).Return(updatedSection, nil)

		reqBody := []byte(`{"current_temperature": 14.0}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/1", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Update(response, request)

		expectedSectionJSON := `{
		"data":
		{
			"id": 1,
			"section_number": "S01",
			"current_temperature": 12.0,
			"minimum_temperature": 5.0,
			"current_capacity": 10,
			"minimum_capacity": 5,
			"maximum_capacity": 20,
			"warehouse_id": 1,
			"product_type_id": 1
		}
		}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given a invalid section to update then return an error", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Update", 50, &model.Section{ID: 50, CurrentTemperature: 5.0}).Return(model.Section{}, customerror.HandleError("section", customerror.ErrorNotFound, ""))

		reqBody := []byte(`{
			"current_temperature": 5.0
		}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/50", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Update(response, request)

		expectedSectionJSON := `{"message": "section not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid param then return an error", func(t *testing.T) {
		hd := setupSectionService(t)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/aaa", nil)
		response := httptest.NewRecorder()

		hd.Update(response, request)

		expectedSectionJSON := `{"message": "invalid id param"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
	})

	t.Run("given an invalid request body then return an error", func(t *testing.T) {
		hd := setupSectionService(t)

		reqBody := []byte(`{
			"current_temperature": 5.0,
			InvalidRequest
		}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/50", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Update(response, request)

		expectedSectionJSON := `{"message": "invalid request body"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
	})

	t.Run("given an empty request body then return an error", func(t *testing.T) {
		hd := setupSectionService(t)

		reqBody := []byte(`{}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/50", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Update(response, request)

		expectedSectionJSON := `{"message": "request body cannot be empty"}`

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
	})

	t.Run("given an invalid section to update then return an error", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Update", 50, &model.Section{ID: 50, CurrentTemperature: 5.0}).Return(model.Section{}, errors.New("unable to update section"))

		reqBody := []byte(`{
			"current_temperature": 5.0
		}`)

		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/50", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()

		hd.Update(response, request)

		expectedSectionJSON := `{"message": "unable to update section"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})
}

func TestHandlerDeleteSection(t *testing.T) {
	t.Run("given a valid section id then delete this section", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Delete", 1).Return(nil)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()

		hd.Delete(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid section id then return error", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Delete", 50).Return(customerror.HandleError("section", customerror.ErrorNotFound, ""))

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/50", nil)
		response := httptest.NewRecorder()

		hd.Delete(response, request)

		expectedSectionJSON := `{"message": "section not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid section id then return error", func(t *testing.T) {
		hd := setupSectionService(t)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/aaa", nil)
		response := httptest.NewRecorder()

		hd.Delete(response, request)

		expectedSectionJSON := `{"message": "invalid id param"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
	})

	t.Run("given an invalid section then return error", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("Delete", 50).Return(errors.New("unable to delete section"))

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/50", nil)
		response := httptest.NewRecorder()

		hd.Delete(response, request)

		expectedSectionJSON := `{"message": "unable to delete section"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})
}

func TestHandlerCountProductBatchesSections(t *testing.T) {
	t.Run("return the count of all product batches successfully", func(t *testing.T) {
		hd := setupSectionService(t)

		countProductBatchesSections := []model.SectionProductBatches{
			{ID: 1, SectionNumber: "S01", ProductsCount: 1},
		}

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("CountProductBatchesSections").Return(countProductBatchesSections, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/reportProducts", nil)
		response := httptest.NewRecorder()

		hd.CountProductBatchesSections(response, request)

		expectedJson := `{
    "data": [
        {
            "id": 1,
            "section_number": "S01",
            "products_count": 1
        }
    ]
}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return a error when count product batches", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("CountProductBatchesSections").Return([]model.SectionProductBatches{}, errors.New("unable to count section product batches"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/reportProducts", nil)
		response := httptest.NewRecorder()

		hd.CountProductBatchesSections(response, request)

		expectedJson := `{
    "message": "unable to count section product batches"
}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return the count of a product batches section successfully", func(t *testing.T) {
		hd := setupSectionService(t)

		countProductBatchesSection := model.SectionProductBatches{
			ID: 1, SectionNumber: "S01", ProductsCount: 1,
		}

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("CountProductBatchesBySectionID", countProductBatchesSection.ID).Return(countProductBatchesSection, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/reportProducts?id=1", nil)
		response := httptest.NewRecorder()

		hd.CountProductBatchesSections(response, request)

		expectedJson := `{
			"data":
			  {
				"id": 1,
				"products_count": 1,
				"section_number": "S01"
			  }
		  }`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return the unmmaped error", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("CountProductBatchesBySectionID", 1).Return(model.SectionProductBatches{}, errors.New("unable to count section product batches"))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/reportProducts?id=1", nil)
		response := httptest.NewRecorder()

		hd.CountProductBatchesSections(response, request)

		expectedJson := `{"message":"error"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return the custom error", func(t *testing.T) {
		hd := setupSectionService(t)

		mockService := hd.Sv.(*mocks.MockISectionService)
		mockService.On("CountProductBatchesBySectionID", 1).Return(model.SectionProductBatches{}, customerror.HandleError("section", 0, ""))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/reportProducts?id=1", nil)
		response := httptest.NewRecorder()

		hd.CountProductBatchesSections(response, request)

		expectedJson := `{"message":"section unknow server error"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("return error if id is invalid", func(t *testing.T) {
		hd := setupSectionService(t)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/reportProducts?id=aaa", nil)
		response := httptest.NewRecorder()

		hd.CountProductBatchesSections(response, request)

		expectedJson := `{"message":"invalid id"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJson, response.Body.String())
	})
}
