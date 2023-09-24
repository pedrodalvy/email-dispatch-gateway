package campaign

import (
	"email-dispatch-gateway/internal/contract"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) Create(dto contract.NewCampaignDTO) (id string, err error) {
	campaign, err := NewCampaign(dto.Name, dto.Content, dto.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", err
	}

	return campaign.ID, err
}
