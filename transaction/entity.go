package transaction

import (
	"crowfundig/campaign"
	"crowfundig/user"
	"time"
)

type Transaction struct {
	ID         int `gorm:"primaryKey"`
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
