package main

import (
	"net/http"

	"github.com/OPengXJ/GoPro/interner/router"
	"github.com/gin-gonic/gin"
	"github.com/OPengXJ/GoPro/pkg/log"
)

func main() {
	//config:=configs.Get()
	gin.SetMode("test")
	router:=router.InitRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
	}
	log.Debug("test",log.String("hello","log work success"))
	log.Error("test",log.String("hello","log work"))
	server.ListenAndServe()
}
