package repository

type QrCodeRepository interface {
	Set(string, []byte) error
}
