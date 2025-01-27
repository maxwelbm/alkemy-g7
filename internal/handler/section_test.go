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

func setupSectionService() *handler.SectionController {
	mockSectionService := new(service.MockSectionService)
	hd := handler.CreateHandlerSections(mockSectionService)
	return hd
}

func TestHandlerGet(t *testing.T) {
	t.Run("return a list of all existing sections successfully", func(t *testing.T) {
		hd := setupSectionService()

		expectedSections := []model.Section{{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}, {ID: 2, SectionNumber: "S02", CurrentTemperature: 15.0, MinimumTemperature: 10.0, CurrentCapacity: 20, MinimumCapacity: 10, MaximumCapacity: 30, WarehouseID: 2, ProductTypeID: 2}}

		mockService := hd.Sv.(*service.MockSectionService)
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
		hd := setupSectionService()

		mockService := hd.Sv.(*service.MockSectionService)
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
		hd := setupSectionService()

		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockService := hd.Sv.(*service.MockSectionService)
		mockService.On("GetById", expectedSection.ID).Return(expectedSection, nil)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()

		hd.GetById(response, request)

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
		hd := setupSectionService()

		mockService := hd.Sv.(*service.MockSectionService)
		mockService.On("GetById", 100).Return(model.Section{}, custom_error.HandleError("section", custom_error.ErrorNotFound, ""))

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/100", nil)
		response := httptest.NewRecorder()

		hd.GetById(response, request)

		expectedSectionJSON := `{"message": "section not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, expectedSectionJSON, response.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("given an invalid id then return error", func(t *testing.T) {
		hd := setupSectionService()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/aaa", nil)
		response := httptest.NewRecorder()

		hd.GetById(response, request)

		expectedJSON := `{"message": "invalid id param"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, expectedJSON, response.Body.String())
	})
}

func TestHandlerCreateSection(t *testing.T) {
	t.Run("given a valid section then create successfully", func(t *testing.T) {
		hd := setupSectionService()

		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockService := hd.Sv.(*service.MockSectionService)
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
		hd := setupSectionService()
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
		hd := setupSectionService()

		mockService := hd.Sv.(*service.MockSectionService)
		mockService.On("Post", &model.Section{SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}).Return(model.Section{}, custom_error.HandleError("section", custom_error.ErrorConflict, ""))

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
}

func TestHandlerUpdateSection(t *testing.T) {
	t.Run("given a valid section to update then update it with no error", func(t *testing.T) {
		hd := setupSectionService()

		updatedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 12.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockService := hd.Sv.(*service.MockSectionService)
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
		hd := setupSectionService()

		mockService := hd.Sv.(*service.MockSectionService)
		mockService.On("Update", 50, &model.Section{ID: 50, CurrentTemperature: 5.0}).Return(model.Section{}, custom_error.HandleError("section", custom_error.ErrorNotFound, ""))

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
}
