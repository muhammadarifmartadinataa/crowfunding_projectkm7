package user

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"type:varchar(255)"`
	Occupation     string    `gorm:"type:varchar(255)"`
	Email          string    `gorm:"type:varchar(255);unique"`
	PasswordHash   string    `gorm:"type:varchar(255)"`
	AvatarFileName string    `gorm:"type:varchar(255)"`
	Role           string    `gorm:"type:varchar(50)"`
	Token          string    `gorm:"type:varchar(255)"`
	CreatedAt      time.Time `gorm:"type:datetime"`
	UpdatedAt      time.Time `gorm:"type:datetime"`
	// Campaigns      []Campaign    `gorm:"foreignKey:UserID"`
	// Transactions   []Transaction `gorm:"foreignKey:UserID"`
}
