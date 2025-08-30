package models

import "errors"

var (
	ErrInvalidName  = errors.New("invalid name field")
	ErrInvalidEmail = errors.New("invalid email field")
)

type User struct {
	CommonsModel
	Name   string  `json:"name"`
	Email  string  `json:"email" gorm:"unique"`
	Orders []Order `json:"orders,omitempty" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrInvalidName
	}
	if u.Email == "" {
		return ErrInvalidEmail
	}
	return nil
}
