package service

type QrCodeService interface {
	Generate(string) ([]byte, error)
}
