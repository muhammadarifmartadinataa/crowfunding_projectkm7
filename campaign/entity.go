package campaign

import "time"

type Campaign struct {
	ID                int
	Name              string
	ShortDescription  string
	Defaultescription string
	GoalAmount        int
	CurrentAmount     int
	Perks             string
	Slug              string
	BackerCount       int
	UserID            int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CampaignImages    []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
