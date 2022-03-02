package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anime454/google-two-factor-authen/service"
	"github.com/gin-gonic/gin"
)

type twoFactorCodeHandler struct {
	twoFactorHl service.TwoFactorService
}

func NewTwoFactorHandler(twhl service.TwoFactorService) twoFactorCodeHandler {
	return twoFactorCodeHandler{twoFactorHl: twhl}
}

func (t twoFactorCodeHandler) ValidateToken() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		username := "mockupUsername@gmail.com"
		token := c.PostForm("token")
		trimTokem := strings.TrimSpace(token)
		fmt.Println("DEBUG: Token is ", trimTokem)
		intToken, err := strconv.Atoi(trimTokem)
		if err != nil {
			fmt.Println("Convert string to int error ", err)
			c.JSON(500, gin.H{"code": 50000, "message": "Internal Server Error", "data": nil})
			return
		}

		err = t.twoFactorHl.ValidateToken(username, intToken)
		if err != nil {
			fmt.Println("DEBUG: err is ", err.Error())
			if err.Error() == "InvalidToken" {
				fmt.Println("Invalid token", err)
				c.JSON(200, gin.H{"code": 20001, "message": "Invalid token", "data": nil})
				return
			} else {
				fmt.Println("QrCode generate handler error ", err)
				c.JSON(500, gin.H{"code": 50000, "message": "Internal Server Error", "data": nil})
				return
			}
		}
		c.JSON(200, gin.H{"code": 20000, "message": "Success", "data": nil})
	}
	return fn
}
