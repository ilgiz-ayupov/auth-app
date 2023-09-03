package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/repository"
)

const (
	salt       = "qfcv123leivn94"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type AuthService struct {
	repo *repository.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	auth.UserTokenClaims
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user auth.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateJWTToken(login string, password string) (string, error) {
	password = generatePasswordHash(password)
	id, err := s.repo.AuthentificationUser(login, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		auth.UserTokenClaims{
			Login:  login,
			UserId: id,
		},
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseJWTToken(token string) (auth.UserTokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return auth.UserTokenClaims{}, err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok {
		return auth.UserTokenClaims{}, errors.New("token claims are not of type *tokenClaims")
	}

	return auth.UserTokenClaims{
		Login:  claims.Login,
		UserId: claims.UserId,
	}, nil
}

func (s *AuthService) AuthorizationToken(claims auth.UserTokenClaims) error {
	return s.repo.AuthorizationToken(claims)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
