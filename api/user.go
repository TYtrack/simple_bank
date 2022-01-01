/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-29 16:57:59
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-01 17:16:46
 * @FilePath: /bank_project/api/user.go
 */

package api

import (
	db "bank_project/db/sqlc"
	"bank_project/util"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateUserRequest struct {
	UserName string `json:"username" binding:"required,alphanum" `
	Password string `json:"password" binding:"required,min=6" `
	FullName string `json:"fullname" binding:"required" `
	Email    string `json:"email" binding:"required,email" `
}

type LoginUserRequest struct {
	UserName string `json:"username" binding:"required,alphanum" `
	Password string `json:"password" binding:"required,min=6" `
}

type UserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"fullName"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"createdAt"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
}
type LoginUserResponse struct {
	User  UserResponse `json:"user" `
	Token string       `json:"token" `
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:     req.UserName,
		HashPassword: hashPassword,
		Email:        req.Email,
		FullName:     req.FullName,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := GetUserResponse(user)

	ctx.JSON(http.StatusOK, res)
}

type GetUserRequest struct {
	Username string `uri:"username" binding:"required,min=1" `
}

func (server *Server) getUser(ctx *gin.Context) {
	var req GetUserRequest
	fmt.Println(ctx.Params)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	userResponse := GetUserResponse(user)

	ctx.JSON(http.StatusOK, userResponse)
}

//封装： 在User的基础上不要hash_password
func GetUserResponse(user db.User) UserResponse {
	res := UserResponse{
		CreatedAt:         user.CreatedAt,
		Username:          user.Username,
		PasswordChangedAt: user.PasswordChangedAt,
		FullName:          user.FullName,
		Email:             user.Email,
	}
	return res
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(user.HashPassword, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	token, err := server.tokenMaker.CreateToken(req.UserName, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := LoginUserResponse{
		User:  GetUserResponse(user),
		Token: token,
	}

	ctx.JSON(http.StatusOK, rsp)

}
