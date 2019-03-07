package model

import "time"

type User struct {
	ID        int        `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

func NewUser(name string) (*User, error) {
	// TODO - validate
	return &User{
		Name: name,
	}, nil
}
