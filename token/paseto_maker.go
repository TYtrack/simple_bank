/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-31 15:19:13
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-01 16:47:41
 * @FilePath: /bank_project/token/paseto_maker.go
 */

package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto *paseto.V2
	//只在本地使用令牌API，所以用对称加密
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) < chacha20poly1305.KeySize {
		return nil, errors.New(fmt.Sprintf("the length of symmetricKey is not enough%v", len(symmetricKey)))
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (pasetoMaker *PasetoMaker) CreateToken(username string, duration time.Duration) (token string, err error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return pasetoMaker.paseto.Encrypt(pasetoMaker.symmetricKey, payload, nil)
}

func (pasetoMaker *PasetoMaker) VerifyToken(token string) (payload *Payload, err error) {
	payload = &Payload{}

	err = pasetoMaker.paseto.Decrypt(token, pasetoMaker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, ErrExpiredToken
	}
	return payload, nil
}
