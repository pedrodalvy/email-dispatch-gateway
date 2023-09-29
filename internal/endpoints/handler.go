package endpoints

import "email-dispatch-gateway/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.ServiceInterface
}

func NewHandler(cs campaign.ServiceInterface) Handler {
	return Handler{CampaignService: cs}
}
