package business

import (
	"authentication/models"
	"authentication/repositories"
)

type ILoginService interface {
	CreateUser(user *models.LoginRequestModel) error
}

type LoginService struct {
	repo repositories.ILoginRepository
}

func NewLoginService(repo repositories.ILoginRepository) ILoginService {
	return &LoginService{repo: repo}
}

func (service *LoginService) CreateUser(user *models.LoginRequestModel) error {
	if err := service.repo.CreateUser(user); err != nil {
		return err
	}
	return nil

}
