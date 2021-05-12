package main

import (
	"fmt"
	"goginapi/controllers"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("failed to find .env file")
	}

	fmt.Println("ok")

	router := gin.Default()

	r := router.Group("/v1")

	home := new(controllers.UserController)

	r.GET("/", home.Retrieve)

	r.GET("/user/:id", func(c *gin.Context) {
		user := c.Params.ByName("id")
		u := c.Param("id")
		b := c.Request.Host
		r := c.Query("name")

		headerGetToken := c.GetHeader("token")

		c.JSON(200, gin.H{
			"uid":         user,
			"u":           u,
			"fullpath":    c.FullPath(),
			"a":           b,
			"query-name":  r,
			"header-name": headerGetToken,
		})
	})

	r.GET("env", func(c *gin.Context) {

		s3Bucket := os.Getenv("SECRET_KEY")

		c.JSON(200, gin.H{
			"s3-bucket": s3Bucket,
		})
	})

	router.Run(":8000")
}
