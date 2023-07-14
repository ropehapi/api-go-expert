package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ropehapi/api-go-expert/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_api_go_expert"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)

	assert.NoError(t, err)
	assert.NotEqual(t, product.ID, "")
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_api_go_expert"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i := 1; 1 < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)

		db.Create(product)
	}
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, productDB, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, productDB, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, productDB, 10)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[9].Name)
}

func TestFindProductById(t *testing.T){
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_api_go_expert"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T){
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_api_go_expert"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)
	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T){
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_api_go_expert"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	_, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
} 
