package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/temp-tracker/service"
)

func Weather(c *gin.Context) {
	city := c.Param("city")
	data, err := service.FindTemp(city)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"name": data.Name,
		"temp": data.Main.Kelvin,
	})
}
