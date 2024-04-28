package main

import (
	"fmt"
	"ginWeb/models"
	"ginWeb/pkg/setting"
	"ginWeb/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	models.Setup()
}

func main() {
	gin.SetMode(setting.RunMode)

	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	routerInit := routers.InitRouter()
	readTimeout := setting.ReadTimeout
	WriteTimeout := setting.WriteTimeout
	MaxHeaderBytes := 1 << 20
	s := &http.Server{
		Addr:           endPoint,
		Handler:        routerInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	s.ListenAndServe()
}
