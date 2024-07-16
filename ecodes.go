package ecodes

import "net/http"

const (
	OK            = http.StatusOK
	BadRequest    = http.StatusBadRequest
	Unauthorized  = http.StatusUnauthorized
	Forbidden     = http.StatusForbidden
	NotFound      = http.StatusNotFound
	InternalError = http.StatusInternalServerError
)
