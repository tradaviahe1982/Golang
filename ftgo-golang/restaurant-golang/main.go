package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"restaurant-golang/controllers"
	"restaurant-golang/initializers"
	"restaurant-golang/routes"
)

var (
	server *gin.Engine

	RestaurantController      controllers.RestaurantController
	RestaurantRouteController routes.RestaurantRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("ðŸš€ Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	RestaurantController = controllers.NewRestaurantController(initializers.DB)
	RestaurantRouteController = routes.NewRouteRestaurantController(RestaurantController)

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

	RestaurantRouteController.RestaurantRoute(router)
	log.Fatal(server.Run(":" + "8000"))
}
