package usecase

import (
	"chathenon/constant"
	"chathenon/entity"
	"chathenon/repository"
	"chathenon/util"
	"context"
)

type IUserUseCase interface {
	Register(ctx context.Context, user entity.User) error
}

type UserUsecase struct {
	ur repository.IUserRepo
}

func NewUserUseCase(ur repository.IUserRepo) UserUsecase {
	return UserUsecase{
		ur: ur,
	}
}

func (uuc UserUsecase) Register(ctx context.Context, user entity.User) error {
	isExist, err := uuc.ur.CheckUserExistByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if isExist {
		return constant.ErrEmailAlreadyUsed
	}

	// TODO: hash password
	hashedPw, err := util.GenerateBcrypt(user.Password)
	if err != nil {
		return err
	}

	user.Password = *hashedPw

	err = uuc.ur.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
