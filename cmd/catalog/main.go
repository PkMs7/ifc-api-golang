package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PkMs7/ifc-api-produtos-golang/internal/database"
	"github.com/PkMs7/ifc-api-produtos-golang/internal/service"
	"github.com/PkMs7/ifc-api-produtos-golang/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ifc")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	//Middlewares
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	//Routes
	c.Get("/category/{id}", webCategoryHandler.GetCategoryHandler)
	c.Get("/category", webCategoryHandler.GetCategoriesHandler)
	c.Post("/category", webCategoryHandler.CreateCategoryHandler)

	c.Get("/product/{id}", webProductHandler.GetProductHandler)
	c.Get("/product", webProductHandler.GetProductsHandler)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductsByCategoryHandler)
	c.Post("/product", webProductHandler.CreateProductHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
