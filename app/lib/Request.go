package lib

import (
	"net/http"
)

// Request ...
type Request struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}
