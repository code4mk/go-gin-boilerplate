package main

import (
	"fmt"
	"goginapi/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("failed to find .env file")
	}

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
