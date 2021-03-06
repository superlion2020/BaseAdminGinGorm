package Service

import (
	"Server/DB"
	"Server/Models/DbModel"
	"Server/Models/ViewModel"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type LoginRequest struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"name" json:"user" uri:"name" xml:"user" binding:"required"`
	Pssword string `form:"pwd" json:"password" uri:"pwd" xml:"password" binding:"required"`
}

func DoLogin(c *gin.Context) (res ViewModel.ResultModel) {
	var user DbModel.User
	for i := 10; i < 20; i++ {
		user.ID = uint(i)
		user.Name = "name" + fmt.Sprintln(i)
		DB.MysqlDB.Create(&user)
	}

	var db = DB.MysqlDB.Table("User").Where("ID>0")

	var count int64
	db.Count(&count)
	fmt.Println(count)
	var list []DbModel.User
	db.Find(&list)
	fmt.Println(list)

	DB.MysqlDB.First(&user)
	fmt.Println(user)
	//panic("a yi ya")
	var m LoginRequest
	// 将request的body中的数据，绑定到模型
	if err := c.ShouldBind(&m); err != nil {
		// 返回错误信息
		res.Code = 0
		res.Msg = "参数错误" + err.Error()
		res.Data = ""
		res.Count = 0
		return res
	}
	// 判断用户名密码是否正确
	if m.User != "root" || m.Pssword != "admin" {
		res.Code = 1
		res.Msg = "ok"
		res.Data = "chenggong+admin"
		res.Count = 0

	} else {
		res.Code = 0
		res.Msg = "登录失败"
		res.Data = ""
		res.Count = 0
	}
	return res

}

func DoLoginPost(c *gin.Context) (res ViewModel.ResultModel) {

	var m LoginRequest
	// 将request的body中的数据，绑定到模型
	if err := c.ShouldBind(&m); err != nil {
		// 返回错误信息
		res.Code = 0
		res.Msg = "参数错误" + err.Error()
		res.Data = ""
		res.Count = 0
		return res
	}
	// 判断用户名密码是否正确
	if m.User != "root" || m.Pssword != "admin" {
		res.Code = 1
		res.Msg = "ok"
		res.Data = "chenggong+admin"
		res.Count = 0

	} else {
		res.Code = 0
		res.Msg = "登录失败"
		res.Data = ""
		res.Count = 0
	}
	return res

}
