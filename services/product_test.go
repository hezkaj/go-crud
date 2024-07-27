package services

import (
	"testing"

	"example.com/go-guid/database"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateProduct(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO products").
		WithArgs("milk", 2, true).WillReturnResult(sqlmock.NewResult(1, 1))
	CreateProduct(&database.Product{Name: "milk", Count: 2, Check_box: true}, mockDb)
}
func TestFindOneProduct(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	id := "1"
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT * FROM products WHERE id=$1`).WithArgs(1)
	FindOneProduct(mockDb, &id)
}
func TestFindAllProduct(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT * FROM products`)
	FindAllProduct(mockDb)
}
func TestUpdateProduct(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	id := "1"
	mock.ExpectBegin()
	mock.ExpectExec(`update Products set name = $1 count=$2 where id = $3`).
		WithArgs("milk", 2, "1").WillReturnResult(sqlmock.NewResult(1, 1))
	UpdateProduct(&database.Product{Name: "milk", Count: 2, Check_box: true}, mockDb, &id)
}
func TestUpdateCheckBoxProduct(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	id := "1"
	mock.ExpectBegin()
	mock.ExpectExec(`update Products set check_box = $1 where id = $2`).
		WithArgs(true, "1").WillReturnResult(sqlmock.NewResult(1, 1))
	UpdateCheckBoxProduct(&database.Product{Name: "milk", Count: 2, Check_box: true}, mockDb, &id)
}
func TestDeleteProduct(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	id := "1"
	mock.ExpectBegin()
	mock.ExpectExec(`delete from Products where id = 1`)
	DeleteProduct(mockDb, &id)
}
