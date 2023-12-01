package repository

import (
	"database/sql"
	"fmt"
	"microservice/model"
)

type BasicRepo struct {
	DB *sql.DB
}

func (o *BasicRepo) Insert() {
	fmt.Println("o", o)
	fmt.Println("Insert")
}
func (o *BasicRepo) Get() ([]model.Order, error) {

	var orders []model.Order

	return orders, nil
}
