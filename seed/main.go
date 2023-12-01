package main

import (
	"fmt"
	"microservice/application"
)

type User struct {
	firstName string
	lastName  string
	email     string
	password  string
}
type Order struct {
	CustomerID       int64
	OrderDescription string
}

var orders = []Order{
	{
		CustomerID:       1,
		OrderDescription: "Orderdescription 1",
	},
	{
		CustomerID:       2,
		OrderDescription: "Orderdescription 2",
	},
}

var users = []User{
	{
		firstName: "Daniyal",
		lastName:  "Kozhakmetov",
		email:     "daniyalkozhakmetov@gmail.com",
		password:  "password",
	},
	{
		firstName: "Dummy",
		lastName:  "Dummieov",
		email:     "dummy@gmail.com",
		password:  "password",
	},
}

func main() {
	db := application.ConnectDB()
	for _, v := range users {
		_, err := db.Exec("INSERT INTO users (first_name, last_name, email,password) VALUES (?, ?, ?, ?)", v.firstName, v.lastName, v.email, v.password)
		if err != nil {
			fmt.Printf("users: %v", err)
			return
		}
	}
	for _, v := range orders {
		_, err := db.Exec("INSERT INTO orders (order_description, user_id) VALUES (?, ?)", v.OrderDescription, v.CustomerID)
		if err != nil {
			fmt.Printf("orders: %v", err)
			return
		}
	}

}
