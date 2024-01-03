package repository

import "im/services/user/domain/entity"

type UserRepository interface {
	GetUserInfoByEmail(email string) (*entity.User, error)
	GetUserInfoByUid(id string) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	InsertUser(user *entity.User) (*entity.User, error)
}
