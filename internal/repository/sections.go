package repository

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

var (
	NotFoundError = errors.New("there's no section with this id")
	ConflictError = errors.New("section with this id already exists")
)

type SectionRepository struct {
	db *sql.DB
}

func CreateRepositorySections(db *sql.DB) *SectionRepository {
	return &SectionRepository{db}
}

func (r *SectionRepository) Get() (sections []model.Section, err error) {
	queryGetAll := "SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM `sections`"
	rows, err := r.db.Query(queryGetAll)
	if err != nil {
		return
	}

	for rows.Next() {
		var section model.Section
		err = rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
		if err != nil {
			return
		}
		sections = append(sections, section)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *SectionRepository) GetById(id int) (section model.Section, err error) {
	// if _, exists := r.dbSection.TbSections[id]; !exists {
	// 	err = NotFoundError
	// 	return
	// }
	// section = r.dbSection.TbSections[id]
	// return
	return model.Section{}, nil
}

func (r *SectionRepository) Post(section model.Section) (s model.Section, err error) {
	// lastId := getLastId(r.dbSection.TbSections)
	// sectionExists := sectionNumberExists(section.SectionNumber, r)
	// if _, exists := r.dbSection.TbSections[section.ID]; exists {
	// 	err = ConflictError
	// 	return
	// }
	// if sectionExists {
	// 	err = ConflictError
	// 	return
	// }
	// section.ID = lastId
	// r.dbSection.TbSections[section.ID] = section
	// s = section
	// return
	return model.Section{}, nil
}

func (r *SectionRepository) Update(id int, section model.Section) (newSec model.Section, err error) {
	// if _, exists := r.dbSection.TbSections[id]; !exists {
	// 	err = NotFoundError
	// 	return
	// }
	// sectionExists := sectionNumberExists(section.SectionNumber, r)
	// if r.dbSection.TbSections[id].SectionNumber != section.SectionNumber {
	// 	if sectionExists {
	// 		err = ConflictError
	// 		return
	// 	}
	// }
	// newSec = section
	// r.dbSection.TbSections[id] = newSec
	// return
	return model.Section{}, nil
}

func (r *SectionRepository) Delete(id int) (err error) {
	// if _, exists := r.dbSection.TbSections[id]; !exists {
	// 	err = NotFoundError
	// 	return
	// }
	// delete(r.dbSection.TbSections, id)
	// return
	return err
}

func sectionNumberExists(sectionNumber string, sr *SectionRepository) bool {
	// for _, section := range sr.dbSection.TbSections {
	// 	if section.SectionNumber == sectionNumber {
	// 		return true
	// 	}
	// }
	// return false
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
