package router

import (
	"github.com/OPengXJ/GoPro/interner/api/admin"
	"github.com/gin-gonic/gin"
)

func setApiRouter(r *gin.Engine){
	//admin
	adminHandler:=admin.New()
	login:=r.Group("/api")
	{
		login.GET("/login",adminHandler.Login())
	}
}