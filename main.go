package main

import (
	"net/http"
	"github.com/OPengXJ/GoPro/interner/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//config:=configs.Get()
	gin.SetMode("test")
	router:=router.InitRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
	}
	server.ListenAndServe()
}
