package repository

import (
	"database/sql"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type ProductRepository struct {
	DB  *sql.DB
	log logger.Logger
}

func NewProductRepository(db *sql.DB, log logger.Logger) *ProductRepository {
	return &ProductRepository{DB: db, log: log}
}

func (pr *ProductRepository) GetAll() (map[int]model.Product, error) {
	pr.log.Log("ProductRepository", "INFO", "GetAll function initializing")

	query := "SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products"

	var products = make(map[int]model.Product)

	rows, err := pr.DB.Query(query)

	if err != nil {
		pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error executing query: %v", err))
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
			pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error scanning row: %v", err))
			return nil, err
		}

		products[product.ID] = product
	}

	if err = rows.Err(); err != nil {
		pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error during row iteration: %v", err))
		return nil, err
	}

	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("Retrieved products: %+v", products))
	pr.log.Log("ProductRepository", "INFO", "GetAll function completed")

	return products, nil
}

func (pr *ProductRepository) GetByID(id int) (model.Product, error) {
	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("GetByID function initializing for ID: %d", id))

	var product model.Product

	row := pr.DB.QueryRow("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products WHERE id = ?", id)

	err := row.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Width, &product.Height, &product.Length, &product.NetWeight, &product.ExpirationRate, &product.RecommendedFreezingTemperature, &product.FreezingRate, &product.ProductTypeID, &product.SellerID)

	if err != nil {
		if err == sql.ErrNoRows {
			pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("No product found with ID: %d", id))
			return product, appErr.HandleError("product", appErr.ErrorNotFound, "")
		}
		pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error scanning product with ID %d: %v", id, err))
		return product, err
	}

	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("Retrieved product: %+v", product))

	return product, nil
}

func (pr *ProductRepository) Create(product model.Product) (model.Product, error) {
	pr.log.Log("ProductRepository", "INFO", "Create function initializing")

	result, err := pr.DB.Exec("INSERT INTO products (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID)
	
	if err != nil {
		pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error inserting product: %v", err))
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error retrieving last inserted ID: %v", err))
		return product, err
	}

	product.ID = int(id)
	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("Product created successfully with ID: %d", product.ID))

	return product, nil
}

func (pr *ProductRepository) Update(id int, product model.Product) (model.Product, error) {
	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("Update function initializing for ID: %d", id))

	_, err := pr.DB.Exec("UPDATE products SET product_code = ?, description = ?, width = ?, height = ?, length = ?, net_weight = ?, expiration_rate = ?, recommended_freezing_temperature = ?, freezing_rate = ?, product_type_id = ?, seller_id = ? WHERE id = ?",
		product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID, id)
	
	if err != nil {
		pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error updating product with ID %d: %v", id, err))
		return product, err
	}

	product.ID = id
	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("Product updated successfully with ID: %d", product.ID))

	return product, nil
}

func (pr *ProductRepository) Delete(id int) error {
	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("Delete function initializing for ID: %d", id))

	_, err := pr.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		pr.log.Log("ProductRepository", "ERROR", fmt.Sprintf("Error deleting product with ID %d: %v", id, err))
		return appErr.HandleError("product", appErr.ErrorDep, "product record")
	}

	pr.log.Log("ProductRepository", "INFO", fmt.Sprintf("Product with ID %d deleted successfully", id))
	return nil
}
