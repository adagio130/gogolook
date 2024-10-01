package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"tasks/errors"
)

const (
	ContextKeyError      = "error"
	ContextKeyCode       = "code"
	ContextKeyStackTrace = "stackTrace"
	ContextKeyTraceId    = "traceId"
	ContextKeyCustomLog  = "customLog"
	ContextKeyRespStatus = "respStatus"
	ContextKeyRespType   = "respType"
	ContextKeyResp       = "resp"
	ContextKeyUserID     = "user_id"
	RespTypeJSON         = "json"
	RespTypePlain        = "plain"
)

// Default Error Info
const (
	defaultErrorStatus  = http.StatusInternalServerError
	defaultErrorCode    = 99998
	defaultErrorMessage = "internal server error"
)

// Panic Error Info
const (
	panicErrorStatus  = http.StatusInternalServerError
	panicErrorCode    = 99999
	panicErrorMessage = "server panic error"
)

const (
	traceHeaderKey = "X-Trace-Id"
)

type ResponseMiddleware struct{}

func NewResponseMiddleware() *ResponseMiddleware {
	return &ResponseMiddleware{}
}

func (m *ResponseMiddleware) GetResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetString(ContextKeyTraceId)
		c.Writer.Header().Set(traceHeaderKey, traceID)

		defer func() {
			if err := recover(); err != nil {
				c.Set(ContextKeyError, fmt.Errorf("panic: %s", err).Error())
				c.Set(ContextKeyCode, panicErrorCode)
				c.Set(ContextKeyStackTrace, string(debug.Stack()))

				c.JSON(
					panicErrorStatus,
					m.makeErrorResp(panicErrorCode, panicErrorMessage, traceID),
				)
				return
			}
		}()

		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err

			customError := errors.CauseCustomError(err)
			if customError.IsEmpty() {
				c.Set(ContextKeyError, err.Error())
				c.Set(ContextKeyCode, defaultErrorCode)
				c.Set(ContextKeyStackTrace, errors.CauseStackTrace(err))

				c.JSON(
					defaultErrorStatus,
					m.makeErrorResp(defaultErrorCode, defaultErrorMessage, traceID),
				)
				return
			}

			c.Set(ContextKeyError, err.Error())
			c.Set(ContextKeyCode, customError.Code())
			c.Set(ContextKeyStackTrace, errors.CauseStackTrace(err))
			c.JSON(
				customError.Status().ToHTTPStatus(),
				m.makeErrorResp(customError.Code(), customError.Message(), traceID),
			)
			return
		}
	}
}

func (m *ResponseMiddleware) makeErrorResp(code int, message string, traceID string) gin.H {
	return gin.H{
		"error": gin.H{
			"code":     code,
			"message":  message,
			"trace_id": traceID,
		},
	}
}
