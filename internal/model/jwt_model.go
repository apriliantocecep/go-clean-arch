package model

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-clean-arch/internal/entity"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (wrapper JwtWrapper) GenerateToken(user entity.User) (signedToken string, err error) {
	claims := &JwtClaims{
		Id:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    wrapper.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(wrapper.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (wrapper JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(wrapper.SecretKey), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
