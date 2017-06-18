package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

var DB = make(map[string]string)

func main() {
    r := gin.Default()

    // Ping test
    r.GET("/todobuys", func(c *gin.Context) {
        c.String(http.StatusOK, "['IMAC', 'CELLPHONE']")
    })

    r.GET("/todobuys/:Id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "name": "RUBICON",
            "imgUrl": "https://s-media-cache-ak0.pinimg.com/736x/0c/a5/90/0ca590b8330c80257c36ca137486244c.jpg",
            "price": 40000,
        })
    })

    // Listen and Server in 0.0.0.0:8080
    r.Run(":8080")
}