package models

import "time"

type CampaignImage struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	FileName   string    `gorm:"type:varchar(255);not null"`
	IsPrimary  int       `gorm:"type:tinyint;not null"`
	CampaignID uint      `gorm:"index;not null"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
	DeletedAt  time.Time `gorm:"type:datetime"`
	Campaign   Campaign  `gorm:"foreignKey:CampaignID"`
}
