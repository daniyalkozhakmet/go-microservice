package main

import (
	"context"
	"fmt"
	"log"
	"microservice/application"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func init() {
	LoadEnvVariables()
}

func main() {

	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}

}
