package campaign

import (
	internalErrors "email-dispatch-gateway/internal/internal-errors"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"gte=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (campaign *Campaign, err error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign = &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedOn: time.Now(),
	}

	err = internalErrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
