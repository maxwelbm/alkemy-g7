package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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
	queryGetById := "SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM `sections` WHERE `id` = ?"
	row := r.db.QueryRow(queryGetById, id)

	err = row.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.NotFoundErrorSection
		}
		return
	}
	return
}

func (r *SectionRepository) Post(section *model.Section) (s model.Section, err error) {
	queryPost := "INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(queryPost, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = custom_error.ConflictErrorSection
			default:
			}
			return
		}
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	(*section).ID = int(id)

	s, _ = r.GetById(int(id))

	return
}

func (r *SectionRepository) Update(id int, section *model.Section) (newSec model.Section, err error) {
	queryUpdate := "UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?"
	_, err = r.db.Exec(queryUpdate, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID, id)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = custom_error.ConflictErrorSection
			default:
			}
			return
		}
		return
	}

	newSec, _ = r.GetById(id)

	return
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
