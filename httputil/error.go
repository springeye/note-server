package httputil

import (
	"github.com/go-chi/render"
	"net/http"
)

// NewError example
func NewError(writer http.ResponseWriter, request *http.Request, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	render.JSON(writer, request, &er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
