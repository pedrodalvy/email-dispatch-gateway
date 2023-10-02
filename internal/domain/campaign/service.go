package campaign

import (
	"email-dispatch-gateway/internal/contract"
	internalErrors "email-dispatch-gateway/internal/internal-errors"
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
		return "", internalErrors.ErrInternalServerError
	}

	return campaign.ID, err
}

func (s *Service) GetByID(id string) (contract.CampaignResponse, error) {
	campaign, err := s.findCampaignByID(id)
	if err != nil {
		return contract.CampaignResponse{}, err
	}

	return contract.CampaignResponse{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *Service) CancelByID(id string) (err error) {
	campaign, err := s.findCampaignByID(id)
	if err != nil {
		return err
	}

	if err = campaign.Cancel(); err != nil {
		return err
	}

	if err = s.Repository.Update(campaign); err != nil {
		return internalErrors.ErrInternalServerError
	}

	return nil
}

func (s *Service) DeleteByID(id string) (err error) {
	campaign, err := s.findCampaignByID(id)
	if err != nil {
		return err
	}

	if err = campaign.Delete(); err != nil {
		return err
	}

	if err = s.Repository.Delete(campaign); err != nil {
		return internalErrors.ErrInternalServerError
	}

	return nil
}

func (s *Service) findCampaignByID(id string) (*Campaign, error) {
	campaign, err := s.Repository.GetByID(id)

	if err != nil {
		return nil, internalErrors.ErrInternalServerError
	}

	if campaign == nil {
		return nil, internalErrors.ErrResourceNotFound
	}

	return campaign, nil
}
