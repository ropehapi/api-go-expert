package main

import (
	"net/http"

	"github.com/ropehapi/api-go-expert/configs"
	"github.com/ropehapi/api-go-expert/internal/entity"
	"github.com/ropehapi/api-go-expert/internal/infra/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_api_go_expert"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler
