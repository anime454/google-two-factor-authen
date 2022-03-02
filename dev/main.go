package main

import (
	"log"
	"net/http"

	handler "github.com/anime454/google-two-factor-authen/handler"
	service "github.com/anime454/google-two-factor-authen/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// userRepo := repository.NewUserRepositoryDB(db)
	qrservice := service.NewQrCodeService()
	qrhandler := handler.NewQrCodeHandler(qrservice)

	server := gin.Default()
	server.GET("/getQrCode", qrhandler.Generate())

	srv := &http.Server{
		Addr:    ":" + "10002",
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
