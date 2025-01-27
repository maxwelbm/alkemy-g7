package service

import (
	"reflect"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

type ProductService struct {
	ProductRepository interfaces.IProductsRepo
	SellerRepository  interfaces.ISellerRepo
}

func NewProductService(productRepo interfaces.IProductsRepo, sellerRepo interfaces.ISellerRepo) *ProductService {
	return &ProductService{
		ProductRepository: productRepo,
		SellerRepository:  sellerRepo,
	}
}

func (ps *ProductService) GetAllProducts() ([]model.Product, error) {
	products, err := ps.ProductRepository.GetAll()

	//nolint:prealloc
	var productSlice []model.Product

	if err != nil {
		return productSlice, err
	}

	for _, product := range products {
		productSlice = append(productSlice, product)
	}

	return productSlice, nil
}

func (ps *ProductService) GetProductByID(id int) (model.Product, error) {
	product, err := ps.ProductRepository.GetByID(id)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (ps *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	err := product.Validate()

	if err != nil {
		return model.Product{}, err
	}

	_, err = ps.SellerRepository.GetByID(product.SellerID)
	if err != nil {
		return model.Product{}, err
	}

	productsList, _ := ps.ProductRepository.GetAll()

	existsByCode := existsByProductCode(product.ProductCode, productsList)

	if existsByCode {
		return model.Product{}, customError.CustomError{Object: product.ProductCode, Err: customError.ErrConflict}
	}

	productDB, err := ps.ProductRepository.Create(product)

	if err != nil {
		return model.Product{}, err
	}

	return productDB, nil
}

func (ps *ProductService) UpdateProduct(id int, product model.Product) (model.Product, error) {
	if product.SellerID != 0 {
		_, err := ps.SellerRepository.GetByID(product.SellerID)

		if err != nil {
			return model.Product{}, err
		}
	}

	listOfProducts, _ := ps.ProductRepository.GetAll()

	if existsByProductCode(product.ProductCode, listOfProducts) {
		return model.Product{}, customError.CustomError{Object: product.ProductCode, Err: customError.ErrConflict}
	}

	productInDB, err := ps.ProductRepository.GetByID(id)

	if err != nil {
		return model.Product{}, err
	}

	productAdjusted := updateProduct(productInDB, product)

	productUpdated, _ := ps.ProductRepository.Update(id, productAdjusted)

	return productUpdated, nil
}

func (ps *ProductService) DeleteProduct(id int) error {
	_, err := ps.ProductRepository.GetByID(id)

	if err != nil {
		return customError.HandleError("product", customError.ErrorNotFound, "")
	}

	err = ps.ProductRepository.Delete(id)

	if err != nil {
		return err
	}

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
