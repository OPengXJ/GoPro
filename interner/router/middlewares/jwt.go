package middlewares

import (
	"net/http"
	"github.com/OPengXJ/GoPro/configs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)
type UserClaims struct{
	UserId uint
	UserName string
	jwt.RegisteredClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("x-token")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请登陆",
			})
			ctx.Abort()
			return
		}
		token,err:=jwt.ParseWithClaims(tokenString,&UserClaims{},func(t *jwt.Token) (interface{}, error) {
			jwtConfig:=configs.Get().JwtPass
			return jwtConfig,nil
		})
		if err!=nil{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"msg":"抱歉,请重新登陆后再试",
			})
		}
		if claims,ok:=token.Claims.(*UserClaims);ok&&token.Valid{
			ctx.Set("claims",claims)
			ctx.Set("uid",claims.UserId)
			ctx.Next()
		}else{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"msg":"登陆国旗,请重新登陆",
			})
		}
	}
}
