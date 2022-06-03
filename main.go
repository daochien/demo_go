package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Runner struct {
	ID    int64  `json:"id" gorm:"column:id;"`
	Uuid  int64  `json:"uuid" gorm:"column:uuid;"`
	Phone string `json:"phone" gorm:"column:phone;"`
	Name  string `json:"name" gorm:"column:name;"`
}

func (Runner) TableName() string {
	return "runners"
}

func main() {
	dsn := "host=localhost user=postgres password= dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println(db, err)

	db = db.Debug()

	//newRunner := Runner{Name: "daochien1", Uuid: 33021, Phone: "09891575791"}
	//
	//db.Create(&newRunner)
	//
	//log.Println(newRunner)

	//var runners []Runner
	//
	//db.Find(&runners)
	//
	//log.Println(runners)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("runners", func(c *gin.Context) {
		var newRunner Runner
		if err := c.ShouldBind(&newRunner); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newRunner).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": newRunner})
	})

	r.GET("runners", func(c *gin.Context) {
		var runners []Runner

		if err := db.Find(&runners).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": runners})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
