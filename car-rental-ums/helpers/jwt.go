package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type ClaimsToken struct {
	FirstName string
	LastName  string
	Email     string
	Role      string
	jwt.RegisteredClaims
}

var TokenType = map[string]time.Duration{
	"token":         time.Hour * 3,
	"refresh_token": time.Hour * 24,
}

var secretKey = []byte(GetEnv("JWT_SECRET_KEY"))

func GenerateJWTToken(firstName, lastName, email, role, tokenType string) (string, error) {
	claims := &ClaimsToken{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME"),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenType[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return signedString, nil
}

func ValidateToken(tokenString string) (*ClaimsToken, error) {
	var (
		claims = &ClaimsToken{}
		ok     bool
	)

	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	if claims, ok = jwtToken.Claims.(*ClaimsToken); !ok && !jwtToken.Valid {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	return claims, nil
}
