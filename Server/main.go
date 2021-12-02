package main

import (
	"Server/Models/ViewModel"
	"Server/Routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {

	// 新建一个没有任何默认中间件的路由
	r := Routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
	//r := gin.Default()
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}

func DoERROR(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			url := c.Request.URL
			method := c.Request.Method
			log.Printf("| url [%s] | method | [%s] | error [%s] |", url, method, err)
			var res ViewModel.ResultModel
			res.Code = 0
			res.Msg = fmt.Sprintln(err)
			res.Data = ""
			res.Count = 0
			c.JSON(http.StatusOK, res)
			c.Abort()

		}
	}()
	panic("AI YA" + time.Now().GoString())
}
