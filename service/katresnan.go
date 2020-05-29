package service

import (
	"github.com/budiariyanto/cqrs-dummy/repo"
	"github.com/budiariyanto/cqrs-dummy/model"
)

type DonationService struct {
	repo *repo.DonationRepo
}

func NewDonationService(repo *repo.DonationRepo) *DonationService {
	return &DonationService{
		repo: repo,
	}
}

func (s *DonationService) Create(d model.Donation) error {
	return s.repo.Create(d)
}

func (s *DonationService) Sum(campaignID int64) (int64, error) {
	return s.repo.SumDonation(campaignID)
}