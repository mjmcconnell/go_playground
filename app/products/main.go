package products

import (
	"os"
	"database/sql"
	"fmt"
	"log"
	"github.com/mjmcconnell/go_playground/base"
)

type App struct {
	BaseApp *base.App
	DB     *sql.DB
}


func (a *App) setupDB() {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	ssl_mode := os.Getenv("DATABASE_SSL_MODE")

	connectionString :=
		fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=%s", user, password, host, dbname, ssl_mode)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = a.DB.Ping()
	if err != nil {
		fmt.Println(connectionString)
		fmt.Println(err)
	}
}

func Initialize(baseApp *base.App) {
	productApp := App{BaseApp: baseApp}
	productApp.setupDB()
	productApp.initializeRoutes()
}
