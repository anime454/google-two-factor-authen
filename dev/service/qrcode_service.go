package service

import (
	"encoding/base32"
	"fmt"
	"math/rand"

	"github.com/anime454/google-two-factor-authen/repository"
	"github.com/skip2/go-qrcode"
)

type qrCodeService struct {
	qrCodeRepo repository.QrCodeRepository
}

func NewQrCodeService(qrCodeRepo repository.QrCodeRepository) qrCodeService {
	return qrCodeService{qrCodeRepo: qrCodeRepo}
}

func (qr qrCodeService) Generate(username string) ([]byte, error) {
	var png []byte
	randomStr := randSeq(6)
	fmt.Println(randomStr)
	secret := base32.StdEncoding.EncodeToString([]byte(randomStr))
	authLink := "otpauth://totp/" + username + "?secret=" + secret + "&issuer=BotSystem&algorithm=SHA1&digits=6&period=30"
	png, err := qrcode.Encode(authLink, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("QrCode generate error", err)
		return nil, err
	}
	err = qr.qrCodeRepo.Set(username, []byte(secret))
	fmt.Println("DEBUG: username is ", username)
	fmt.Println("DEBUG: []bytes is ", secret)
	if err != nil {
		fmt.Println("Set secret error", err)
		return nil, err
	}
	return png, nil
}

func randSeq(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
