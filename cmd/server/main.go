package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ropehapi/api-go-expert/configs"
	"github.com/ropehapi/api-go-expert/internal/entity"
	"github.com/ropehapi/api-go-expert/internal/infra/database"
	"github.com/ropehapi/api-go-expert/internal/infra/webserver/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_api_go_expert?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/product/{id}", productHandler.GetProduct)
	r.Post("/product", productHandler.CreateProduct)
	r.Put("/product/{id}", productHandler.UpdateProduct)
	r.Delete("/product/{id}", productHandler.DeleteProduct)

	r.Post("/user", userHandler.Create)

	http.ListenAndServe(":8000", r)
}
