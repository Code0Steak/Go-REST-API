package main

import (
	"database/sql"
	"errors"
	"fmt"
)

//This file will contain all methods related to DB querying

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// getProductsFromDB func gets all products from DB. as per it's name suggests :)
func getProductsFromDB(db *sql.DB) ([]product, error) {
	query := "SELECT id, name, quantity, price FROM products"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	var products []product

	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)

		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil

}

func (p *product) getProductFromDB(db *sql.DB) error {

	query := fmt.Sprintf("SELECT name, quantity, price FROM products WHERE id=%d", p.ID)
	row := db.QueryRow(query)
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)

	if err != nil {
		return err
	}
	return nil
}

func (p *product) createProductinDB(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v(name,quantity,price) values(\"%v\", %v, %v);", DB_Table, p.Name, p.Quantity, p.Price)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)

	return nil
}

func (p *product) updateProduct(db *sql.DB) error {
	query := fmt.Sprintf("UPDATE products SET name=\"%v\", quantity=%v, price=%v  WHERE id=%v", p.Name, p.Quantity, p.Price, p.ID)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("No such row exists.")
	}
	return err
}

func (p *product) deleteProduct(db *sql.DB) error {

	query := fmt.Sprintf("DELETE from products WHERE id=%v", p.ID)
	result, err := db.Exec(query)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Nothing changed, as ID is incorrect.")
	}

	return nil
}
