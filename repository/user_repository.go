package repository

import (
	"todo/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// dbを引数にuserRepositoryの構造体を作成し，そのポインタを返す
// ここで，IUserRepository型を返すときはすべてのメゾットを満たす必要あり
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// ユーザ発見時，引数がさすuser情報を，DBで発見したユーザ情報に書き換え
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	// ユーザ登録時，引数がさすuser情報を，DBに登録したユーザ情報に書き換え
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
