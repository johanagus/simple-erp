package utils

import (
	"fmt"
	"time"

	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/johanagus/simple-erp/config"
)

var (
	jwtSecret        = []byte(config.GetEnv("JWT_SECRET", "secret"))
	jwtRefreshSecret = []byte(config.GetEnv("JWT_REFRESH_SECRET", "secret"))
)

func GenerateAccessToken(UserID int, Email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    UserID,
		"email": Email,
		"exp": func() int64 {
			expStr := config.GetEnv("EXPIRED_JWT_TOKEN", "24") // default expired access token 24 jam
			expInt := 24
			if v, err := strconv.Atoi(expStr); err == nil {
				expInt = v
			}
			return time.Now().Add(time.Hour * time.Duration(expInt)).Unix()
		}(),
	}

	fmt.Println(config.GetEnv("JWT_SECRET", "secret"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetEnv("JWT_SECRET", "secret")))
}

func GenerateRefreshToken(UserID int) (string, error) {
	claims := jwt.MapClaims{
		"id": UserID,
		"exp": func() int64 {
			expStr := config.GetEnv("EXPIRED_JWT_REFRESH_TOKEN", "168") // default expired resfresh token 168 jam / 1 minggu
			expInt := 168
			if v, err := strconv.Atoi(expStr); err == nil {
				expInt = v
			}
			return time.Now().Add(time.Hour * time.Duration(expInt)).Unix()
		}(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetEnv("JWT_REFRESH_SECRET", "secret")))
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtRefreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, err
	}

	return token, claims, nil
}
