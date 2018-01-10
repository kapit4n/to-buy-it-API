package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
)

var DB = make(map[string]string)

type TodoBuys struct {
    Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Name string `gorm:"not null" form:"name" json:"name"`
    ImageUrl string `gorm:"not null" form:"imageUrl" json:"imageUrl"`
    Price  int `gorm:"not null" form:"price" json:"price"`
    Description  string `gorm:"not null" form:"description" json:"description"`
    Done  bool `gorm:"not null" form:"done" json:"done"`
}

func main() {
    r := gin.Default()

    v1 := r.Group("api/v1")
    {
        v1.GET("/todobuys", GetTodoBuys)
        v1.POST("/todobuys", PostTodoBuys)
        v1.POST("/todobuys/:id", UpdateTodoBuy)
        v1.GET("/todobuys/:id", GetTodoBuy)
        v1.DELETE("/todobuys/:id", DeleteTodoBuy)
    }

    r.Use(Cors())

    // Listen and Server in 0.0.0.0:8080
    r.Run(":8080")
}

func Database() *gorm.DB {
   //open a db connection
    db, err := gorm.Open("mysql", "root:root@/todobuy?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
          panic("failed to connect database")
    }
    if !db.HasTable(&TodoBuys{}) {
        db.CreateTable(&TodoBuys{})
        db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&TodoBuys{})
    }

    return db
}

func PostTodoBuys(c *gin.Context) {
    db := Database()
    
    var todoBuy TodoBuys
    c.BindJSON(&todoBuy)
    db.Create(&todoBuy)
    c.JSON(200, todoBuy)
   defer db.Close()
}

func GetTodoBuys(c *gin.Context) {
    // Connection to the database
    db := Database()
    
    var todoBuys []TodoBuys
    // SELECT * FROM todoBuys
    db.Find(&todoBuys)

    // Display JSON result
    c.JSON(200, todoBuys)
    // Close connection database
    defer db.Close()

    // curl -i http://localhost:8080/api/v1/todoBuys
}

func GetTodoBuy(c *gin.Context) {
    // Connection to the database
    db := Database()
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
    fmt.Printf("%s\n", 777777777777)
    db := Database()
    
    // Get id todoBuy
    id := c.Param("id")


    var todoBuy TodoBuys
    // SELECT * FROM todoBuys WHERE id = 1;
    db.First(&todoBuy, id)
    fmt.Printf("%s\n", id)
    fmt.Printf("GET %+v\n", todoBuy)

    c.BindJSON(&todoBuy)
    todoBuy.Done = true;
    db.Save(&todoBuy)
    c.JSON(200, todoBuy)
    defer db.Close()

    // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/todoBuys/1
}

func DeleteTodoBuy(c *gin.Context) {
    // Connection to the database
    db := Database()
    
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
    // Close connection database
    defer db.Close()

    // curl -i -X DELETE http://localhost:8080/api/v1/todoBuys/1
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
        c.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
        c.Next()
    }
}
