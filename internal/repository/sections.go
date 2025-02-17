package repository

import (
	"database/sql"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type SectionRepository struct {
	db  *sql.DB
	log *logger.Logger
}

func CreateRepositorySections(db *sql.DB, log *logger.Logger) *SectionRepository {
	return &SectionRepository{db: db, log: log}
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

func (r *SectionRepository) GetByID(id int) (section model.Section, err error) {
	queryGetByID := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?"
	row := r.db.QueryRow(queryGetByID, id)

	err = row.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
			return
		}

		return
	}

	return
}

func (r *SectionRepository) Post(section *model.Section) (s model.Section, err error) {
	queryPost := "INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(queryPost, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.HandleError("section", customerror.ErrorConflict, "")
		}

		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	(*section).ID = int(id)

	s, _ = r.GetByID(int(id))

	return
}

func (r *SectionRepository) Update(id int, section *model.Section) (newSec model.Section, err error) {
	queryUpdate := "UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?"
	_, err = r.db.Exec(queryUpdate, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID, id)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
			return
		}

		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.HandleError("section", customerror.ErrorConflict, "")
			return
		}

		return
	}

	newSec, _ = r.GetByID(id)

	return
}

func (r *SectionRepository) Delete(id int) (err error) {
	queryDelete := "DELETE FROM `sections` WHERE `id` = ?"
	_, err = r.db.Exec(queryDelete, id)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
			return
		}

		return
	}

	return
}

func (r *SectionRepository) CountProductBatchesBySectionID(id int) (countProdBatches model.SectionProductBatches, err error) {
	query := "SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id WHERE s.id = ? GROUP BY s.id"

	row := r.db.QueryRow(query, id)

	err = row.Scan(&countProdBatches.ID, &countProdBatches.SectionNumber, &countProdBatches.ProductsCount)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
		}

		return
	}

	return
}

func (r *SectionRepository) CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error) {
	query := "SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id GROUP BY s.id"

	rows, err := r.db.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var sectionProductBatches model.SectionProductBatches
		err = rows.Scan(&sectionProductBatches.ID, &sectionProductBatches.SectionNumber, &sectionProductBatches.ProductsCount)

		if err != nil {
			return
		}

		countProductBatches = append(countProductBatches, sectionProductBatches)
	}

	return
}
