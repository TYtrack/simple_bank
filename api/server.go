/*
 * @Author: your name
 * @Date: 2021-12-22 12:14:06
 * @LastEditTime: 2022-01-01 20:17:11
 * @LastEditors: TYtrack
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /bank_project/api/server.go
 */

package api

import (
	db "bank_project/db/sqlc"
	"bank_project/token"
	"bank_project/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user/login", server.loginUser)
	router.POST("/user", server.createUser)

	AuthRouter := router.Group("/").Use(authMiddleWare(server.tokenMaker))
	AuthRouter.POST("/account", server.createAccount)
	AuthRouter.GET("/account/:id", server.getAccount)
	AuthRouter.GET("/accountlist", server.getAccountList)
	AuthRouter.POST("/transfers", server.createTransfer)
	AuthRouter.GET("/user/:username", server.getUser)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
