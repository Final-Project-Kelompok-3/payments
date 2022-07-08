package main

import (
	"flag"
	"os"

	"github.com/Final-Project-Kelompok-3/payments/database"
	"github.com/Final-Project-Kelompok-3/payments/database/migration"
	"github.com/Final-Project-Kelompok-3/payments/database/seeder"
	"github.com/Final-Project-Kelompok-3/payments/internal/factory"
	"github.com/Final-Project-Kelompok-3/payments/internal/http"
	"github.com/Final-Project-Kelompok-3/payments/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
)

// load env file
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	database.CreateConnection()
	xendit.Opt.SecretKey = os.Getenv("XENDIT_KEY")
	

	var m string // for check migration

	flag.StringVar(
		&m,
		"db",
		"run",
		`this argument for check if user want to migrate table, rollback table, or status migration

		to use this flag:
			use -db=migrate for migrate table
			use -db=rollback for rollback table
			use -db=status for get status migration
			use -db=seed for run seeder`,
	)
	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
		return
	} else if m == "rollback" {
		migration.Rollback()
		return
	} else if m == "status" {
		migration.Status()
		return
	} else if m == "seed" {
		seeder.Seed()
		return
	}

	conn := database.GetConnection()
	f := factory.NewFactory(conn)
	e := echo.New()
	http.NewHttp(e, f)
	middleware.Init(e)
	
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}