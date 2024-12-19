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

func (s *SectionService) Get() (sections map[int]model.Section, err error) {
	sections, err = s.rp.Get()
	return
}

func (s *SectionService) GetById(id int) (section model.Section, err error) {
	section, err = s.rp.GetById(id)
	return
}

func (s *SectionService) Post(section model.Section) (sec model.Section, err error) {
	if err := section.Validate(); err != nil {
		return model.Section{}, err
	}
	sec, err = s.rp.Post(section)
	return
}

func (s *SectionService) Update(id int, section model.Section) (sec model.Section, err error) {
	sec, err = s.rp.Update(id, section)
	return
}

func (s *SectionService) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}
