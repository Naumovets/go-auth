package utils

import (
	"fmt"
	"time"

	"github.com/Naumovets/go-auth/internal/entities"
	"github.com/Naumovets/go-auth/internal/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user entities.User, secretKey []byte, duration time.Duration) (string, error) {
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Id:       user.Id,
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&models.UserClaims{},
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %s", err)
	}

	claims, ok := token.Claims.(*models.UserClaims)

	if !ok {
		return nil, fmt.Errorf("invalid token claims: %s", err)
	}

	return claims, nil

}
