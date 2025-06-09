package services

import (
	"fmt"

	"github.com/felipeemora/social-hex-api/internal/infraestructure/drivenadapters/configurations"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secretKey string
	aud       string
	iss       string
}

func NewTokenService(JWTConfig *configurations.JWTConfig) *TokenService {
	return &TokenService{
		secretKey: JWTConfig.SecretKey,
		aud:       JWTConfig.Audience,
		iss:       JWTConfig.Issuer,
	}
}

func (j *TokenService) GenerateToken(claims jwt.Claims) (string, error) {
	claims.(jwt.MapClaims)["aud"] = j.aud
	claims.(jwt.MapClaims)["iss"] = j.iss

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *TokenService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(j.aud),
		jwt.WithIssuer(j.iss),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
