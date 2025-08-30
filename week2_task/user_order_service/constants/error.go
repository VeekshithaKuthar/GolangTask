package constants

import "errors"

const (
	BadRequest      = "Invalid request body"
	OrderIdRequired = "order_id is required"
)

var (
	ErrInvalidCustomerName = errors.New("customer name cannot be empty")
	ErrInvalidAddress      = errors.New("address cannot be empty")
	ErrInvalidItem         = errors.New("item must be provided")
	ErrInvalidSize         = errors.New("size must be one of: small, medium, large")
)
