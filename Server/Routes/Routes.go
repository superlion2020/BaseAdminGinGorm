package Routes

import (
	"Server/Controller"
	"Server/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.NoMethod(Middleware.HandleNotFound)
	r.NoRoute(Middleware.HandleNotRoute)
	r.Use(Middleware.CatchError)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/login", Controller.Login)
	r.POST("/loginpost", Controller.LoginPost)
	return r
}
