package main

import (
	"fmt"
	"go-pet/app/models"
	"go-pet/config"
)

func main() {
	config.Init("development")
	models.Init()
	var _ = models.GetDB()

	fmt.Println("server is running")
}