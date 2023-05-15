package service

import "stad_projekt/storage"

type Service struct {
	repo storage.StorageI
}

func NewService (repository storage.StorageI) *Service{
	return &Service{
		repo: repository,
	}
}