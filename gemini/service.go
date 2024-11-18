package gemini

type GeminiService interface {
	SaveResponse(content string) (GeminiResponse, error)
	GetAllResponses() ([]GeminiResponse, error)
}

type geminiService struct {
	repository GeminiRepository
}

func NewGeminiService(repository GeminiRepository) *geminiService {
	return &geminiService{repository}
}

func (s *geminiService) SaveResponse(content string) (GeminiResponse, error) {
	response := GeminiResponse{
		Content: content,
	}

	savedResponse, err := s.repository.Save(response)
	if err != nil {
		return savedResponse, err
	}
	return savedResponse, nil
}

func (s *geminiService) GetAllResponses() ([]GeminiResponse, error) {
	return s.repository.FindAll()
}
