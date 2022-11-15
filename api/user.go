package api

import (
	"Gin_todo/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBindJSON(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		fmt.Println(err)
		c.JSON(400, err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBindJSON(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		fmt.Println(err)
		c.JSON(400, err)
	}
}
