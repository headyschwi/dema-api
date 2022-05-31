package main

import (
	"dema-api/api"
	"dema-api/models"
)

func main() {
	models.DatabaseConnect()
	models.AutoMigrations()
	api.Run()
}
