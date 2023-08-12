package main

import (
	"cc/controllers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_CONN_STR"))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	controller := controllers.NewController(pool)

	e := gin.New()
	v1 := e.Group("/v1")

	v1.GET("/events", controller.EventsLookup)
}
