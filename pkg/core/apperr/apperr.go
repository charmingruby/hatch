package apperr

import (
	"errors"
)

type ErrorType string

const (
	TypeNotFound   = "NOT_FOUND"
	TypeInternal   = "INTERNAL"
	TypeValidation = "VALIDATION"
	TypeConflict   = "CONFLICT"
)

type Error struct {
	Details any       `json:"details,omitempty"`
	Err     error     `json:"-"`
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    string    `json:"code"`
}

func New(t ErrorType, message string, err error) *Error {
	return &Error{
		Type:    t,
		Message: message,
		Err:     err,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

func NotFound(message string) *Error {
	return New(TypeNotFound, message, nil)
}

func IsNotFound(err error) bool {
	return IsType(err, TypeNotFound)
}

func Conflict(message string) *Error {
	return New(TypeConflict, message, nil)
}

func IsConflict(err error) bool {
	return IsType(err, TypeConflict)
}

func Validation(message string) *Error {
	return New(TypeValidation, message, nil)
}

func IsValidation(err error) bool {
	return IsType(err, TypeValidation)
}

func Internal(message string, err error) *Error {
	return New(TypeInternal, message, nil)
}

func IsInternal(err error) bool {
	return IsType(err, TypeInternal)
}

func (e *Error) WithDetails(details any) *Error {
	e.Details = details
	return e
}

func (e *Error) WithCode(code string) *Error {
	e.Code = code
	return e
}

func IsType(err error, t ErrorType) bool {
	var e *Error
	return errors.As(err, &e) && e.Type == t
}

// package apperrors

// import (
// 	"errors"
// )

// // ErrorType categorizes the error
// type ErrorType string

// const (
// 	TypeValidation ErrorType = "VALIDATION"
// 	TypeNotFound   ErrorType = "NOT_FOUND"
// 	TypeConflict   ErrorType = "CONFLICT"
// 	TypeInternal   ErrorType = "INTERNAL"
// )

// // Error represents an application-level error
// type Error struct {
// 	Type    ErrorType `json:"type"`
// 	Message string    `json:"message"`
// 	Code    string    `json:"code,omitempty"`
// 	Details any       `json:"details,omitempty"`
// 	Err     error     `json:"-"`
// }

// // Error implements the error interface
// func (e *Error) Error() string {
// 	return e.Message
// }

// // Unwrap allows errors.Is / errors.As to work
// func (e *Error) Unwrap() error {
// 	return e.Err
// }

// //
// // Constructors
// //

// // New creates a new Error
// func New(t ErrorType, message string, err error) *Error {
// 	return &Error{
// 		Type:    t,
// 		Message: message,
// 		Err:     err,
// 	}
// }

// // Validation creates a validation error
// func Validation(message string) *Error {
// 	return New(TypeValidation, message, nil)
// }

// // NotFound creates a not found error
// func NotFound(message string) *Error {
// 	return New(TypeNotFound, message, nil)
// }

// // Conflict creates a conflict error
// func Conflict(message string) *Error {
// 	return New(TypeConflict, message, nil)
// }

// // Internal creates an internal error (can wrap another error)
// func Internal(message string, err error) *Error {
// 	return New(TypeInternal, message, err)
// }

// //
// // Helpers
// //

// // IsType checks if error is of a given type
// func IsType(err error, t ErrorType) bool {
// 	var e *Error
// 	return errors.As(err, &e) && e.Type == t
// }

// func IsValidation(err error) bool {
// 	return IsType(err, TypeValidation)
// }

// func IsNotFound(err error) bool {
// 	return IsType(err, TypeNotFound)
// }

// func IsConflict(err error) bool {
// 	return IsType(err, TypeConflict)
// }

// func IsInternal(err error) bool {
// 	return IsType(err, TypeInternal)
// }
