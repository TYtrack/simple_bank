/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-28 14:34:10
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-28 16:21:01
 * @FilePath: /bank_project/api/transfer.go
 */
package api

import (
	db "bank_project/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1" `
	ToAccountID   int64 `json:"to_account_id" binding:"required,min=1" `
	Amount        int64 `json:"amount" binding:"required,gt=0" `
	// 使用currency自己注册的验证器
	Currency string `json:"currency" binding:"required,currency"`
	//Currency string `json:"currency" binding:"required,oneof=EUR USD CNY"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req TransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validTransfer(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validTransfer(ctx, req.ToAccountID, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		From_account_id: req.FromAccountID,
		To_account_id:   req.ToAccountID,
		Amount:          req.Amount,
	}

	account, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) validTransfer(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccountForUpdate(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account %v didn't match the currency[%v]", accountID, currency)
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return false
	}

	return true
}
