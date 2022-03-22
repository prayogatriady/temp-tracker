package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/temp-tracker/controller"
)

var PORT string = "8080"

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/weather/:city", controller.Weather)

	err := router.Run(":" + PORT)
	if err != nil {
		log.Fatalf("Error when running port %s \n", PORT)
	}
}
