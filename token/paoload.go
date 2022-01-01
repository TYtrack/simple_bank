/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-31 00:45:27
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-31 11:19:07
 * @FilePath: /bank_project/token/paoload.go
 */
package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken error = errors.New("token expired")
var ErrInvalidToken error = errors.New("token invalid")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := Payload{
		ID:        ID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return &payload, nil
}

func (payload *Payload) Valid() (err error) {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
