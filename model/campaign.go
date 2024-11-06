package models

import "time"

type Campaign struct {
	ID               uint            `gorm:"primaryKey;autoIncrement"`
	Name             string          `gorm:"type:varchar(255);not null"`
	ShortDescription string          `gorm:"type:varchar(255);not null"`
	Description      string          `gorm:"type:text;not null"`
	GoalAmount       int             `gorm:"not null"`
	CurrentAmount    int             `gorm:"not null"`
	Perks            string          `gorm:"type:text;not null"`
	Slug             string          `gorm:"type:varchar(255);unique;not null"`
	BackerCount      int             `gorm:"not null"`
	UserID           uint            `gorm:"index;not null"`
	CreatedAt        time.Time       `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time       `gorm:"type:datetime;not null"`
	DeletedAt        time.Time       `gorm:"type:datetime"`
	User             User            `gorm:"foreignKey:UserID"`
	CampaignImages   []CampaignImage `gorm:"foreignKey:CampaignID"`
	Transactions     []Transaction   `gorm:"foreignKey:CampaignID"`
}
