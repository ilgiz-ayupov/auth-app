package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/ilgiz-ayupov/auth-app"
)

const salt = "qfcv123leivn94"

type AuthService struct {
}

func InitAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(user auth.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	fmt.Println(user)
	return 0, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
