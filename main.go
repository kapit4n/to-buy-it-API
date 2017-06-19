package main

import (
    "github.com/gin-gonic/gin"
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
    Description  int `gorm:"not null" form:"description" json:"description"`
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
        v1.GET("/todobuys/:id", GetTodoBuy)
        v1.PUT("/todobuys/:id", UpdateTodoBuy)
        v1.DELETE("/todobuys/:id", DeleteTodoBuy)
    }

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

func GetTodoBuy(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    id := c.Params.ByName("id")
    var todoBuy TodoBuys
    // SELECT * FROM todoBuys WHERE id = 1;
    db.First(&todoBuy, id)

    if todoBuy.Id != 0 {
        // Display JSON result
        c.JSON(200, todoBuy)
    } else {
        // Display JSON error
        c.JSON(404, gin.H{"error": "TodoBuy not found"})
    }

    // curl -i http://localhost:8080/api/v1/todoBuys/1
}

func UpdateTodoBuy(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    // Get id todoBuy
    id := c.Params.ByName("id")
    var todoBuy TodoBuys
    // SELECT * FROM todoBuys WHERE id = 1;
    db.First(&todoBuy, id)

    if todoBuy.Name != "" && todoBuy.ImageUrl != "" {

        if todoBuy.Id != 0 {
            var newTodoBuy TodoBuys
            c.Bind(&newTodoBuy)

            result := TodoBuys{
                Id:        todoBuy.Id,
                Name: newTodoBuy.Name,
                ImageUrl:  newTodoBuy.ImageUrl,
                Price:  newTodoBuy.Price,
                Description:  newTodoBuy.Description,
            }

            // UPDATE todoBuys SET firstname='newTodoBuy.Firstname', lastname='newTodoBuy.Lastname' WHERE id = todoBuy.Id;
            db.Save(&result)
            // Display modified data in JSON message "success"
            c.JSON(200, gin.H{"success": result})
        } else {
            // Display JSON error
            c.JSON(404, gin.H{"error": "TodoBuy not found"})
        }

    } else {
        // Display JSON error
        c.JSON(422, gin.H{"error": "Fields are empty"})
    }

    // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/todoBuys/1
}

func DeleteTodoBuy(c *gin.Context) {
    // Connection to the database
    db := InitDb()
    // Close connection database
    defer db.Close()

    // Get id todoBuy
    id := c.Params.ByName("id")
    var todoBuy TodoBuys
    // SELECT * FROM todoBuys WHERE id = 1;
    db.First(&todoBuy, id)

    if todoBuy.Id != 0 {
        // DELETE FROM todoBuys WHERE id = todoBuy.Id
        db.Delete(&todoBuy)
        // Display JSON result
        c.JSON(200, gin.H{"success": "TodoBuy #" + id + " deleted"})
    } else {
        // Display JSON error
        c.JSON(404, gin.H{"error": "TodoBuy not found"})
    }

    // curl -i -X DELETE http://localhost:8080/api/v1/todoBuys/1
}