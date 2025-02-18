package service_test

import (
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setupRepMock(t *testing.T) *service.SectionService {
	mockRep := mocks.NewMockISectionRepo(t)
	return service.CreateServiceSection(mockRep, logMock)
}

func TestGetSections(t *testing.T) {
	t.Run("return a list of all sections successfully", func(t *testing.T) {
		svc := setupRepMock(t)

		expectedSections := []model.Section{{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}, {ID: 2, SectionNumber: "S02", CurrentTemperature: 15.0, MinimumTemperature: 10.0, CurrentCapacity: 20, MinimumCapacity: 10, MaximumCapacity: 30, WarehouseID: 2, ProductTypeID: 2}}

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)
		mockRepo.On("Get").Return(expectedSections, nil)

		sections, err := svc.Get()

		assert.Equal(t, expectedSections, sections)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestSectionByID(t *testing.T) {
	t.Run("given a valid section by id then return a section with no error", func(t *testing.T) {
		svc := setupRepMock(t)

		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)
		mockRepo.On("GetByID", 1).Return(expectedSection, nil)

		section, err := svc.GetByID(1)

		assert.Equal(t, expectedSection, section)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid id return an error", func(t *testing.T) {
		svc := setupRepMock(t)

		expectedError := customerror.HandleError("section", customerror.ErrorNotFound, "")

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)
		mockRepo.On("GetByID", 50).Return(model.Section{}, expectedError)

		section, err := svc.GetByID(50)

		assert.Equal(t, model.Section{}, section)
		assert.ErrorIs(t, err, expectedError)
		mockRepo.AssertExpectations(t)
	})
}

func TestPostSection(t *testing.T) {
	t.Run("given a valid section create it successfully", func(t *testing.T) {
		svc := setupRepMock(t)

		createdSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)
		mockRepo.On("Post", &model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}).Return(createdSection, nil)

		section, err := svc.Post(&createdSection)

		assert.NoError(t, err)
		assert.Equal(t, createdSection, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid section return error", func(t *testing.T) {
		svc := setupRepMock(t)

		expectedErrSection := customerror.HandleError("section", customerror.ErrorConflict, "")
		createdSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)
		mockRepo.On("Post", &createdSection).Return(model.Section{}, expectedErrSection)

		section, err := svc.Post(&createdSection)

		assert.Equal(t, model.Section{}, section)
		assert.ErrorIs(t, err, expectedErrSection)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid section then return an empty section", func(t *testing.T) {
		svc := setupRepMock(t)

		createdSection := model.Section{ID: 1, CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		section, err := svc.Post(&createdSection)

		assert.Equal(t, model.Section{}, section)
		assert.Error(t, err)
	})
}

func TestUpdateSection(t *testing.T) {
	t.Run("given a valid section then update it", func(t *testing.T) {
		svc := setupRepMock(t)

		updatedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)
		mockRepo.On("GetByID", 1).Return(updatedSection, nil)

		mockRepo.On("Update", 1, &updatedSection).Return(updatedSection, nil)

		section, err := svc.Update(1, &updatedSection)

		assert.NoError(t, err)
		assert.Equal(t, updatedSection, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid section id then return error", func(t *testing.T) {
		svc := setupRepMock(t)

		updatedSection := model.Section{ID: 50, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}
		mockRepo := svc.Rp.(*mocks.MockISectionRepo)

		expectedError := customerror.HandleError("section", customerror.ErrorConflict, "")
		mockRepo.On("GetByID", 50).Return(model.Section{}, expectedError)

		section, err := svc.Update(50, &updatedSection)

		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		assert.Equal(t, model.Section{}, section)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteSection(t *testing.T) {
	t.Run("given a valid id section then delete it", func(t *testing.T) {
		svc := setupRepMock(t)

		deletedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}
		mockRepo := svc.Rp.(*mocks.MockISectionRepo)

		mockRepo.On("GetByID", 1).Return(deletedSection, nil)

		mockRepo.On("CountProductBatchesBySectionID", 1).Return(model.SectionProductBatches{}, nil)

		mockRepo.On("Delete", 1).Return(nil)

		err := svc.Delete(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given an invalid section id then return error", func(t *testing.T) {
		svc := setupRepMock(t)

		expectedError := customerror.HandleError("section", customerror.ErrorNotFound, "")

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)

		mockRepo.On("GetByID", 50).Return(model.Section{}, expectedError)

		err := svc.Delete(50)

		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("given a valid section id that have a dependecy with product batches then return error to delete", func(t *testing.T) {
		svc := setupRepMock(t)

		deletedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}
		prodBatches := model.SectionProductBatches{ID: 1, SectionNumber: "S01", ProductsCount: 1}
		expectedError := customerror.HandleError("section", customerror.ErrorDep, "")

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)

		mockRepo.On("GetByID", 1).Return(deletedSection, nil)

		mockRepo.On("CountProductBatchesBySectionID", 1).Return(prodBatches, expectedError)

		err := svc.Delete(1)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCountProductBatchesBySectionID(t *testing.T) {
	t.Run("given a valid product btaches section then return it", func(t *testing.T) {
		svc := setupRepMock(t)

		expectedSPB := model.SectionProductBatches{ID: 1, SectionNumber: "S01", ProductsCount: 1}

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)

		mockRepo.On("CountProductBatchesBySectionID", 1).Return(expectedSPB, nil)

		count, err := svc.CountProductBatchesBySectionID(1)

		assert.Equal(t, expectedSPB, count)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCountProductBatchesBySection(t *testing.T) {
	t.Run("return a slice with all product batches sections", func(t *testing.T) {
		svc := setupRepMock(t)

		expectedSPB := []model.SectionProductBatches{{ID: 1, SectionNumber: "S01", ProductsCount: 1}, {ID: 2, SectionNumber: "S02", ProductsCount: 1}}

		mockRepo := svc.Rp.(*mocks.MockISectionRepo)

		mockRepo.On("CountProductBatchesSections").Return(expectedSPB, nil)

		count, err := svc.CountProductBatchesSections()

		assert.Equal(t, expectedSPB, count)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
