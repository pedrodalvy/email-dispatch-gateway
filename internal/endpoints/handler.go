package endpoints

import "email-dispatch-gateway/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
