package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func New() *App {
	app := &App{
		DB: ConnectDB(),
	}
	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {

	err := a.Router.Run()

	if err != nil {
		return fmt.Errorf("failed to start server: %w ", err)
	}

	fmt.Println("Server running")
	return nil
}
func ConnectDB() *sql.DB {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "golang",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Here Failed connect to DB")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("There", pingErr)
	}
	fmt.Println("Connected!")

	return db
}
