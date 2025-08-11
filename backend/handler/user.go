package handler

import (
	"chathenon/dto"
	"chathenon/entity"
	"chathenon/usecase"
	"chathenon/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
}

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
