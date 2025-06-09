package in

import "github.com/golang-jwt/jwt/v5"

type ValidateTokenPort interface {
	Execute(token *string)  (*jwt.Token, error)
}
