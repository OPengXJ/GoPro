package token

import (
	"net/http"
	"time"

	"github.com/OPengXJ/GoPro/configs"
	"github.com/OPengXJ/GoPro/interner/router/middlewares"
	"github.com/OPengXJ/GoPro/interner/service/admin"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(c *gin.Context, rep admin.LoginReponse) string {
	jwtConfig := []byte(configs.Get().JwtPass)
	claims := middlewares.UserClaims{
		UserId:   rep.UserId,
		UserName: rep.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errormsg": err.Error(),
		})
	}
	return tokenString
}
