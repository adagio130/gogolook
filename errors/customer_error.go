package errors

// general error
var (
	InvalidRequest     = NewCustomError(400, StatusBadRequest, "request is invalid")
	Internal           = NewCustomError(500, StatusInternalServerError, "server internal error")
	ServiceUnavailable = NewCustomError(503, StatusServiceUnavailable, "service unavailable")
	Timeout            = NewCustomError(504, StatusGatewayTimeout, "server timeout")
)
