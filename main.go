package main

import (
	"campaignapi/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Bismillah123_@tcp(127.0.0.1:3306)/dbcampaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good...")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Daniel Ortega"
	userInput.Email = "daniel@ortega.com"
	userInput.Occupation = "Programmer"
	userInput.Password = "qwerty"

	userService.RegisterUser(userInput)

	// router := gin.Default()
	// router.GET("/", handler)
	// router.Run()
}

// func handler(c *gin.Context) {
// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
