package main

import (
	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func main() {
	r := gin.Default()

	// Ping test
	r.GET("/todobuys", func(c *gin.Context) {
		c.String(200, "['IMAC', 'CELLPHONE']")
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}