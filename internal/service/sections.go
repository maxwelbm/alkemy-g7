package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
)

type SectionService struct {
	rp repository.SectionRepository
}

func CreateServiceSection(rp repository.SectionRepository) *SectionService {
	return &SectionService{rp: rp}
}

func (s *SectionService) Get() (sections []model.Section, err error) {
	sections, err = s.rp.Get()
	return
}

func (s *SectionService) GetById(id int) (section model.Section, err error) {
	section, err = s.rp.GetById(id)
	return
}

func (s *SectionService) Post(section *model.Section) (sec model.Section, err error) {
	if err := section.Validate(); err != nil {
		return model.Section{}, err
	}
	sec, err = s.rp.Post(section)
	return
}

func (s *SectionService) Update(id int, section *model.Section) (sec model.Section, err error) {
	existingSection, err := s.GetById(id)
	if err != nil {
		sec = model.Section{}
		return
	}

	updateSectionFields(&existingSection, section)

	sec, err = s.rp.Update(id, &existingSection)
	return
}

func (s *SectionService) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}

func updateSectionFields(existingSection *model.Section, updatedSection *model.Section) {
	if updatedSection.SectionNumber != "" {
		existingSection.SectionNumber = updatedSection.SectionNumber
	}
	if updatedSection.CurrentTemperature != 0 {
		existingSection.CurrentTemperature = updatedSection.CurrentTemperature
	}
	if updatedSection.MinimumTemperature != 0 {
		existingSection.MinimumTemperature = updatedSection.MinimumTemperature
	}
	if updatedSection.CurrentCapacity != 0 {
		existingSection.CurrentCapacity = updatedSection.CurrentCapacity
	}
	if updatedSection.MinimumCapacity != 0 {
		existingSection.MinimumCapacity = updatedSection.MinimumCapacity
	}
	if updatedSection.MaximumCapacity != 0 {
		existingSection.MaximumCapacity = updatedSection.MaximumCapacity
	}
	if updatedSection.WarehouseID != 0 {
		existingSection.WarehouseID = updatedSection.WarehouseID
	}
	if updatedSection.ProductTypeID != 0 {
		existingSection.ProductTypeID = updatedSection.WarehouseID
	}
}

func (s *SectionService) CountProductBatchesBySectionId(id int) (countProdBatches model.SectionProductBatches, err error) {
	countProdBatches, err = s.rp.CountProductBatchesBySectionId(id)
	return
}

func (s *SectionService) CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error) {
	countProductBatches, err = s.rp.CountProductBatchesSections()
	return
}
