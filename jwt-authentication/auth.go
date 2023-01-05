package jwt_authentication

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type CustomClaims struct {
	UserId string
	jwt.StandardClaims
}

var SecretKey = []byte("dxU6g=fb")

func GenerateToken(userId string) string {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "snow",
			Subject:   "go-in-action",
			Audience:  userId,
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			ExpiresAt: now.Add(time.Hour * 24 * 365).Unix(),
		},
	})
	t, _ := token.SignedString(SecretKey)
	return t
}

func VerifyToken(token string) (userId string, err error) {
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := t.Claims.(*CustomClaims); ok {
		return claims.UserId, nil
	}

	return "", errors.New("invalid token")
}
