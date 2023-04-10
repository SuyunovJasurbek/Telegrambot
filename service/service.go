package service

import "telegram_bot/storage"

type Service struct {
	repository storage.StorageI
}

func NewService(repository storage.StorageI) *Service {
	return &Service{
		repository: repository,
	}
}