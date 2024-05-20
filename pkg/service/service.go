package service

import (
	"math/rand"
	"time"
)

type PasswordService interface {
	GeneratePassword(length int) string
}

type passwordService struct{}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (p *passwordService) GeneratePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+=-~`"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}