package repository

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type ProductRecRepository struct {
	DB  *sql.DB
	log logger.Logger
}

func NewProductRecRepository(db *sql.DB, log logger.Logger) *ProductRecRepository {
	return &ProductRecRepository{DB: db, log: log}
}

func (pr *ProductRecRepository) Create(productRec model.ProductRecords) (model.ProductRecords, error) {
	query := `
	INSERT INTO product_records 
	(last_update_date, purchase_price, sale_price, product_id) 
	VALUES (?, ?, ?, ?)
	`

	result, err := pr.DB.Exec(query, productRec.LastUpdateDate,
		productRec.PurchasePrice, productRec.SalePrice, productRec.ProductID)

	if err != nil {
		return model.ProductRecords{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return model.ProductRecords{}, err
	}

	productRec.ID = int(id)

	return productRec, nil
}

func (pr *ProductRecRepository) GetByID(id int) (model.ProductRecords, error) {
	var productRecord model.ProductRecords

	query := `
	SELECT
	id,
	last_update_date, 
	product_id, 
	purchase_price, 
	sale_price
	FROM product_records 
	WHERE id = ?
	`
	row := pr.DB.QueryRow(query, id)

	err := row.Scan(&productRecord.ID, &productRecord.LastUpdateDate,
		&productRecord.ProductID, &productRecord.PurchasePrice, &productRecord.SalePrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ProductRecords{}, appErr.HandleError("product record", appErr.ErrorNotFound, "")
		}

		return model.ProductRecords{}, err
	}

	return productRecord, nil
}

func (pr *ProductRecRepository) GetAll() ([]model.ProductRecords, error) {
	var productRecordList []model.ProductRecords

	query := `
	SELECT
	id,
	last_update_date, 
	product_id, 
	purchase_price, 
	sale_price
	FROM product_records
	`

	rows, err := pr.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var productRecord model.ProductRecords

		err := rows.Scan(&productRecord.ID, &productRecord.LastUpdateDate,
			&productRecord.ProductID, &productRecord.PurchasePrice,
			&productRecord.SalePrice)

		if err != nil {
			return nil, errors.New("error trying to convert row to struct")
		}

		productRecordList = append(productRecordList, productRecord)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return productRecordList, nil
}

func (pr *ProductRecRepository) GetByIDProduct(idProduct int) ([]model.ProductRecords, error) {
	var productRecordList []model.ProductRecords

	query := `
	SELECT
	id,
	last_update_date, 
	product_id, 
	purchase_price, 
	sale_price
	FROM product_records
	where product_id = ?
	`

	rows, err := pr.DB.Query(query, idProduct)

	if err != nil {
		return productRecordList, err
	}

	defer rows.Close()

	for rows.Next() {
		var productRecord model.ProductRecords
		err := rows.Scan(&productRecord.ID, &productRecord.LastUpdateDate,
			&productRecord.ProductID, &productRecord.PurchasePrice,
			&productRecord.SalePrice)

		if err != nil {
			return nil, errors.New("error trying to convert row to struct")
		}

		productRecordList = append(productRecordList, productRecord)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return productRecordList, nil
}

func (pr *ProductRecRepository) GetAllReport() ([]model.ProductRecordsReport, error) {
	var productRecordReport []model.ProductRecordsReport

	query := `
	SELECT  
	p.id, 
	p.description, 
	count(p.id) as record_count 
	FROM products p
	inner join product_records pr on pr.product_id = p.id
	GROUP by p.id, p.description
	`
	rows, err := pr.DB.Query(query)

	if err != nil {
		return productRecordReport, err
	}

	defer rows.Close()

	for rows.Next() {
		var productRecord model.ProductRecordsReport

		err := rows.Scan(&productRecord.ProductID, &productRecord.Description, &productRecord.RecordsCount)

		if err != nil {
			return nil, errors.New("error trying to convert row to struct")
		}

		productRecordReport = append(productRecordReport, productRecord)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return productRecordReport, nil
}
