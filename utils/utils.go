package utils

import (
	"go-start-project/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateJWT(payload model.JWTPayload) (string, error) {
    claims := jwt.MapClaims{
        "sub": payload.Sub,
		"email": payload.Email,
		"is_verified": payload.IsVerified,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // срок действия 24 часа
		"iat": time.Now().Unix(), // время создания токена
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}