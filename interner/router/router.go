package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine{
	r:=gin.Default()
	//注册api
	pprof.Register(r)
	setApiRouter(r)
	return r
}
