package util

import "errors"

var (
	// errors
	ErrDataNotFound       = errors.New("data not found")
	ErrInvalidData        = errors.New("data is invalid")
	ErrInvalidCredentials = errors.New("credentials is invalid")
	ErrUnauthorized       = errors.New("sign in to proceed")
	ErrForbidden          = errors.New("you are not allowed to access this resource")

	// message
	DataNotFoundMessage       = "Data not found"
	InvalidDataMessage        = "Data is invalid"
	InvalidCredentialsMessage = "Credentials is invalid"
	UnauthorizedMessage       = "Sign in to proceed"
	ForbiddenMessage          = "You are not allowed to access this resource"
)