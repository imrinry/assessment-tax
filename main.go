package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/imrinry/assessment-tax/handlers"
	"github.com/imrinry/assessment-tax/logs"
	"github.com/imrinry/assessment-tax/middlewares"
	"github.com/imrinry/assessment-tax/repositories"
	"github.com/imrinry/assessment-tax/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func InitializeDatabase(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS deductions (
		id SERIAL PRIMARY KEY,
		name TEXT,
		value FLOAT
	)
	`)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT INTO deductions (name, value) VALUES
	('personalDeduction', 0),
	('kReceipt', 0)
	ON CONFLICT DO NOTHING;
`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		logs.Error("PORT must be set")
		panic("PORT must be set")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logs.Error("DATABASE_URL must be set")
		panic("DATABASE_URL must be set")
	}

	db, err := InitializeDatabase(dbURL)
	if err != nil {
		panic(err)
	}

	repo := repositories.New(db)
	services := services.New(repo)
	handlers := handlers.New(services)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.Group("/admin").Use(middlewares.BasicAuthMiddleware)
	e.POST("/admin/deductions/personal", handlers.DeductionPersonalSetting)
	e.POST("/admin/deductions/k-receipt", handlers.DeductionkReceiptSetting)

	e.POST("/tax/calculations", handlers.TaxCalculations)
	e.POST("tax/calculations/upload-csv", handlers.CsvFileTaxCalculations)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		fmt.Println("\n Shutting down the server...")
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	e.Start(fmt.Sprintf(":%s", port))

}
