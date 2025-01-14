package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	er "github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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

func (rp *LocalitiesRepository) GetReportCarriersWithId(id int) (locality []model.LocalitiesJSONCarriers, err error) {
	if _, err := rp.GetById(id); err != nil {
		return locality, err
	}

	query := "SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name"
	row := rp.db.QueryRow(query, id)

	var c model.LocalitiesJSONCarriers
	err = row.Scan(&c.ID, &c.Locality, &c.Carriers)
	if errors.Is(err, sql.ErrNoRows) {
		return locality, er.HandleError("locality", er.ErrNotFound, "")
	}

	locality = append(locality, c)
	return
}

func (rp *LocalitiesRepository) GetReportCarriersWithId(id int) (locality []model.LocalitiesJSONCarriers, err error) {
	if _, err := rp.GetById(id); err != nil {
		return locality, err
	}

	query := "SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c INNER JOIN `locality` l ON c.locality_id = l.id WHERE c.locality_id = ? GROUP BY l.id, l.locality_name"
	row := rp.db.QueryRow(query, id)

	var c model.LocalitiesJSONCarriers
	err = row.Scan(&c.ID, &c.Locality, &c.Carriers)
	if errors.Is(err, sql.ErrNoRows) {
		err = model.ErrorLocalityNotFound
		return
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

func (rp *LocalitiesRepository) GetReportSellersWithId(id int) (locality []model.LocalitiesJSONSellers, err error) {
	if _, err := rp.GetById(id); err != nil {
		return locality, err
	}

	query := "SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name"
	row := rp.db.QueryRow(query, id)

	var l model.LocalitiesJSONSellers
	err = row.Scan(&l.ID, &l.Locality, &l.Sellers)
	if errors.Is(err, sql.ErrNoRows) {
		return locality, er.HandleError("locality", er.ErrNotFound, "")
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

func (rp *LocalitiesRepository) GetById(id int) (l model.Locality, err error) {
	query := "SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?"
	row := rp.db.QueryRow(query, id)

	err = row.Scan(&l.ID, &l.Locality, &l.Province, &l.Country)

	if errors.Is(err, sql.ErrNoRows) {
		return l, er.HandleError("locality", er.ErrNotFound, "")
	}
	return
}

func (rp *LocalitiesRepository) CreateLocality(locality *model.Locality) (l model.Locality, err error) {
	query := "INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)"
	result, err := rp.db.Exec(query, (*locality).Locality, (*locality).Province, (*locality).Country)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1064:
				err = er.ErrorInvalidLocalityJSONFormat
			case 1048:
				err = er.ErrorNullLocalityAttribute
			}
			return
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	l, _ = rp.GetById(int(id))

	return
}
