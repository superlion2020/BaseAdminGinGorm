package Controller

import (
	"Server/Models/ViewModel"
	"Server/Service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var res = Service.DoLogin(c)
	c.JSON(200, res)
}
func Login2(c *gin.Context) {

	var res ViewModel.ResultModel
	res.Code = 1
	res.Data = "haole"
	c.JSON(200, res)
}
func LoginPost(c *gin.Context) {
	var res = Service.DoLoginPost(c)
	c.JSON(200, res)
}
