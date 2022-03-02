package qrcode

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

type qrCodeService struct{}

func NewQrCodeService() qrCodeService {
	return qrCodeService{}
}

func (qr qrCodeService) Generate() ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		fmt.Println("QrCode generate error", err)
		return nil, err
	}
	return png, nil
}
