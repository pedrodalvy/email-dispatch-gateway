package database

import (
	"email-dispatch-gateway/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	DB *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{DB: db}
}

func (cr *CampaignRepository) Create(campaign *campaign.Campaign) error {
	tx := cr.DB.Create(campaign)
	return tx.Error
}

func (cr *CampaignRepository) GetByID(id string) (campaign *campaign.Campaign, err error) {
	tx := cr.DB.First(&campaign, "id = ?", id)
	return campaign, tx.Error
}

func (cr *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := cr.DB.Save(campaign)
	return tx.Error
}

func (cr *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	tx := cr.DB.Select("Contacts").Delete(campaign)
	return tx.Error
}
