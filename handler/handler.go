package handler

import (
	"Flights/model"
	"Flights/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFlights() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestData := model.Flights{}
		c.Bind(&requestData)
		flightResult := service.GetFlightsService(requestData)
		c.JSON(http.StatusOK, flightResult)
	}
}
