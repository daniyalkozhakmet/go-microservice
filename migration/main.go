package main

import (
	"fmt"
	"log"
	"microservice/application"
)

var scripts = []string{
	`CREATE TABLE IF NOT EXISTS users (
	id INT PRIMARY KEY AUTO_INCREMENT,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	email VARCHAR(255),
	password VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);`,
	`CREATE TABLE IF NOT EXISTS orders (
	order_id INT PRIMARY KEY AUTO_INCREMENT,
    order_description VARCHAR(255),
    user_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);`,
}

func main() {
	db := application.ConnectDB()
	for _, v := range scripts {
		_, err := db.Exec(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Created table")
}
