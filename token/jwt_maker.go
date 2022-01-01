/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-31 00:53:16
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-31 12:49:54
 * @FilePath: /bank_project/token/jwt_maker.go
 */

package token

import (
	"errors"
	"time"

	// jwt "github.com/dgrijalva/jwt-go"
	jwt "github.com/golang-jwt/jwt/v4"
)

const MinSecretKey int = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < MinSecretKey {
		return nil, errors.New("the length of secretKey is not enough")
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

func (jwt_maker *JWTMaker) CreateToken(username string, duration time.Duration) (token string, err error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(jwt_maker.secretKey))
}

func (jwt_maker *JWTMaker) VerifyToken(token string) (payload *Payload, err error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwt_maker.secretKey), nil
	}
	jwt_token, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwt_token.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
