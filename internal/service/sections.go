package service

import (
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type SectionService struct {
	Rp  interfaces.ISectionRepo
	log logger.Logger
}

func CreateServiceSection(rp interfaces.ISectionRepo, log logger.Logger) *SectionService {
	return &SectionService{Rp: rp, log: log}
}

func (s *SectionService) Get() (sections []model.Section, err error) {
	s.log.Log("SectionService", "INFO", "initializing Get function")
	sections, err = s.Rp.Get()

	return
}

func (s *SectionService) GetByID(id int) (section model.Section, err error) {
	s.log.Log("SectionService", "INFO", "initializing GetByID function with id param")
	section, err = s.Rp.GetByID(id)

	return
}

func (s *SectionService) Post(section *model.Section) (sec model.Section, err error) {
	s.log.Log("SectionService", "INFO", "initializing Post function with section param")

	if err := section.Validate(); err != nil {
		s.log.Log("SectionService", "ERROR", fmt.Sprintf("Error: %v", err))
		return model.Section{}, err
	}

	sec, err = s.Rp.Post(section)

	s.log.Log("SectionService", "INFO", "successfully executed post function")

	return
}

func (s *SectionService) Update(id int, section *model.Section) (sec model.Section, err error) {
	s.log.Log("SectionService", "INFO", "initializing Update function with id and section param")

	existingSection, err := s.GetByID(id)
	if err != nil {
		sec = model.Section{}

		s.log.Log("SectionService", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	updateSectionFields(&existingSection, section)

	sec, err = s.Rp.Update(id, &existingSection)

	s.log.Log("SectionService", "INFO", "successfully executed update function")

	return
}

func (s *SectionService) Delete(id int) (err error) {
	s.log.Log("SectionService", "INFO", "initializing Delete function with id param")

	_, err = s.GetByID(id)
	if err != nil {
		s.log.Log("SectionService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	secProdBatches, _ := s.Rp.CountProductBatchesBySectionID(id)
	if secProdBatches.ProductsCount > 0 {
		s.log.Log("SectionService", "ERROR", fmt.Sprintf("Error: %v", err))
		return customerror.HandleError("section", customerror.ErrorDep, "")
	}

	err = s.Rp.Delete(id)

	s.log.Log("SectionService", "INFO", "successfully executed delete function")

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

func (s *SectionService) CountProductBatchesBySectionID(id int) (countProdBatches model.SectionProductBatches, err error) {
	s.log.Log("SectionService", "INFO", "initializing CountProductBatchesBySectionID function with id param")
	countProdBatches, err = s.Rp.CountProductBatchesBySectionID(id)

	return
}

func (s *SectionService) CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error) {
	s.log.Log("SectionService", "INFO", "initializing CountProductBatchesSections function")
	countProductBatches, err = s.Rp.CountProductBatchesSections()

	return
}
