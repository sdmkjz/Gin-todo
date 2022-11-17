package api

import (
	"Gin_todo/pkg/utils"
	"Gin_todo/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&createTask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	}
}

func ShowTask(c *gin.Context) {
	//var showTask service.ShowTaskService
	//claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	showTaskService := service.ShowTaskService{}
	res := showTaskService.Show(c.Param("id"))
	c.JSON(200, res)
}

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&listTask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	}
}
