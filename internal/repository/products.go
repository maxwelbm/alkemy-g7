package repository

import (
	"database/sql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) GetAll() (map[int]model.Product, error) {
	 query := "SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product"
	var products = make(map[int]model.Product)
	rows, err := pr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.ProductCode, &product.Description, 
			&product.Width, &product.Height, &product.Length, &product.NetWeight, 
			&product.ExpirationRate, &product.RecommendedFreezingTemperature, 
			&product.FreezingRate, &product.ProductTypeID, &product.SellerID)
		if err != nil {
			return nil, err
		}
		products[product.ID] = product
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) GetById(id int) (model.Product, error) {
	var product model.Product
	row := pr.DB.QueryRow("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product WHERE id = ?", id)
	err := row.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Width, &product.Height, &product.Length, &product.NetWeight, &product.ExpirationRate, &product.RecommendedFreezingTemperature, &product.FreezingRate, &product.ProductTypeID, &product.SellerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, custom_error.CustomError{Object: id, Err: custom_error.NotFound}
		}
		return product, err
	}

	return product, nil
}

func (pr *ProductRepository) Create(product model.Product) (model.Product, error) {
	result, err := pr.DB.Exec("INSERT INTO product (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 
		product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID)
	if err != nil {
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}
	product.ID = int(id)

	return product, nil
}

func (pr *ProductRepository) Update(id int, product model.Product) (model.Product, error) {
	_, err := pr.DB.Exec("UPDATE product SET product_code = ?, description = ?, width = ?, height = ?, length = ?, net_weight = ?, expiration_rate = ?, recommended_freezing_temperature = ?, freezing_rate = ?, product_type_id = ?, seller_id = ? WHERE id = ?", 
		product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID, id)
	if err != nil {
		return product, err
	}

	product.ID = id
	return product, nil
}

func (pr *ProductRepository) Delete(id int) error {
	_, err := pr.DB.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}