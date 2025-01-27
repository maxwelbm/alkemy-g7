package service_test

import (
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
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
