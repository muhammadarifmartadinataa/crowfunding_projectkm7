package config

import (
	"crowfundig/campaign"
	"crowfundig/transaction"
	"crowfundig/user"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &campaign.Campaign{}, &campaign.CampaignImage{}, &transaction.Transaction{})
}
