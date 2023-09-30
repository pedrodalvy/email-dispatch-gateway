package campaign

import (
	internalErrors "email-dispatch-gateway/internal/internal-errors"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	ID         string `validate:"required" gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignID string `gorm:"size:20"`
}

const (
	Pending string = "Pending"
)

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24"  gorm:"size:24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"  gorm:"size:1024"`
	Contacts  []Contact `validate:"gte=1,dive"`
	Status    string    `gorm:"size:20"`
}

func NewCampaign(name string, content string, emails []string) (campaign *Campaign, err error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}

	campaign = &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedOn: time.Now(),
		Status:    Pending,
	}

	err = internalErrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
