package service

import (
	"github.com/budiariyanto/cqrs-dummy/repo"
)

type CampaignershipService struct {
	repo *repo.CampaignershipRepo
}

func NewCampaignershipService(repo *repo.CampaignershipRepo) *CampaignershipService {
	return &CampaignershipService{
		repo: repo,
	}
}

func (s *CampaignershipService) Close(id int64, reason string) error {
	return s.repo.CloseCampaign(id, reason)
	// publish close campaign event
}
