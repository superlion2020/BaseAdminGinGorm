package Controller

import (
	"Server/Service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var res = Service.DoLogin(c)
	c.JSON(200, res)
}
func LoginPost(c *gin.Context) {
	var res = Service.DoLoginPost(c)
	c.JSON(200, res)
}
