package middleware

import (
	"Flights/routes"
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func IintMiddleware(g *gin.Engine) {

	g.Use(cors.Default())

	r := g.Group("/r")
	r.Use(RestrictedRequest())

	routes.InitRoutes(r)

}

func RestrictedRequest() gin.HandlerFunc {
	dbToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlJhaHVsIiwicGFzc3dvcmQiOiJSQGh1bEsiLCJ0b2tlbiI6IiIsImV4cCI6MTYyNDk1NzM5NH0.fA10RzTginH-GPRAam-d1JeUPbCxUfZfigIy0yJJA9M"
	return func(c *gin.Context) {
		fmt.Println("header data :", c.GetHeader("Authorization"))

		token := c.GetHeader("Authorization")

		if strings.Trim(token, "") == "" {
			fmt.Println("Token not available")
			c.AbortWithStatusJSON(401, gin.H{"error": "Token not available"})
		}

		if token != "" && token == dbToken {
			c.Next()

		} else {
			fmt.Println("Unauthorised Access")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorised Access"})
		}

	}
}
