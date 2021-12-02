package Middleware

import (
	"Server/Models/ViewModel"
	"fmt"
	"log"
	"runtime/debug"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CatchError(c *gin.Context) {

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
	c.Next()

}

func CatchError2(cc *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
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
				cc.JSON(http.StatusOK, res)
				cc.Abort()
			}
		}()
		c.Next()
	}
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			c.JSON(200, gin.H{
				"code": "4444",
				"msg":  "服务器内部错误",
			})
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

func HandleNotFound(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "4444",
		"msg":  "路径不对啊",
	})
	//加载完 defer recover，继续后续接口调用
	c.Abort()
}
func HandleNotRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "4444",
		"msg":  "查看POST和GET是否正确",
	})
	//加载完 defer recover，继续后续接口调用
	c.Abort()
}
func Auth(c *gin.Context) {
	token, b := c.GetQuery("token")
	if !b {
		c.JSON(200, gin.H{
			"code": "1",
			"msg":  "请登录1",
		})
		c.Abort()
	} else if fmt.Sprint(token) != "12345" {
		fmt.Println(fmt.Sprintln(token))
		c.JSON(200, gin.H{
			"code": "1",
			"msg":  "请登录2",
		})
		c.Abort()
	}

	c.Next()

	//加载完 defer recover，继续后续接口调用

}
