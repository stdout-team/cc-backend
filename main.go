package main

import (
	"cc/controllers"
	"cc/services"
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

	service := services.NewEventsService(pool)

	mainController := controllers.NewController(service)

	e := gin.New()
	v1 := e.Group("/v1").Use(controllers.CorsMiddleware()).Use(controllers.ErrorsMiddleware)

	v1.GET("/events", mainController.EventsLookup)
	v1.GET("/getCoordinates", controllers.ResolveCoordinates)
	v1.POST("/countMeIn/:id", mainController.CountMeIn)
	v1.GET("/interests", mainController.GetAllInterests)
	v1.POST("/addEvent", mainController.AddEvent)

	e.Run(":8084")
}
