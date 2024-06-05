package handler

import (
	service "project/service"
)

type HTTPHandler struct {
	Service *service.Service
}

func NewHTTPHandler(service service.Service) *HTTPHandler {
	return &HTTPHandler{Service: &service}
}
