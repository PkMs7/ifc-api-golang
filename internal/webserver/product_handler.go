package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/PkMs7/ifc-api-produtos-golang/internal/entity"
	"github.com/PkMs7/ifc-api-produtos-golang/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProdtuctService *service.ProdtuctService
}

func NewWebProductHandler(productService *service.ProdtuctService) *WebProductHandler {
	return &WebProductHandler{ProdtuctService: productService}
}

func (wph *WebProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProdtuctService.GetProductsService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	product, err := wph.ProdtuctService.GetProductService(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProductsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	products, err := wph.ProdtuctService.GetProductsByCategoryService(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wph.ProdtuctService.CreateProductService(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
