package repository

import (
	"errors"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

var (
	NotFoundError = errors.New("there's no section with this id")
	ConflictError = errors.New("section with this id already exists")
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

func (r *SectionRepository) Post(section model.Section) (s model.Section, err error) {
	lastId := getLastId(r.dbSection.TbSections)
	sectionExists := sectionNumberExists(section.SectionNumber, r)
	if _, exists := r.dbSection.TbSections[section.ID]; exists {
		err = ConflictError
		return
	}
	if sectionExists {
		err = ConflictError
		return
	}
	section.ID = lastId
	r.dbSection.TbSections[section.ID] = section
	s = section
	return
}

func (r *SectionRepository) Update(id int, section model.Section) (newSec model.Section, err error) {
	newSec = section
	r.dbSection.TbSections[id] = newSec
	return
}

func (r *SectionRepository) Delete(id int) (err error) {
	if _, exists := r.dbSection.TbSections[id]; !exists {
		err = NotFoundError
		return
	}
	delete(r.dbSection.TbSections, id)
	return
}

func sectionNumberExists(sectionNumber int, sr *SectionRepository) bool {
	for _, section := range sr.dbSection.TbSections {
		if section.SectionNumber == sectionNumber {
			return true
		}
	}
	return false
}

func getLastId(db map[int]model.Section) int {
	lastId := 0
	for _, section := range db {
		if section.ID > lastId {
			lastId = section.ID
		}
	}
	return lastId + 1
}
