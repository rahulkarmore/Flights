package routes

import (
	"Flights/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/getFlights", handler.GetFlights())
	
}
