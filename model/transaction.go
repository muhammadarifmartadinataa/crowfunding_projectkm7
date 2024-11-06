package models

import "time"

type Transaction struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Amount     int       `gorm:"not null"`
	Status     string    `gorm:"type:varchar(50);not null"`
	Code       string    `gorm:"type:varchar(255);not null"`
	UserID     uint      `gorm:"index;not null"`
	CampaignID uint      `gorm:"index;not null"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
	DeletedAt  time.Time `gorm:"type:datetime"`
	User       User      `gorm:"foreignKey:UserID"`
	Campaign   Campaign  `gorm:"foreignKey:CampaignID"`
}
