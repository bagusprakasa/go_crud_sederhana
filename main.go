package main

import (
	"crud_sederhana/models"
	"crud_sederhana/routes"
)

func main() {
	db := models.SetupDB()
	// Migrate Table From Models
	db.AutoMigrate(&models.User{})

	router := routes.SetupRoutes(db)
	router.Run()
}
