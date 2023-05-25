package jwtx

import (
	"github.com/golang-jwt/jwt/v5"
)

type CusClaim struct {
	Exp int64  `json:"exp"`
	Iat int64  `json:"iat"`
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

func NewCusClaim(iat, seconds int64, uid string) *CusClaim {
	return &CusClaim{
		Exp: iat + seconds,
		Iat: iat,
		Uid: uid,
	}
}

func GenerateToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewCusClaim(iat, seconds, uid))
	return token.SignedString([]byte(secretKey))

}

func ParseToken(secretKey string, token string) (string, error) {
	claim := &CusClaim{}
	_, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Uid, err
}
