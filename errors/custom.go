package errors

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

var (
	Unauthorized        = NewCustomError(559201002, StatusUnauthorized, "invalid authorization")
	InvalidArgument     = NewCustomError(559001001, StatusBadRequest, "invalid argument")
	InternalServerPanic = NewCustomError(559201000, StatusInternalServerError, "internal server panic")
	InternalServerError = NewCustomError(559201001, StatusInternalServerError, "internal server error")
)

type CustomError struct {
	code         int
	status       Status
	message      string
	errorMessage string
	cause        error
}

func NewCustomError(code int, status Status, message string) CustomError {
	return CustomError{
		code:    code,
		status:  status,
		message: message,
	}
}

func (e CustomError) New(msg string) error {
	e.cause = errors.New(msg)
	return errors.WithStack(e)
}

func (e CustomError) Errorf(format string, args ...interface{}) error {
	e.cause = errors.Errorf(format, args...)
	return errors.WithStack(e)
}

func (e CustomError) Error() string {
	if e.cause != nil {
		if e.errorMessage != "" {
			return fmt.Sprintf("%s: %v", e.errorMessage, e.cause.Error())
		} else {
			return e.cause.Error()
		}
	}
	if e.errorMessage != "" {
		return e.errorMessage
	}
	return e.message
}

func (e CustomError) Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	e.cause = err
	e.errorMessage = msg

	return errors.WithStack(e)
}

func (e CustomError) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	e.cause = err
	e.errorMessage = fmt.Sprintf(format, args...)

	return errors.WithStack(e)
}

// Cause 錯誤發生原因
func (e CustomError) Cause() error {
	return e.cause
}

// Message 取得對外錯誤訊息
func (e CustomError) Message() string {
	return e.message
}

// Code 取得錯誤碼
func (e CustomError) Code() int {
	return e.code
}

// Status 取得錯誤狀態
func (e CustomError) Status() Status {
	return e.status
}

func (e CustomError) IsEmpty() bool {
	return e == CustomError{}
}

func (e CustomError) Is(err error) bool {
	if err == nil {
		return false
	}

	if ae, ok := err.(CustomError); ok {
		return e.code == ae.code
	}
	return e.Error() == err.Error()
}

func (e CustomError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n", e.Cause())
			_, _ = io.WriteString(s, e.message)
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, e.Error())
	}
}

// CauseCustomError 解析內層符合 CustomError 結構
func CauseCustomError(err error) CustomError {
	type causer interface {
		Cause() error
	}

	var result CustomError

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		customError, ok := err.(CustomError)
		if ok {
			result = customError
		}

		err = cause.Cause()
	}

	return result
}
