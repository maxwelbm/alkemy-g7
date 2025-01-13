package repository

import (
	"database/sql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type ProductRecRepository struct {
	DB *sql.DB
}

// GetById implements interfaces.IProductRecRepository.
func (pr *ProductRecRepository) GetById(id int) (model.ProductRecords, error) {
	panic("unimplemented")
}

func NewProductRecRepository(db *sql.DB) *ProductRecRepository {
	return &ProductRecRepository{DB: db}
}

func (pr *ProductRecRepository) Create(productRec model.ProductRecords) (model.ProductRecords, error) {
	query := `
	INSERT INTO product_records 
	(last_update_date, purchase_price, sale_price, product_id) 
	VALUES (?, ?, ?, ?)
	`

	result, err := pr.DB.Exec(query, productRec.LastUpdateDate,
		productRec.PurchasePrice, productRec.SalePrice, productRec.ProductId)
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

func (pr *ProductRecRepository) GetAll() ([]model.ProductRecordsReport, error) {
	panic("unimplemented")
}
