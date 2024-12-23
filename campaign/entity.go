package campaign

import (
	"crowfundig/user"
	"time"
)

type Campaign struct {
	ID               int `gorm:"primaryKey"`
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	Slug             string
	BackerCount      int
	UserID           int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int `gorm:"primaryKey"`
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
