package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockSectionService struct {
	mock.Mock
}

func (sm *MockSectionService) Get() (sections []model.Section, err error) {
	args := sm.Called()
	sections = args.Get(0).([]model.Section)
	err = args.Error(1)

	return
}

func (sm *MockSectionService) GetById(id int) (section model.Section, err error) {
	args := sm.Called(id)
	section = args.Get(0).(model.Section)
	err = args.Error(1)

	return
}

func (sm *MockSectionService) Post(section *model.Section) (sec model.Section, err error) {
	args := sm.Called(section)
	sec = args.Get(0).(model.Section)
	err = args.Error(1)

	return
}

func (sm *MockSectionService) Update(id int, section *model.Section) (sec model.Section, err error) {
	args := sm.Called(id, section)
	sec = args.Get(0).(model.Section)
	err = args.Error(1)

	return
}

func (sm *MockSectionService) Delete(id int) (err error) {
	panic("needs implementation...")
}

func (sm *MockSectionService) CountProductBatchesBySectionId(id int) (countProdBatches model.SectionProductBatches, err error) {
	panic("needs implementation...")
}

func (sm *MockSectionService) CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error) {
	panic("needs implementation...")
}
