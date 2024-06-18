package util

import (
	"acervocomics/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const secret = "superDUPERSecret"

func MintToken(user *models.User) (string, error) {
	// Cria um token JWT
	claims := jwt.MapClaims{
		"id":    user.ID,                                  // ID do usuário
		"email": user.Email,                               // Email do usuário
		"admin": user.Admin,                               // Se o usuário é admin
		"exp":   time.Now().Add(time.Minute * 120).Unix(), // Token expira em 2 horas
	}
	// Cria um novo token com as claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	// Faz o parse do token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
