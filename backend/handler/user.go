package handler

import (
	"chathenon/dto"
	"chathenon/entity"
	"chathenon/usecase"
	"chathenon/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uuc usecase.IUserUseCase
}

func NewUserHandler(uuc usecase.UserUsecase) UserHandler {
	return UserHandler{uuc: uuc}
}

func (uh UserHandler) RegisterHandler(ctx *gin.Context) {
	var user dto.RegisterUserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Error(err)
		return
	}

	data := entity.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	err = uh.uuc.Register(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}

	util.ResponseMsg(ctx, true, nil, nil, http.StatusOK)
}

func (uh UserHandler) LoginHandler(ctx *gin.Context) {
	var user dto.LoginUserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Error(err)
		return
	}

	data := entity.User{
		Email:    user.Email,
		Password: user.Password,
	}

	token, err := uh.uuc.Login(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}

	respData := dto.LoginUserRes{Token: token}

	util.ResponseMsg(ctx, true, nil, respData, http.StatusOK)
}
