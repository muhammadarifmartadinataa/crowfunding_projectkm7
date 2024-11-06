package models

import "time"

type User struct {
	ID             uint          `gorm:"primaryKey;autoIncrement"`
	Name           string        `gorm:"type:varchar(255);not null"`
	Occupation     string        `gorm:"type:varchar(255);not null"`
	Email          string        `gorm:"type:varchar(255);unique;not null"`
	PasswordHash   string        `gorm:"type:varchar(255);not null"`
	AvatarFileName string        `gorm:"type:varchar(255);not null"`
	Role           string        `gorm:"type:varchar(50);not null"`
	Token          string        `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time     `gorm:"type:datetime;not null"`
	UpdatedAt      time.Time     `gorm:"type:datetime;not null"`
	DeletedAt      time.Time     `gorm:"type:datetime"`
	Campaigns      []Campaign    `gorm:"foreignKey:UserID"`
	Transactions   []Transaction `gorm:"foreignKey:UserID"`
}
