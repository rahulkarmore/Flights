package main

import (
	"Flights/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.GET("/", checkServer())

	middleware.IintMiddleware(route)
	serve := &http.Server{
		Addr:    ":8081",
		Handler: route,
	}
	fmt.Println("server running...")
	serve.ListenAndServe()
}

func checkServer() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running")

	}

}
