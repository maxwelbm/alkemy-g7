package repository

import (
	"database/sql"
	"errors"
	"fmt"

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
	pr.log.Log("ProductRecRepository", "INFO", "Create function initializing")

	query := `
	INSERT INTO product_records 
	(last_update_date, purchase_price, sale_price, product_id) 
	VALUES (?, ?, ?, ?)
	`

	result, err := pr.DB.Exec(query, productRec.LastUpdateDate,
		productRec.PurchasePrice, productRec.SalePrice, productRec.ProductID)

	if err != nil {
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error inserting product record: %v", err))
		return model.ProductRecords{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error retrieving last inserted ID: %v", err))
		return model.ProductRecords{}, err
	}

	productRec.ID = int(id)
	pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("Product record created successfully with ID: %d", productRec.ID))

	return productRec, nil
}

func (pr *ProductRecRepository) GetByID(id int) (model.ProductRecords, error) {
	pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("GetByID function initializing for ID: %d", id))
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
			pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("No product record found with ID: %d", id))
			return model.ProductRecords{}, appErr.HandleError("product record", appErr.ErrorNotFound, "")
		}
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error scanning product record with ID %d: %v", id, err))
		return model.ProductRecords{}, err
	}

	pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("Retrieved product record: %+v", productRecord))

	return productRecord, nil
}

func (pr *ProductRecRepository) GetAll() ([]model.ProductRecords, error) {
	pr.log.Log("ProductRecRepository", "INFO", "GetAll function initializing")
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
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error executing query: %v", err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productRecord model.ProductRecords
		err := rows.Scan(&productRecord.ID, &productRecord.LastUpdateDate,
			&productRecord.ProductID, &productRecord.PurchasePrice,
			&productRecord.SalePrice)

		if err != nil {
			pr.log.Log("ProductRecRepository", "ERROR", "Error trying to convert row to struct")
			return nil, errors.New("Error trying to convert row to struct")
		}

		productRecordList = append(productRecordList, productRecord)
	}

	if err = rows.Err(); err != nil {
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error during row iteration: %v", err))
		return nil, err
	}

	pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("Retrieved all product records: %+v", productRecordList))
	return productRecordList, nil
}

func (pr *ProductRecRepository) GetByIDProduct(idProduct int) ([]model.ProductRecords, error) {
	pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("GetByIDProduct function initializing for Product ID: %d", idProduct))
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
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error executing query for product ID %d: %v", idProduct, err))
		return productRecordList, err
	}
	defer rows.Close()

	for rows.Next() {
		var productRecord model.ProductRecords
		err := rows.Scan(&productRecord.ID, &productRecord.LastUpdateDate,
			&productRecord.ProductID, &productRecord.PurchasePrice,
			&productRecord.SalePrice)

		if err != nil {
			pr.log.Log("ProductRecRepository", "ERROR", "Error trying to convert row to struct")
			return nil, errors.New("Error trying to convert row to struct")
		}

		productRecordList = append(productRecordList, productRecord)
	}

	if err = rows.Err(); err != nil {
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error during row iteration: %v", err))
		return nil, err
	}

	pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("Retrieved product records for Product ID %d: %+v", idProduct, productRecordList))
	return productRecordList, nil
}

func (pr *ProductRecRepository) GetAllReport() ([]model.ProductRecordsReport, error) {
	pr.log.Log("ProductRecRepository", "INFO", "GetAllReport function initializing")
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
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error executing report query: %v", err))
		return productRecordReport, err
	}
	defer rows.Close()

	for rows.Next() {
		var productRecord model.ProductRecordsReport
		err := rows.Scan(&productRecord.ProductID, &productRecord.Description, &productRecord.RecordsCount)

		if err != nil {
			pr.log.Log("ProductRecRepository", "ERROR", "Error trying to convert row to struct")
			return nil, errors.New("Error trying to convert row to struct")
		}

		productRecordReport = append(productRecordReport, productRecord)
	}

	if err = rows.Err(); err != nil {
		pr.log.Log("ProductRecRepository", "ERROR", fmt.Sprintf("Error during row iteration: %v", err))
		return nil, err
	}

	pr.log.Log("ProductRecRepository", "INFO", fmt.Sprintf("Retrieved all product record reports: %+v", productRecordReport))
	return productRecordReport, nil
}