package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iliaposmac/todo-app"
	"github.com/iliaposmac/todo-app/pkg/repository"
)

const (
	salt      = "asdas123asdasdasdasd"
	tokenTTL  = 12 * time.Hour
	signinKey = "asdasdasdasdasd"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUser(username, generateHashPassword(password))

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

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
