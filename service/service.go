package service

import "stad_projekt/storage"

type Service struct {
	s storage.StorageI
}

func NewService (repository storage.StorageI) *Service{
	return &Service{
		s: repository,
	}
}