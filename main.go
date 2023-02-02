package main

import (
	"campaignapi/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	router.GET("/", handler)
	router.Run()
}

func handler(c *gin.Context) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Bismillah123_@tcp(127.0.0.1:3306)/dbcampaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good...")

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
