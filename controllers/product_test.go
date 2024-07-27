package controllers

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestProductRouter(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	r := gin.Default()
	ProductRouter(r, mockDb)
}
func TestCreateProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	r := gin.Default()
	createProduct(r, mockDb)
}
func TestFindOneProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	r := gin.Default()
	findOneProduct(r, mockDb)
}
func TestFindAllProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	r := gin.Default()
	findAllProduct(r, mockDb)
}
func TestUpdateProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	r := gin.Default()
	updateProduct(r, mockDb)
}
func TestUpdateCheckBoxProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	r := gin.Default()
	updateCheckBoxProduct(r, mockDb)
}
func TestDeleteProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	r := gin.Default()
	deleteProduct(r, mockDb)
}
