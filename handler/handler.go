package handler

import "stad_projekt/service"

type Handler struct {
	service *service.Service
}

func Newhandler(service *service.Service) Handler {
	return Handler{
		service: service,
	}
}