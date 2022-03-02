package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/anime454/google-two-factor-authen/repository"
	"github.com/dgryski/dgoogauth"
)

type twoFactorService struct {
	twoFactorRepo repository.TwoFactorRepository
}

func NewTwoFactorService(twoFactorRepo repository.TwoFactorRepository) twoFactorService {
	return twoFactorService{twoFactorRepo: twoFactorRepo}
}

func (twoFactorSv twoFactorService) ValidateToken(username string, token int) error {
	s, err := twoFactorSv.twoFactorRepo.Get(username)
	if err != nil {
		fmt.Println("Get token from username error ", err)
		return err
	}

	fmt.Println("DEBUG: s is ", s)
	secret := s.Value
	fmt.Println("DEBUG: secret is ", secret)
	otpConfig := &dgoogauth.OTPConfig{
		Secret:      strings.TrimSpace(secret),
		WindowSize:  3,
		HotpCounter: 0,
	}

	strToken := strconv.Itoa(token)
	fmt.Println("DEBUG: token is ", strToken)
	trimmedToken := strings.TrimSpace(strToken)

	// Validate token
	ok, err := otpConfig.Authenticate(trimmedToken)
	if err != nil {
		fmt.Println("Validate token error ", err)
		return err
	}

	if !ok {
		fmt.Println("Validate token return false", ok)
		return errors.New("InvalidToken")
	}

	return nil
}
