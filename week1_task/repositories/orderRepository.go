package repositories

import (
	"errors"
	"ordersApi/models"

	"gorm.io/gorm"
)

type OrderDb struct {
	DB   *gorm.DB
	Jobs chan Job
}
type Job struct {
	Id      uint
	OrderId uint
}
type IOrderDB interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetBy(id uint) (*models.User, error)
	GetOrderById(id uint) ([]*models.Order, error)
	ConfirmOrder(id uint) error
}

func NewOrderDB(db *gorm.DB, job chan Job) IOrderDB {
	return &OrderDb{db, job}
}

func (udb *OrderDb) CreateOrder(order *models.Order) (*models.Order, error) {
	_, err := udb.GetBy(order.UserID)
	if err != nil {
		return nil, errors.New("invalid user or user not found")
	}
	// var user models.User
	// if err := udb.DB.First(&user, order.UserID).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return nil, errors.New("user not found")
	// 	}
	// 	return nil, err // any other DB error
	// }
	tx := udb.DB.Create(order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

func (udb *OrderDb) GetBy(id uint) (*models.User, error) {
	user := new(models.User)
	tx := udb.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil

}

func (udb *OrderDb) GetOrderById(id uint) ([]*models.Order, error) {
	var orders []*models.Order
	tx := udb.DB.Where("user_id= ?", id).Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func (udb *OrderDb) ConfirmOrder(id uint) error {
	order := new(models.Order)
	tx := udb.DB.First(order, id)
	if tx.Error != nil {
		return tx.Error
	}
	job := Job{Id: id, OrderId: id}
	udb.Jobs <- job

	return nil
}

// func (udb *OrderDb) UpdateOrder(chan1 Job, order *models.Order) {
// 	time.Sleep(time.Second * 2)
// 	order.Status = "failed"
// 	if rand.Intn(2) == 0 {
// 		order.Status = "confirmed"
// 	}
// 	udb.DB.Model(order).Where("id= ?", chan1.Id).Update("status", order.Status)
// }
