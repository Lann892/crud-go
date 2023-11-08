package main

import (
	"go-backend/config"
	"go-backend/models"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Person{})
}
