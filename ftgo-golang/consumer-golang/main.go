package main

import (
	"consumer-golang/controllers"
	"consumer-golang/initializers"
	"consumer-golang/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine

	ConsumerController      controllers.ConsumerController
	ConsumerRouteController routes.ConsumerRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("ðŸš€ Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	ConsumerController = controllers.NewConsumerController(initializers.DB)
	ConsumerRouteController = routes.NewRouteConsumerController(ConsumerController)

	server = gin.Default()
}

func main() {
	_, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ConsumerRouteController.ConsumerRoute(router)
	log.Fatal(server.Run(":" + "8000"))
}
