package repository

import (
	"database/sql"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type SectionRepository struct {
	db  *sql.DB
	log logger.Logger
}

func CreateRepositorySections(db *sql.DB, log logger.Logger) *SectionRepository {
	return &SectionRepository{db: db, log: log}
}

func (r *SectionRepository) Get() (sections []model.Section, err error) {
	r.log.Log("SectionRepository", "INFO", "initializing Get function")

	queryGetAll := "SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM `sections`"
	rows, err := r.db.Query(queryGetAll)

	if err != nil {
		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	for rows.Next() {
		var section model.Section
		err = rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)

		if err != nil {
			r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		sections = append(sections, section)
	}

	err = rows.Err()
	if err != nil {
		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	r.log.Log("SectionRepository", "INFO", fmt.Sprintf("returning a slice of sections with no error: %v", sections))

	return
}

func (r *SectionRepository) GetByID(id int) (section model.Section, err error) {
	r.log.Log("SectionRepository", "INFO", "initializing GetByID function with id param")

	queryGetByID := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?"

	row := r.db.QueryRow(queryGetByID, id)

	err = row.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
			r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	r.log.Log("SectionRepository", "INFO", fmt.Sprintf("returning a section from the database based on the id parameter: %v", section))

	return
}

func (r *SectionRepository) Post(section *model.Section) (s model.Section, err error) {
	r.log.Log("SectionRepository", "INFO", "initializing Post function with a section")

	queryPost := "INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(queryPost, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.HandleError("section", customerror.ErrorConflict, "")
		}

		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	(*section).ID = int(id)

	s, _ = r.GetByID(int(id))

	r.log.Log("SectionRepository", "INFO", fmt.Sprintf("saving a section to the database: %v", s))

	return
}

func (r *SectionRepository) Update(id int, section *model.Section) (newSec model.Section, err error) {
	r.log.Log("SectionRepository", "INFO", "initializing Update function with id and section parameters")

	queryUpdate := "UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?"
	_, err = r.db.Exec(queryUpdate, (*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID, id)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
			r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.HandleError("section", customerror.ErrorConflict, "")
			r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	newSec, _ = r.GetByID(id)

	r.log.Log("SectionRepository", "INFO", fmt.Sprintf("updating a section based on the id and section parameter to database: %v", newSec))

	return
}

func (r *SectionRepository) Delete(id int) (err error) {
	r.log.Log("SectionRepository", "INFO", "initializing Delete function with id parameter")

	queryDelete := "DELETE FROM `sections` WHERE `id` = ?"
	_, err = r.db.Exec(queryDelete, id)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
			r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	r.log.Log("SectionRepository", "INFO", "delete a section based on the id parameter from the database")

	return
}

func (r *SectionRepository) CountProductBatchesBySectionID(id int) (countProdBatches model.SectionProductBatches, err error) {
	r.log.Log("SectionRepository", "INFO", "initializing CountProductBatchesBySectionID function with id parameter")

	query := "SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id WHERE s.id = ? GROUP BY s.id"

	row := r.db.QueryRow(query, id)

	err = row.Scan(&countProdBatches.ID, &countProdBatches.SectionNumber, &countProdBatches.ProductsCount)

	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("section", customerror.ErrorNotFound, "")
		}

		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	r.log.Log("SectionRepository", "INFO", fmt.Sprintf("returning a count of product batches by section id: %v", countProdBatches))

	return
}

func (r *SectionRepository) CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error) {
	r.log.Log("SectionRepository", "INFO", "initializing CountProductBatchesSections function")

	query := "SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id GROUP BY s.id"

	rows, err := r.db.Query(query)

	if err != nil {
		r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	defer rows.Close()

	for rows.Next() {
		var sectionProductBatches model.SectionProductBatches
		err = rows.Scan(&sectionProductBatches.ID, &sectionProductBatches.SectionNumber, &sectionProductBatches.ProductsCount)

		if err != nil {
			r.log.Log("SectionRepository", "ERROR", fmt.Sprintf("Error: %v", err))
			return
		}

		countProductBatches = append(countProductBatches, sectionProductBatches)
	}

	r.log.Log("SectionRepository", "INFO", fmt.Sprintf("returning all count of product batches per section: %v", countProductBatches))

	return
}
