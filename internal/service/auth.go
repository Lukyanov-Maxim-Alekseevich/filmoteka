package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/MrSaveliy/vk-filmoteca/internal/repository"
	"github.com/MrSaveliy/vk-filmoteca/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt      = "adsfsdglls213lkdf21k2"
	signinKey = "asdf3232kj42^$^"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)

}

func (s *AuthService) CreateRole(role models.Role) (int, error) {
	return s.repo.CreateRole(role)

}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {

	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signinKey))
}

func (c *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid string method")
		}
		return []byte(signinKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("tokem claims are not of type *tikenClaims")
	}
	return claims.UserId, err
}
