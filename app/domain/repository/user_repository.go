package repository

import "goraphql/app/domain/model"

type UserRepository interface {
	FindOne(id int) (*model.User, error)
	List() (*[]model.User, error)
	Store(newUser *model.User) error
}
