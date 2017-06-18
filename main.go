package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

var DB = make(map[string]string)

type TodoBuys struct {
    Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Name string `gorm:"not null" form:"name" json:"name"`
    imageUrl string `gorm:"not null" form:"imageUrl" json:"imageUrl"`
    price  int `gorm:"not null" form:"price" json:"price"`
}

func InitDb() *gorm.DB {
    // Openning file
    db, err := gorm.Open("sqlite3", "./data.db")
    db.LogMode(true)
    // Error
    if err != nil {
        panic(err)
    }
    // Creating the table
    if !db.HasTable(&TodoBuys{}) {
        db.CreateTable(&TodoBuys{})
        db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&TodoBuys{})
    }

    return db
}

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