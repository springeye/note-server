package httputil

import (
	"github.com/go-chi/render"
	"net/http"
)

// NewResult example
func NewResult[T any](writer http.ResponseWriter, request *http.Request, data T) {
	re := HTTPResult[T]{
		Code:    0,
		Message: "success",
		Data:    data,
	}
	render.JSON(writer, request, &re)
}

// HTTPResult example
type HTTPResult[T any] struct {
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"success"`
	Data    T      `json:"data"`
}
