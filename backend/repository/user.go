package repository

import (
	"chathenon/entity"
	"context"

	"gorm.io/gorm"
)

type IUserRepo interface {
	CreateUser(context.Context, entity.User) error
	CheckUserExistByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (ur UserRepo) CreateUser(ctx context.Context, user entity.User) error {
	tx := ur.db.WithContext(ctx).Create(&user)

	err := tx.Error
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepo) CheckUserExistByEmail(ctx context.Context, email string) (bool, error) {
	var exist bool

	err := ur.db.WithContext(ctx).Model(&entity.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Scan(&exist).
		Error

	if err != nil {
		return exist, err
	}

	return exist, nil
}

func (ur UserRepo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var res entity.User

	err := ur.db.WithContext(ctx).Take(&res, "email = ?", email).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
