package service

import (
	"github.com/PkMs7/ifc-api-produtos-golang/internal/database"
	"github.com/PkMs7/ifc-api-produtos-golang/internal/entity"
)

type ProdtuctService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProdtuctService {
	return &ProdtuctService{ProductDB: productDB}
}

func (ps *ProdtuctService) GetProductsService() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProdtuctService) GetProductService(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProdtuctService) GetProductsByCategoryService(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductsByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProdtuctService) CreateProductService(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryID, imageURL, price)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
