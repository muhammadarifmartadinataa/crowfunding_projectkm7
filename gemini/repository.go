package gemini

import (
	"gorm.io/gorm"
)

type GeminiRepository interface {
	Save(response GeminiResponse) (GeminiResponse, error)
	FindAll() ([]GeminiResponse, error)
}

type geminiRepository struct {
	db *gorm.DB
}

func NewGeminiRepository(db *gorm.DB) *geminiRepository {
	return &geminiRepository{db}
}

func (r *geminiRepository) Save(response GeminiResponse) (GeminiResponse, error) {
	err := r.db.Create(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *geminiRepository) FindAll() ([]GeminiResponse, error) {
	var responses []GeminiResponse
	err := r.db.Find(&responses).Error
	return responses, err
}
