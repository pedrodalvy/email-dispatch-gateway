package database

import (
	"email-dispatch-gateway/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	cr.campaigns = append(cr.campaigns, *campaign)
	return nil
}

func (cr *CampaignRepository) GetByID(id string) (campaign *campaign.Campaign, err error) {
	for _, c := range cr.campaigns {
		if c.ID == id {
			return &c, nil
		}
	}

	return nil, errors.New("campaign not found")
}
