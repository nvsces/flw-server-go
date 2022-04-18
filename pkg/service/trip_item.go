package service

import (
	trip "github.com/nvsces/flw-server-go"
	"github.com/nvsces/flw-server-go/pkg/repository"
)

type TripItemService struct {
	repo repository.TripItem
}

func NewTripItemService(repo repository.TripItem) *TripItemService {
	return &TripItemService{repo: repo}
}

func (s *TripItemService) Create(item trip.TripItem) (int, error) {
	return s.repo.Create(item)
 }

func (s *TripItemService) GetAll() ([]repository.ObjectOutputJson, error) {
	return s.repo.GetAll()
}

func (s *TripItemService) GetById(itemId int) (trip.TripItem, error) {
	return s.repo.GetById(itemId)
}

func (s *TripItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
