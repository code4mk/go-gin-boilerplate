package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gogin api",
		"api":     "golang",
		"version": "v1.0.2",
		"author":  "kamal 2121",
	})
	return
}
