package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	er "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

func CreateRepositoryLocalities(db *sql.DB, log logger.Logger) *LocalitiesRepository {
	return &LocalitiesRepository{db, log}
}

type LocalitiesRepository struct {
	db  *sql.DB
	log logger.Logger
}

func (rp *LocalitiesRepository) GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error) {
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Carriers function initializing")

	query := "SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name"
	rows, err := rp.db.Query(query)

	if err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	defer rows.Close()

	for rows.Next() {
		var c model.LocalitiesJSONCarriers
		err = rows.Scan(&c.ID, &c.Locality, &c.Carriers)

		if err != nil {
			rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		report = append(report, c)
	}

	rp.log.Log("LocalitiesRepository", "INFO", fmt.Sprintf("Retrieved report carriers: %+v", report))
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Carriers function completed")

	return
}

func (rp *LocalitiesRepository) GetReportCarriersWithID(id int) (locality []model.LocalitiesJSONCarriers, err error) {
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Carrier by ID function initializing")

	if _, err := rp.GetByID(id); err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return locality, err
	}

	query := "SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name"
	row := rp.db.QueryRow(query, id)

	var c model.LocalitiesJSONCarriers
	err = row.Scan(&c.ID, &c.Locality, &c.Carriers)

	if err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	locality = append(locality, c)

	rp.log.Log("LocalitiesRepository", "INFO", fmt.Sprintf("Retrieved report carrier: %+v", locality))
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Carrier by ID function completed")

	return
}

func (rp *LocalitiesRepository) GetSellers(id int) (report []model.LocalitiesJSONSellers, err error) {
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Sellers function initializing")

	query := "SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name"
	rows, err := rp.db.Query(query)

	if err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	defer rows.Close()

	for rows.Next() {
		var l model.LocalitiesJSONSellers
		err = rows.Scan(&l.ID, &l.Locality, &l.Sellers)

		if err != nil {
			rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		report = append(report, l)
	}

	rp.log.Log("LocalitiesRepository", "INFO", fmt.Sprintf("Retrieved report sellers: %+v", report))
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Sellers function completed")

	return
}

func (rp *LocalitiesRepository) GetReportSellersWithID(id int) (locality []model.LocalitiesJSONSellers, err error) {
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Seller by ID function initializing")

	if _, err := rp.GetByID(id); err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return locality, err
	}

	query := "SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name"
	row := rp.db.QueryRow(query, id)

	var s model.LocalitiesJSONSellers
	err = row.Scan(&s.ID, &s.Locality, &s.Sellers)

	if err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	locality = append(locality, s)

	rp.log.Log("LocalitiesRepository", "INFO", fmt.Sprintf("Retrieved report seller: %+v", locality))
	rp.log.Log("LocalitiesRepository", "INFO", "Get report Seller by ID function completed")

	return
}

func (rp *LocalitiesRepository) Get() (localities []model.Locality, err error) {
	rp.log.Log("LocalitiesRepository", "INFO", "Get localities function initializing")

	query := "SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality`"
	rows, err := rp.db.Query(query)

	if err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	defer rows.Close()

	for rows.Next() {
		var locality model.Locality
		err = rows.Scan(&locality.ID, &locality.Locality, &locality.Province, &locality.Country)

		if err != nil {
			rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		localities = append(localities, locality)
	}

	rp.log.Log("LocalitiesRepository", "INFO", fmt.Sprintf("Retrieved localities: %+v", localities))
	rp.log.Log("LocalitiesRepository", "INFO", "Get localities function completed")

	return
}

func (rp *LocalitiesRepository) GetByID(id int) (l model.Locality, err error) {
	rp.log.Log("LocalitiesRepository", "INFO", "Get locality by ID function initializing")

	query := "SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?"
	row := rp.db.QueryRow(query, id)

	err = row.Scan(&l.ID, &l.Locality, &l.Province, &l.Country)

	if errors.Is(err, sql.ErrNoRows) {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		e := er.ErrLocalityNotFound

		return l, e
	}

	rp.log.Log("LocalitiesRepository", "INFO", fmt.Sprintf("Retrieved locality: %+v", l))
	rp.log.Log("LocalitiesRepository", "INFO", "Get locality by ID function completed")

	return
}

func (rp *LocalitiesRepository) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	rp.log.Log("LocalitiesRepository", "INFO", "Create locality function initializing")

	query := "INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)"
	result, err := rp.db.Exec(query, (*locality).Locality, (*locality).Province, (*locality).Country)
	err = rp.validateSQLError(err)

	if err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		rp.log.Log("LocalitiesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	l, _ = rp.GetByID(int(id))

	rp.log.Log("LocalitiesRepository", "INFO", fmt.Sprintf("Created locality: %+v", l))
	rp.log.Log("LocalitiesRepository", "INFO", "Create locality function completed")

	return
}

func (rp *LocalitiesRepository) validateSQLError(err error) (e error) {
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1064:
				e = er.ErrInvalidLocalityJSONFormat
			case 1048:
				e = er.ErrNullLocalityAttribute
			default:
				e = er.ErrDefaultLocalitySQL
			}
		}
	}

	return e
}
