package handler

import (
	"telegram_bot/service"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) Handler {
	return Handler{
		s: s,
	}
}