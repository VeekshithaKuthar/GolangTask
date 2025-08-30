package repositories

import (
	"paymenst/models"

	"gorm.io/gorm"
)

type PaymentOrderDb struct {
	DB *gorm.DB
}

func NewPaymentDB(db *gorm.DB) IPaymentDB {
	return &PaymentOrderDb{db}
}

type IPaymentDB interface {
	CreatePayment(payment *models.Payments) (*models.Payments, error)
	UpdatePaymentStatus(order_Id string, paymentstatus string) error
}

func (udb *PaymentOrderDb) UpdatePaymentStatus(order_Id string, paymentstatus string) error {
	tx := udb.DB.Model(&models.Payments{}).
		Where("order_id= ?", order_Id).Update("payment_status", paymentstatus)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (udb *PaymentOrderDb) CreatePayment(payment *models.Payments) (*models.Payments, error) {
	tx := udb.DB.Create(payment)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return payment, nil
}
