package admin

import "github.com/gin-gonic/gin"

func(h *Handle)Login()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.String(200,"👌👌")
	}
}