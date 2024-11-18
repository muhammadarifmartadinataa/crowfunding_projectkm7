package gemini

import "time"

type GeminiResponse struct {
	ID        int    `gorm:"primaryKey"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
