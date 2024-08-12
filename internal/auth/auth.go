package auth

import (
	"url/internal/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

var mySigningKey = []byte("secret")

func CreateNewToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	})

	return token.SignedString(mySigningKey)
}

func ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.UnexpectedSignMethod
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.InvalidToken
	}

	return claims, nil
}
