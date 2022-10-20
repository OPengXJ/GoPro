package router

import "github.com/gin-gonic/gin"


func InitRouter()*gin.Engine{
	r:=gin.Default()
	//注册api
	setApiRouter(r)
	return r
}
