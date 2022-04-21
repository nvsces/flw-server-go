package service

import (
	trip "github.com/nvsces/flw-server-go"
	"github.com/nvsces/flw-server-go/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) Info(userId int) (trip.User, error) {
	return s.repo.Info(userId)
}
