package main

import (
	"fmt"
	"goginapi/model"
	"goginapi/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	fmt.Println("mm")

	// schedule.Task()
	err := godotenv.Load()

	if err != nil {
		fmt.Println("failed to find .env file")
	}

	dsn := "root:mysql@tcp(127.0.0.1:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to find .env file")
	}

	user := model.User{Name: "Jinzhu", Age: 18}

	db.AutoMigrate(&user, model.Book{})

	db.Create(&user)

	// declare router
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"*"},
	// 	AllowHeaders:     []string{"*"},
	// 	AllowCredentials: true,
	// }))

	router.Use(cors.Default())

	// api route call
	apiRoute := (new(routes.ApiRoutes))
	apiRoute.GetApiRoutes(router)

	// server run
	router.Run(":8000")
}
