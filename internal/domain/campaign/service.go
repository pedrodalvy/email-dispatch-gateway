package campaign

import (
	"email-dispatch-gateway/internal/contract"
	internalerrors "email-dispatch-gateway/internal/internal-errors"
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

	err = s.Repository.Create(campaign)
	if err != nil {
		return "", internalerrors.ErrInternalServerError
	}

	return campaign.ID, err
}

func (s *Service) GetByID(id string) (contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetByID(id)
	if err != nil {
		return contract.CampaignResponse{}, internalerrors.ErrInternalServerError
	}

	return contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}

func (s *Service) CancelByID(id string) (err error) {
	campaign, err := s.Repository.GetByID(id)
	if err != nil {
		return internalerrors.ErrInternalServerError
	}

	err = campaign.Cancel()
	if err != nil {
		return err
	}

	err = s.Repository.Update(campaign)
	if err != nil {
		return internalerrors.ErrInternalServerError
	}

	return nil
}
