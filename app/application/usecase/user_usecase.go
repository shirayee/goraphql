package usecase

import (
	"goraphql/app/domain/model"
	"goraphql/app/domain/repository"
)

type UserUseCase interface {
	GetUserByID(id int) (*model.User, error)
	GetUsers() (*[]model.User, error)
	CreateUser(user *model.User) (*model.User, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u userUseCase) GetUserByID(id int) (*model.User, error) {
	user, err := u.UserRepository.FindOne(id)
	if err != nil {
		user = nil
	}
	return user, err
}

func (u userUseCase) GetUsers() (*[]model.User, error) {
	users, err := u.UserRepository.List()
	if err != nil {
		users = nil
	}
	return users, err
}

func (u userUseCase) CreateUser(user *model.User) (*model.User, error) {
	err := u.UserRepository.Store(user)
	if err != nil {
		user = nil
	}
	return user, err
}
