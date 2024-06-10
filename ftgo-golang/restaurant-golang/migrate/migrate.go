package main

import (
	"fmt"
	"log"
	"restaurant-golang/initializers"
	"restaurant-golang/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.Restaurant{}, &models.Menu{},
		&models.MenuItem{}, &models.Ticket{})
	fmt.Println("👍 Migration complete")
}
