package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

func CreateRepositoryLocalities(db *sql.DB) *LocalitiesRepository {
	return &LocalitiesRepository{db}
}

type LocalitiesRepository struct {
	db *sql.DB
}

// GetCarries implements interfaces.ILocalityRepo.
func (rp *LocalitiesRepository) GetCarries(id int) (locality model.LocalitiesJSONCarries, err error) {
	panic("unimplemented")
}

// GetSellers implements interfaces.ILocalityRepo.
func (rp *LocalitiesRepository) GetSellers(id int) (locality model.LocalitiesJSONSellers, err error) {
	panic("unimplemented")
}

func (rp *LocalitiesRepository) Get() (localities []model.Locality, err error) {
	query := "SELECT `id`, `locality_name`, `province_id` FROM `locality`"
	rows, err := rp.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var locality model.Locality
		err = rows.Scan(&locality.ID, &locality.Locality, &locality.Province)
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
	query := "SELECT `id`, `locality_name`, `province_id` FROM `locality` WHERE `id` = ?"
	row := rp.db.QueryRow(query, id)

	err = row.Scan(&l.ID, &l.Locality, &l.Province)

	if errors.Is(err, sql.ErrNoRows) {
		err = model.ErrorLocalityNotFound
		return
	}
	return
}

func (rp *LocalitiesRepository) Post(locality *model.Locality) (l model.Locality, err error) {
	query := "INSERT INTO `locality` (`id`, `locality_name`, `province_id`) VALUES (?, ?, ?)"
	result, err := rp.db.Exec(query, (*locality).ID, (*locality).Locality, (*locality).Province)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = model.ErrorIDAlreadyExist
			case 1064:
				err = model.ErrorInvalidLocalityJSONFormat
			case 1048:
				err = model.ErrorNullLocalityAttribute
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
