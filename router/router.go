package router

import (
	"go-world/handler"
	"go-world/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())

	r.GET("ping", middleware.AuthMiddleware(), handler.PingHandler)

	r.GET("/user/:id", handler.UserHandler)

	r.GET("/search", handler.SearchHandler)

	r.POST("/login", handler.LoginHandler)

	r.GET("/protected", middleware.AuthMiddleware(), handler.ProtectedHandler)

	return r
}
