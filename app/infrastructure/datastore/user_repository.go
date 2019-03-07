package datastore

import (
	"goraphql/app/domain/model"
	"goraphql/app/domain/repository"
	"goraphql/app/domain/service"
)

type UserRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r UserRepository) FindOne(id int) (*model.User, error) {
	db, err := service.NewDbService().Connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var user model.User
	err = db.Find(&user, id).Error
	return &user, err
}

func (r UserRepository) List() (*[]model.User, error) {
	db, err := service.NewDbService().Connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var users []model.User
	err = db.Find(&users).Error
	return &users, err
}

func (r UserRepository) Store(newUser *model.User) error {
	db, err := service.NewDbService().Connect()
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.Create(&newUser).Error
	return err
}
