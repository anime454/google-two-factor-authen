package handler

import (
	"fmt"

	service "github.com/anime454/google-two-factor-authen/service"
	"github.com/gin-gonic/gin"
)

type qrCodeHandler struct {
	qrSv service.QrCodeService
}

func NewQrCodeHandler(qr service.QrCodeService) qrCodeHandler {
	return qrCodeHandler{qrSv: qr}
}

func (qrhl qrCodeHandler) Generate() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		png, err := qrhl.qrSv.Generate()
		if err != nil {
			fmt.Println("QrCode generate handler error ", err)
			c.JSON(500, gin.H{
				"code":    50000,
				"message": "Internal Server Error",
				"data":    nil,
			})
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Content-Type", "image/png")
			c.Header("Content-Disposition", "attachment; filename=response.png")
			c.Writer.Write(png)
		}
	}
	return fn
}
