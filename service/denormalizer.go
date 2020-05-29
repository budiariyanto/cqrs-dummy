package service

import (
	"github.com/budiariyanto/cqrs-dummy/repo"
	"github.com/budiariyanto/cqrs-dummy/model"
)

type DenormalizeService struct {
	repo *repo.DenormalizeRepo
}

func NewDenormalizeService(repo *repo.DenormalizeRepo) *DenormalizeService {
	return &DenormalizeService{
		repo: repo,
	}
}

func (s *DenormalizeService) GetCampaignDetail(id int64) (cd model.CampaignDetail, err error) {
	return s.repo.GetCampaignDetail(id)
}

func (s *DenormalizeService) UpdateCampaignStatus(campaignID int64, status bool, reason string) error {
	return s.repo.UpdateDonatonStatus(campaignID, status, reason)
}

func (s *DenormalizeService) UpdateTotalDonation(campaignID, totalDonation int64) error {
	return s.repo.UpdateTotalDonation(campaignID, totalDonation)
}