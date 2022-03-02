package qrcode

type QrCodeService interface {
	Generate() ([]byte, error)
}
