package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
)

var DB = make(map[string]string)

type TodoBuys struct {
    Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Name string `gorm:"not null" form:"name" json:"name"`
    ImageUrl string `gorm:"not null" form:"imageUrl" json:"imageUrl"`
    Price  int `gorm:"not null" form:"price" json:"price"`
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

    v1 := r.Group("api/v1")
    {
        v1.GET("/todobuys", GetTodoBuys)
        v1.POST("/todobuys", PostTodoBuys)
    }

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

func PostTodoBuys(c *gin.Context) {
    db := InitDb()
    defer db.Close()

    var todoBuy TodoBuys
    c.Bind(&todoBuy)
    fmt.Printf("%s\n", "This is the common")
    fmt.Printf("%+v\n", todoBuy)
    if todoBuy.Name != "" && todoBuy.ImageUrl != "" {
        // INSERT INTO "todoBuys" (name) VALUES (todoBuy.Name);
        db.Create(&todoBuy)
        // Display error
        c.JSON(201, gin.H{"success": todoBuy})
    } else {
        // Display error
        c.JSON(422, gin.H{"error": "Fields are empty"})
    }

    // curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"RUBICON\", \"imageUrl\": \"https://s-media-cache-ak0.pinimg.com/736x/0c/a5/90/0ca590b8330c80257c36ca137486244c.jpg\" , \"price\": \"40000\" }" http://localhost:8080/api/v1/todoBuys
}

func GetTodoBuys(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    var todoBuys []TodoBuys
    // SELECT * FROM todoBuys
    db.Find(&todoBuys)

    // Display JSON result
    c.JSON(200, todoBuys)

    // curl -i http://localhost:8080/api/v1/todoBuys
}