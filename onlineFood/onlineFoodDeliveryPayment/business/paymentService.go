package business

import (
	"paymenst/models"
	"paymenst/repositories"
)

type IPaymentService interface {
	CreatePayment(payment *models.Payments) (*models.Payments, error)
}
type PaymentService struct {
	repo repositories.IPaymentDB
}

func NewPaymentService(repo repositories.IPaymentDB) IPaymentService {
	return &PaymentService{repo: repo}
}
func (service *PaymentService) CreatePayment(payment *models.Payments) (*models.Payments, error) {
	if int(payment.PaymentAmount)%2 == 0 {
		payment.PaymentStatus = "Success"
	} else {
		payment.PaymentStatus = "Failure"
	}
	payments, err := service.repo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}
	return payments, nil
}
