package service

import (
	"fmt"
	"reflect"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	customerror "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type ProductService struct {
	ProductRepository interfaces.IProductsRepo
	SellerRepository  interfaces.ISellerRepo
	log               logger.LoggerDefault
}

func NewProductService(productRepo interfaces.IProductsRepo, sellerRepo interfaces.ISellerRepo, logger logger.LoggerDefault) *ProductService {
	return &ProductService{
		ProductRepository: productRepo,
		SellerRepository:  sellerRepo,
		log:               logger,
	}
}

func (ps *ProductService) GetAllProducts() ([]model.Product, error) {
	ps.log.Log("ProductService", "INFO", "GetAllProducts function initializing")

	products, err := ps.ProductRepository.GetAll()

	var productSlice []model.Product

	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Error retrieving all products: %v", err))
		return productSlice, err
	}

	for _, product := range products {
		productSlice = append(productSlice, product)
	}

	ps.log.Log("ProductService", "INFO", fmt.Sprintf("Retrieved all products count: %d", len(productSlice)))
	return productSlice, nil
}

func (ps *ProductService) GetProductByID(id int) (model.Product, error) {
	ps.log.Log("ProductService", "INFO", fmt.Sprintf("GetProductByID function initializing for ID: %d", id))

	product, err := ps.ProductRepository.GetByID(id)
	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Error retrieving product with ID: %d, error: %v", id, err))
		return model.Product{}, err
	}

	ps.log.Log("ProductService", "INFO", fmt.Sprintf("Retrieved product: %+v", product))
	return product, nil
}

func (ps *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	ps.log.Log("ProductService", "INFO", "CreateProduct function initializing")

	err := product.Validate()
	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Validation error: %v", err))
		return model.Product{}, err
	}

	_, err = ps.SellerRepository.GetByID(product.SellerID)
	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Seller not found with ID: %d", product.SellerID))
		return model.Product{}, err
	}

	productsList, _ := ps.ProductRepository.GetAll()
	existsByCode := existsByProductCode(product.ProductCode, productsList)

	if existsByCode {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Product code already exists: %s", product.ProductCode))
		return model.Product{}, customerror.CustomError{Object: product.ProductCode, Err: customerror.ErrConflict}
	}

	productDB, err := ps.ProductRepository.Create(product)
	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Error creating product: %v", err))
		return model.Product{}, err
	}

	ps.log.Log("ProductService", "INFO", fmt.Sprintf("Product created successfully: %+v", productDB))
	return productDB, nil
}

func (ps *ProductService) UpdateProduct(id int, product model.Product) (model.Product, error) {
	ps.log.Log("ProductService", "INFO", fmt.Sprintf("UpdateProduct function initializing for ID: %d", id))

	if product.SellerID != 0 {
		_, err := ps.SellerRepository.GetByID(product.SellerID)
		if err != nil {
			ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Seller not found with ID: %d", product.SellerID))
			return model.Product{}, err
		}
	}

	listOfProducts, _ := ps.ProductRepository.GetAll()
	if existsByProductCode(product.ProductCode, listOfProducts) {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Product code already exists conflicting during update: %s", product.ProductCode))
		return model.Product{}, customerror.CustomError{Object: product.ProductCode, Err: customerror.ErrConflict}
	}

	productInDB, err := ps.ProductRepository.GetByID(id)
	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Error retrieving product with ID: %d, error: %v", id, err))
		return model.Product{}, err
	}

	productAdjusted := updateProduct(productInDB, product)
	productUpdated, err := ps.ProductRepository.Update(id, productAdjusted)

	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Error updating product: %v", err))
		return model.Product{}, err
	}

	ps.log.Log("ProductService", "INFO", fmt.Sprintf("Product updated successfully: %+v", productUpdated))
	return productUpdated, nil
}

func (ps *ProductService) DeleteProduct(id int) error {
	ps.log.Log("ProductService", "INFO", fmt.Sprintf("DeleteProduct function initializing for ID: %d", id))

	_, err := ps.ProductRepository.GetByID(id)
	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Error retrieving product for deletion with ID: %d, error: %v", id, err))
		return customerror.HandleError("product", customerror.ErrorNotFound, "")
	}

	err = ps.ProductRepository.Delete(id)
	if err != nil {
		ps.log.Log("ProductService", "ERROR", fmt.Sprintf("Error deleting product with ID: %d, error: %v", id, err))
		return err
	}

	ps.log.Log("ProductService", "INFO", fmt.Sprintf("Product with ID: %d deleted successfully", id))
	return nil
}

func existsByProductCode(productCode string, products map[int]model.Product) bool {
	for _, product := range products {
		if product.ProductCode == productCode {
			return true
		}
	}
	return false
}

func updateProduct(existingProduct model.Product, newProduct model.Product) model.Product {
	valueOfNewProduct := reflect.ValueOf(newProduct)
	valueOfExistingProduct := reflect.ValueOf(&existingProduct).Elem()

	for i := 0; i < valueOfNewProduct.NumField(); i++ {
		newValue := valueOfNewProduct.Field(i)

		if !newValue.IsZero() {
			valueOfExistingProduct.Field(i).Set(newValue)
		}
	}

	return existingProduct
}
