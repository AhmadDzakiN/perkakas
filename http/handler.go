package http

import (
	"net/http"
)

type HttpHandler struct {
	// H is handler, with return interface{} as data object, *string for token next page, error for error type
	H func(w http.ResponseWriter, r *http.Request) (interface{}, *string, error)
	CustomWriter
}

func NewHttpHandler(c HttpHandlerContext) func(handler func(w http.ResponseWriter, r *http.Request) (interface{}, *string, error)) HttpHandler {
	return func(handler func(w http.ResponseWriter, r *http.Request) (interface{}, *string, error)) HttpHandler {
		return HttpHandler{H: handler, CustomWriter: CustomWriter{C: c}}
	}
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, pageToken, err := h.H(w, r)
	if err != nil {
		h.WriteError(w, err)
		return
	}

	h.Write(w, data, pageToken)
}
