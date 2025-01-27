package repository

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockSectionRepository struct {
	mock.Mock
}

func (m *MockSectionRepository) Get() ([]model.Section, error) {
	args := m.Called()
	sections := args.Get(0).([]model.Section)
	err := args.Error(1)

	return sections, err
}

func (m *MockSectionRepository) GetById(id int) (model.Section, error) {
	panic("needs implementation...")
}

func (m *MockSectionRepository) Post(section *model.Section) (model.Section, error) {
	panic("needs implementation...")
}
func (m *MockSectionRepository) Update(id int, section *model.Section) (model.Section, error) {
	panic("needs implementation...")
}

func (m *MockSectionRepository) Delete(id int) error {
	panic("needs implementation...")
}

func (sm *MockSectionRepository) CountProductBatchesBySectionId(id int) (countProdBatches model.SectionProductBatches, err error) {
	panic("needs implementation...")
}

func (sm *MockSectionRepository) CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error) {
	panic("needs implementation...")
}
