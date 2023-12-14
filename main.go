package main

import (
	// book adalah directory root project go yang kita buat
	"crud_go/models" // memanggil package models pada directory models
	"crud_go/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run()
}
