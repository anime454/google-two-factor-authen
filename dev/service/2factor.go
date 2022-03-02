package service

type TwoFactorService interface {
	ValidateToken(string, int) error
}
