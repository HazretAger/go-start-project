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

func GetAccessAndRefreshTokens(payload model.JWTPayload) (model.Tokens, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": payload.Sub,
		"email": payload.Email,
		"is_verified": payload.IsVerified,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // срок действия 24 часа
		"iat": time.Now().Unix(), // время создания токена
    }).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return model.Tokens{}, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": payload.Sub,
		"email": payload.Email,
		"is_verified": payload.IsVerified,
        "exp":   time.Now().Add((time.Hour * 24) * 30).Unix(), // срок действия 30 дней
		"iat": time.Now().Unix(), // время создания токена
    }).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return model.Tokens{}, err
	}
	
    return model.Tokens{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}