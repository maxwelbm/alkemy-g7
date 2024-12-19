package repository

import (
	"errors"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

var (
	NotFoundError = errors.New("there's no section with this id")
)

type SectionRepository struct {
	dbSection database.Database
}

func CreateRepositorySections(db database.Database) *SectionRepository {
	return &SectionRepository{dbSection: db}
}

func (r *SectionRepository) Get() (sections map[int]model.Section, err error) {
	sections = make(map[int]model.Section)

	for _, section := range r.dbSection.TbSections {
		sections[section.ID] = section
	}

	return
}

func (r *SectionRepository) GetById(id int) (section model.Section, err error) {
	if _, exists := r.dbSection.TbSections[id]; !exists {
		err = NotFoundError
		return
	}
	section = r.dbSection.TbSections[id]
	return
}

func (r *SectionRepository) Post(section model.Section) (model.Section, error) {
	return model.Section{}, nil
}

func (r *SectionRepository) Update(id int, section model.Section) (model.Section, error) {
	return model.Section{}, nil
}

func (r *SectionRepository) Delete(id int) error {
	return nil
}
