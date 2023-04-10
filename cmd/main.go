package main

import (
	"telegram_bot/config"
	"telegram_bot/handler"
	"telegram_bot/service"
	"telegram_bot/storage/postgres"
)

func main() {
	cfg := config.Load()
	strg := postgres.NewPostgres(cfg)
	s := service.NewService(strg)
	h := handler.NewHandler(s)
	h.TelegramHandler()
}