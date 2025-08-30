package constants

import "errors"

const (
	BadRequest = "Bad request Error"
)

var (
	RequiredNameField        = errors.New("name is required")
	RequiredEmailField       = errors.New("email is required")
	RequiredPhoneNumberField = errors.New("phone number is required")
)
