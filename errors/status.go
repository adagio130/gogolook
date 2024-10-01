package errors

import (
	"net/http"
)

type Status string

const (
	StatusBadRequest          Status = "BadRequest"
	StatusUnauthorized        Status = "Unauthorized"
	StatusForbidden           Status = "Forbidden"
	StatusNotFound            Status = "NotFound"
	StatusTooManyRequests     Status = "TooManyRequests"
	StatusBadGateway          Status = "BadGateway"
	StatusInternalServerError Status = "InternalServerError"
	StatusServiceUnavailable  Status = "ServiceUnavailable"
	StatusGatewayTimeout      Status = "GatewayTimeout"
	StatusPaymentRequire      Status = "PaymentRequire"
)

func (s Status) ToHTTPStatus() int {
	switch s {
	case StatusBadRequest:
		return http.StatusBadRequest
	case StatusUnauthorized:
		return http.StatusUnauthorized
	case StatusForbidden:
		return http.StatusForbidden
	case StatusNotFound:
		return http.StatusNotFound
	case StatusTooManyRequests:
		return http.StatusTooManyRequests
	case StatusBadGateway:
		return http.StatusBadGateway
	case StatusInternalServerError:
		return http.StatusInternalServerError
	case StatusServiceUnavailable:
		return http.StatusServiceUnavailable
	case StatusGatewayTimeout:
		return http.StatusGatewayTimeout
	case StatusPaymentRequire:
		return http.StatusPaymentRequired
	default:
		return http.StatusInternalServerError
	}
}
