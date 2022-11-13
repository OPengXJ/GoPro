package admin

import (
	"net/http"

	"github.com/OPengXJ/GoPro/interner/repository/redis"

	"github.com/gin-gonic/gin"
)

func (h *Handle) ATest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		redis := redis.GetRedisRepo()
		redis.Set("hello", "world", 0)
		ctx.String(http.StatusOK, redis.Get("hello").String())
	}
}
