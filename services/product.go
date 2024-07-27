package services

import (
	"database/sql"
	"fmt"

	"example.com/go-guid/database"
)

type Product struct {
	id        int
	name      string
	count     int
	check_box bool
}

func CreateProduct(prod *database.Product, db *sql.DB) *error {
	db.Begin()
	_, err := db.Exec(`INSERT INTO products(name,count,check_box)VALUES ($1,$2,$3)`,
		prod.Name, prod.Count, prod.Check_box)
	if err != nil {
		return &err
	}
	return nil
}

func FindOneProduct(db *sql.DB, id *string) (*string, *error) {
	db.Begin()
	rows, err := db.Query(`SELECT * FROM products WHERE id=$1`, id)
	if err != nil {
		return nil, &err
	}
	defer rows.Close()
	p := Product{}
	for rows.Next() {
		err := rows.Scan(&p.id, &p.name, &p.count, &p.check_box)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	if p.id == 0 {
		return nil, nil
	}
	str := fmt.Sprint(p)
	return &str, nil
}
func FindAllProduct(db *sql.DB) ([]string, *error) {
	db.Begin()
	rows, err := db.Query(`SELECT * FROM Products`)
	if err != nil {
		return nil, &err
	}
	defer rows.Close()
	products := []string{}
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.id, &p.name, &p.count, &p.check_box)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, fmt.Sprint(p))
	}
	return products, nil
}
func UpdateProduct(prod *database.Product, db *sql.DB, id *string) *error {
	db.Begin()
	_, err := db.Exec(`update Products set name = $1, count = $2 where id = $3`,
		prod.Name, prod.Count, id)
	if err != nil {
		return &err
	}
	return nil
}
func UpdateCheckBoxProduct(prod *database.Product, db *sql.DB, id *string) *error {
	db.Begin()
	_, err := db.Exec(`update Products set check_box = $1 where id = $2`,
		prod.Check_box, id)
	if err != nil {
		return &err
	}
	return nil
}
func DeleteProduct(db *sql.DB, id *string) *error {
	db.Begin()
	_, err := db.Exec(`delete from Products where id = $1`, id)
	if err != nil {
		return &err
	}
	return nil
}
