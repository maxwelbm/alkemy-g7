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
	args := m.Called(id)
	section := args.Get(0).(model.Section)
	err := args.Error(1)

	return section, err
}

func (m *MockSectionRepository) Post(section *model.Section) (model.Section, error) {
	args := m.Called(section)
	sec := args.Get(0).(model.Section)
	err := args.Error(1)

	return sec, err
}
func (m *MockSectionRepository) Update(id int, section *model.Section) (model.Section, error) {
	args := m.Called(id, section)
	sec := args.Get(0).(model.Section)
	err := args.Error(1)

	return sec, err
}

func (m *MockSectionRepository) Delete(id int) error {
	args := m.Called(id)
	err := args.Error(0)

	return err
}

func (sm *MockSectionRepository) CountProductBatchesBySectionId(id int) (countProdBatches model.SectionProductBatches, err error) {
	args := sm.Called(id)
	countProdBatches = args.Get(0).(model.SectionProductBatches)
	err = args.Error(1)

	return
}

func (sm *MockSectionRepository) CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error) {
	panic("needs implementation...")
}
