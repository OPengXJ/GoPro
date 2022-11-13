package router

import (
	"github.com/OPengXJ/GoPro/interner/api/admin"
	"github.com/OPengXJ/GoPro/interner/router/middlewares"
	"github.com/gin-gonic/gin"
)

func setApiRouter(r *gin.Engine){
	//admin
	adminHandler:=admin.New()
	r.POST("/login",adminHandler.Login())
	r.POST("/create",adminHandler.Create())
	login:=r.Group("/api")
	login.Use(middlewares.JWTAuth())
	{
		login.GET("/test",adminHandler.ATest())
	}

}