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
	Login(ctx context.Context, user entity.User) (string, error)
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

func (uuc UserUsecase) Login(ctx context.Context, user entity.User) (string, error) {
	res, err := uuc.ur.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	// compare password
	err = util.CompareHashPassword([]byte(res.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// generate jwt token
	jwtToken, err := util.CreateToken(int(res.ID))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
