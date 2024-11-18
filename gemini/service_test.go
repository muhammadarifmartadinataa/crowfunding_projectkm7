package gemini

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock GeminiRepository
type MockGeminiRepository struct {
	mock.Mock
}

func (m *MockGeminiRepository) Save(response GeminiResponse) (GeminiResponse, error) {
	args := m.Called(response)
	return args.Get(0).(GeminiResponse), args.Error(1)
}

func (m *MockGeminiRepository) FindAll() ([]GeminiResponse, error) {
	args := m.Called()
	return args.Get(0).([]GeminiResponse), args.Error(1)
}

func TestSaveResponse_Success(t *testing.T) {
	mockRepo := new(MockGeminiRepository)
	service := NewGeminiService(mockRepo)

	// Mock data
	content := "Test response"
	mockResponse := GeminiResponse{
		Content: content,
	}

	// Mock behavior
	mockRepo.On("Save", mock.Anything).Return(mockResponse, nil)

	// Call the method
	result, err := service.SaveResponse(content)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockResponse, result)
	mockRepo.AssertExpectations(t)
}

func TestSaveResponse_Failure(t *testing.T) {
	mockRepo := new(MockGeminiRepository)
	service := NewGeminiService(mockRepo)

	// Mock data
	content := "Test response"
	mockError := errors.New("Failed to save response")

	// Mock behavior
	mockRepo.On("Save", mock.Anything).Return(GeminiResponse{}, mockError)

	// Call the method
	result, err := service.SaveResponse(content)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "Failed to save response", err.Error())
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllResponses_Success(t *testing.T) {
	mockRepo := new(MockGeminiRepository)
	service := NewGeminiService(mockRepo)

	// Mock data
	mockResponses := []GeminiResponse{
		{Content: "Response 1"},
		{Content: "Response 2"},
	}

	// Mock behavior
	mockRepo.On("FindAll").Return(mockResponses, nil)

	// Call the method
	result, err := service.GetAllResponses()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockResponses, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllResponses_Failure(t *testing.T) {
	mockRepo := new(MockGeminiRepository)
	service := NewGeminiService(mockRepo)

	// Mock data
	mockError := errors.New("Failed to fetch responses")

	// Mock behavior
	mockRepo.On("FindAll").Return(nil, mockError)

	// Call the method
	result, err := service.GetAllResponses()

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "Failed to fetch responses", err.Error())
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
