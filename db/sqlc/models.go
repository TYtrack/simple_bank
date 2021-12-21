// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Account struct {
	ID        int64        `json:"id"`
	Owner     string       `json:"owner"`
	Balance   int64        `json:"balance"`
	Currency  string       `json:"currency"`
	CreatedAt sql.NullTime `json:"createdAt"`
}

type Entry struct {
	ID        int64        `json:"id"`
	AccountID int64        `json:"accountID"`
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"createdAt"`
}

type Transfer struct {
	ID            int64        `json:"id"`
	FromAccountID int64        `json:"fromAccountID"`
	ToAccountID   int64        `json:"toAccountID"`
	Amount        int64        `json:"amount"`
	CreatedAt     sql.NullTime `json:"createdAt"`
}
