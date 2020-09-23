package auth

import (
	"errors"
	"time"

	"github.com/kenshiro41/go_app/gql/models"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

var jwtSecret []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenToken(id int, userName string) (*models.Token, error) {
	iat := time.Now().Unix()
	exp := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"userName": userName,
		"iat":      iat,
		"exp":      exp,
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return &models.Token{Token: tokenString, Iat: int(iat), Exp: int(exp)}, nil
}

func DecodeUser(t string) (*UserInfo, error) {
	if t == "" {
		return nil, errors.New("token not found")
	}

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("unexpected signing method: in token.Header[alg]")
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	decodeID := claims["id"].(float64)
	id := int(decodeID)
	userName := claims["userName"].(string)

	return &UserInfo{ID: id, UserName: userName}, nil
}
