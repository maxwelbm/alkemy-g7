package service_test

import (
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
	"github.com/stretchr/testify/assert"
)

func setupRepMock() *service.SectionService {
	mockRep := new(repository.MockSectionRepository)
	return service.CreateServiceSection(mockRep)
}

func TestGetSections(t *testing.T) {
	t.Run("return a list of all sections successfully", func(t *testing.T) {
		svc := setupRepMock()

		expectedSections := []model.Section{{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}, {ID: 2, SectionNumber: "S02", CurrentTemperature: 15.0, MinimumTemperature: 10.0, CurrentCapacity: 20, MinimumCapacity: 10, MaximumCapacity: 30, WarehouseID: 2, ProductTypeID: 2}}

		mockRepo := svc.Rp.(*repository.MockSectionRepository)
		mockRepo.On("Get").Return(expectedSections, nil)

		sections, err := svc.Get()

		assert.Equal(t, expectedSections, sections)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestSectionByID(t *testing.T) {
	t.Run("given a valid section by id then return a section with no error", func(t *testing.T) {
		svc := setupRepMock()

		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*repository.MockSectionRepository)
		mockRepo.On("GetById", 1).Return(expectedSection, nil)

		section, err := svc.GetById(1)

		assert.Equal(t, expectedSection, section)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid id return an error", func(t *testing.T) {
		svc := setupRepMock()

		expectedError := customError.HandleError("section", customError.ErrorNotFound, "")

		mockRepo := svc.Rp.(*repository.MockSectionRepository)
		mockRepo.On("GetById", 50).Return(model.Section{}, expectedError)

		section, err := svc.GetById(50)

		assert.Equal(t, model.Section{}, section)
		assert.ErrorIs(t, err, expectedError)
		mockRepo.AssertExpectations(t)
	})
}

func TestPostSection(t *testing.T) {
	t.Run("given a valid section create it successfully", func(t *testing.T) {
		svc := setupRepMock()

		createdSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*repository.MockSectionRepository)
		mockRepo.On("Post", &model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}).Return(createdSection, nil)

		section, err := svc.Post(&createdSection)

		assert.NoError(t, err)
		assert.Equal(t, createdSection, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid section return error", func(t *testing.T) {
		svc := setupRepMock()

		expectedErrSection := customError.HandleError("section", customError.ErrorConflict, "")
		createdSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*repository.MockSectionRepository)
		mockRepo.On("Post", &createdSection).Return(model.Section{}, expectedErrSection)

		section, err := svc.Post(&createdSection)

		assert.Equal(t, model.Section{}, section)
		assert.ErrorIs(t, err, expectedErrSection)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateSection(t *testing.T) {
	t.Run("given a valid section then update it", func(t *testing.T) {
		svc := setupRepMock()

		updatedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*repository.MockSectionRepository)
		mockRepo.On("GetById", 1).Return(updatedSection, nil)

		mockRepo.On("Update", 1, &updatedSection).Return(updatedSection, nil)

		section, err := svc.Update(1, &updatedSection)

		assert.NoError(t, err)
		assert.Equal(t, updatedSection, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid section id then return error", func(t *testing.T) {
		svc := setupRepMock()

		updatedSection := model.Section{ID: 50, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}
		mockRepo := svc.Rp.(*repository.MockSectionRepository)

		expectedError := customError.HandleError("section", customError.ErrorConflict, "")
		mockRepo.On("GetById", 50).Return(model.Section{}, expectedError)

		section, err := svc.Update(50, &updatedSection)

		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		assert.Equal(t, model.Section{}, section)
		mockRepo.AssertExpectations(t)
	})
}
