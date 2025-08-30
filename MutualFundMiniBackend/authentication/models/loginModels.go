package models

import (
	"authentication/constants"
	"time"
)

type LoginRequestModel struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	PhoneNumber uint64    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

type LoginResponseModel struct {
	Success string `json:"message"`
}

func (response *LoginRequestModel) Validate() error {
	if response.Name == "" {
		return constants.RequiredNameField
	}
	if response.Email == "" {
		return constants.RequiredEmailField
	}
	if response.PhoneNumber == 0 {
		return constants.RequiredPhoneNumberField
	}
	return nil
}
