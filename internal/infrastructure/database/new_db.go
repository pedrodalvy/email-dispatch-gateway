package database

import (
	"email-dispatch-gateway/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATA_SOURCE_NAME")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
