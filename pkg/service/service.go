package service

import (
	"math/rand"
	"time"
)

type PasswordService interface {
	GeneratePassword(length int, includeUppercase, includeLowercase, includeNumbers, includeSpecial bool) (string, error)
}

type passwordService struct{}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (s *passwordService) GeneratePassword(length int, includeUppercase, includeLowercase, includeNumbers, includeSpecial bool) (string, error) {
    var charset string
    if includeUppercase {
        charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    }
    if includeLowercase {
        charset += "abcdefghijklmnopqrstuvwxyz"
    }
    if includeNumbers {
        charset += "0123456789"
    }
    if includeSpecial {
        charset += "!@#$%^&*()-_+=<>?"
    }
    if charset == "" {
        charset = "abcdefghijklmnopqrstuvwxyz"
    }

    rand.Seed(time.Now().UnixNano())
    password := make([]byte, length)
    for i := range password {
        password[i] = charset[rand.Intn(len(charset))]
    }
    return string(password), nil
}