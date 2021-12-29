// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	AddAccountAmount(ctx context.Context, arg AddAccountAmountParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	DeleteTransfer(ctx context.Context, id int64) error
	FromListTransfer(ctx context.Context, arg FromListTransferParams) ([]Transfer, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntriesById(ctx context.Context, arg ListEntriesByIdParams) ([]Entry, error)
	ToListTransfer(ctx context.Context, arg ToListTransferParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	UpdateAccount2(ctx context.Context, arg UpdateAccount2Params) (Account, error)
	UpdateTransfer(ctx context.Context, arg UpdateTransferParams) error
}

var _ Querier = (*Queries)(nil)
