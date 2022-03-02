package main

import (
	"log"
	"net/http"

	"github.com/anime454/google-two-factor-authen/handler"
	"github.com/anime454/google-two-factor-authen/repository"
	"github.com/anime454/google-two-factor-authen/service"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
)

func main() {
	mc := memcache.New("172.17.0.3:11211")
	qrRepo := repository.NewQrCodeRepositoryDB(mc)
	qrService := service.NewQrCodeService(qrRepo)
	qrHandler := handler.NewQrCodeHandler(qrService)

	tfRepo := repository.NewTwoFactorRepositoryDB(mc)
	tfService := service.NewTwoFactorService(tfRepo)
	tfHandler := handler.NewTwoFactorHandler(tfService)

	server := gin.Default()
	server.GET("/getQrCode", qrHandler.Generate())
	server.POST("/validateToken", tfHandler.ValidateToken())

	srv := &http.Server{
		Addr:    ":" + "10002",
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
