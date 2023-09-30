package campaign

import "email-dispatch-gateway/internal/contract"

type ServiceInterface interface {
	Create(dto contract.NewCampaignDTO) (id string, err error)
	GetByID(id string) (campaignResponse contract.CampaignResponse, err error)
}
