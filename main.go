package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/imrinry/assessment-tax/logs"
	"github.com/imrinry/assessment-tax/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func initDB(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		logs.Error(err)
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

	db, err := initDB(dbURL)
	if err != nil {
		panic(err)
	}

	repo := repositories.New(db)
	_, err = repo.ExamRepo(context.Background())
	fmt.Println("err", err)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

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
