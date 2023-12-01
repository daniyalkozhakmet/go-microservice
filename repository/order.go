package repository

import (
	"database/sql"
	"fmt"
	"microservice/model"
)

type OrderRepo struct {
	DB *sql.DB
}

func (o *OrderRepo) Insert() {
	fmt.Println("o", o)
	fmt.Println("Insert")
}
func (o *OrderRepo) Get() ([]model.Order, error) {

	var orders []model.Order

	rows, err := o.DB.Query("SELECT order_id,order_description,created_at,updated_at FROM orders")
	if err != nil {

		return nil, fmt.Errorf("order error: %v", err)
	}
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var ord model.Order
		if err := rows.Scan(&ord.OrderID, &ord.OrderDescription, &ord.CreatedAt, &ord.UpdatedAt); err != nil {
			return nil, fmt.Errorf("order err : %v", err)
		}
		orders = append(orders, ord)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("order err : %v", err)
	}
	return orders, nil
}
func (o *OrderRepo) GetByID() {

}
func (o *OrderRepo) UpdateByID() {

}
func (o *OrderRepo) DeleteByID() {

}
