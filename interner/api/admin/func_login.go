package admin

import (
	"net/http"

	"github.com/OPengXJ/GoPro/interner/pkg/token"
	"github.com/OPengXJ/GoPro/interner/repository/mysql"
	"github.com/OPengXJ/GoPro/interner/service/admin"
	"github.com/gin-gonic/gin"
)


func(h *Handle)Login()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		req:=&admin.LoginRequest{}
		err:=ctx.ShouldBind(&req)
		if err!=nil{
			ctx.String(http.StatusBadRequest,"绑定信息出错：%v",err)
		}
		rep:=mysql.GetMysqlRepo()
		service:=admin.New(*rep)
		res,err:=service.Login(req)
		if err!=nil{
			ctx.String(http.StatusBadRequest,"登陆名或者密码不正确：%v",err)
		}
		tokenString:=token.CreateToken(ctx,*res)
		ctx.SetCookie("token",tokenString,3600*24,"/","",false,true)
		ctx.JSON(http.StatusOK,gin.H{
			"user":res,
			"token":tokenString,
		})
	}
}