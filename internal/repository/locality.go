package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	er "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

func CreateRepositoryLocalities(db *sql.DB) *LocalitiesRepository {
	return &LocalitiesRepository{db}
}

type LocalitiesRepository struct {
	db *sql.DB
}

func (rp *LocalitiesRepository) GetCarriers(id int) (report []model.LocalitiesJSONCarriers, err error) {
	query := "SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name"
	rows, err := rp.db.Query(query)

	if err != nil {
		return
	}

	for rows.Next() {
		var c model.LocalitiesJSONCarriers
		err = rows.Scan(&c.ID, &c.Locality, &c.Carriers)

		if err != nil {
			return
		}

		report = append(report, c)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (rp *LocalitiesRepository) GetReportCarriersWithID(id int) (locality []model.LocalitiesJSONCarriers, err error) {
	if _, err := rp.GetByID(id); err != nil {
		return locality, err
	}

	query := "SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name"
	row := rp.db.QueryRow(query, id)

	var c model.LocalitiesJSONCarriers
	err = row.Scan(&c.ID, &c.Locality, &c.Carriers)

	if errors.Is(err, sql.ErrNoRows) {
		e := er.ErrLocalityNotFound
		return locality, e
	}

	locality = append(locality, c)

	return
}

func (rp *LocalitiesRepository) GetSellers(id int) (report []model.LocalitiesJSONSellers, err error) {
	query := "SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name"
	rows, err := rp.db.Query(query)

	if err != nil {
		return
	}

	for rows.Next() {
		var l model.LocalitiesJSONSellers
		err = rows.Scan(&l.ID, &l.Locality, &l.Sellers)

		if err != nil {
			return
		}

		report = append(report, l)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (rp *LocalitiesRepository) GetReportSellersWithID(id int) (locality []model.LocalitiesJSONSellers, err error) {
	if _, err := rp.GetByID(id); err != nil {
		return locality, err
	}

	query := "SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name"
	row := rp.db.QueryRow(query, id)

	var l model.LocalitiesJSONSellers
	err = row.Scan(&l.ID, &l.Locality, &l.Sellers)

	if err != nil {
		return
	}

	locality = append(locality, l)

	return
}

func (rp *LocalitiesRepository) Get() (localities []model.Locality, err error) {
	query := "SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality`"
	rows, err := rp.db.Query(query)

	if err != nil {
		return
	}

	for rows.Next() {
		var locality model.Locality
		err = rows.Scan(&locality.ID, &locality.Locality, &locality.Province, &locality.Country)

		if err != nil {
			return
		}

		localities = append(localities, locality)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (rp *LocalitiesRepository) GetByID(id int) (l model.Locality, err error) {
	query := "SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?"
	row := rp.db.QueryRow(query, id)

	err = row.Scan(&l.ID, &l.Locality, &l.Province, &l.Country)

	if errors.Is(err, sql.ErrNoRows) {
		e := er.ErrLocalityNotFound
		return l, e
	}

	return
}

func (rp *LocalitiesRepository) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	query := "INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)"
	result, err := rp.db.Exec(query, (*locality).Locality, (*locality).Province, (*locality).Country)
	err = rp.validateSQLError(err)

	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	l, _ = rp.GetByID(int(id))

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
