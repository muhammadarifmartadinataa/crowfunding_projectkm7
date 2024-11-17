package campaign

import (
	"crowfundig/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindAll() ([]Campaign, error) {
	args := m.Called()
	return args.Get(0).([]Campaign), args.Error(1)
}

func (m *MockRepository) FindByUserID(userID int) ([]Campaign, error) {
	args := m.Called(userID)
	return args.Get(0).([]Campaign), args.Error(1)
}

func (m *MockRepository) FindByID(ID int) (Campaign, error) {
	args := m.Called(ID)
	return args.Get(0).(Campaign), args.Error(1)
}

func (m *MockRepository) Save(campaign Campaign) (Campaign, error) {
	args := m.Called(campaign)
	return args.Get(0).(Campaign), args.Error(1)
}

func (m *MockRepository) Update(campaign Campaign) (Campaign, error) {
	args := m.Called(campaign)
	return args.Get(0).(Campaign), args.Error(1)
}

func (m *MockRepository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	args := m.Called(campaignImage)
	return args.Get(0).(CampaignImage), args.Error(1)
}

func (m *MockRepository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	args := m.Called(campaignID)
	return args.Bool(0), args.Error(1)
}

// Test GetCampaigns
func TestGetCampaigns(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	campaigns := []Campaign{
		{ID: 1, Name: "Campaign 1", UserID: 1},
		{ID: 2, Name: "Campaign 2", UserID: 2},
	}

	mockRepo.On("FindAll").Return(campaigns, nil)

	result, err := service.GetCampaigns(0)

	assert.NoError(t, err)
	assert.Equal(t, len(campaigns), len(result))
	mockRepo.AssertExpectations(t)
}

// Test CreateCampaign
func TestCreateCampaign(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	input := CreateCampaignInput{
		Name:             "New Campaign",
		ShortDescription: "Short Description",
		Description:      "Full Description",
		Perks:            "Perk1, Perk2",
		GoalAmount:       10000,
		User:             user.User{ID: 1},
	}

	campaign := Campaign{
		ID:               1,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		UserID:           int(input.User.ID),
		Slug:             "new-campaign-1",
	}

	mockRepo.On("Save", mock.Anything).Return(campaign, nil)

	result, err := service.CreateCampaign(input)

	assert.NoError(t, err)
	assert.Equal(t, campaign.ID, result.ID)
	assert.Equal(t, campaign.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

// Test UpdateCampaign
func TestUpdateCampaign(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	inputID := GetCampaignDetailInput{ID: 1}
	inputData := CreateCampaignInput{
		Name:             "Updated Campaign",
		ShortDescription: "Updated Short Description",
		Description:      "Updated Full Description",
		Perks:            "Perk1, Perk2",
		GoalAmount:       20000,
		User:             user.User{ID: 1},
	}

	existingCampaign := Campaign{
		ID:               1,
		Name:             "Old Campaign",
		UserID:           1,
		ShortDescription: "Old Short Description",
		Description:      "Old Full Description",
		Perks:            "Perk1, Perk2",
		GoalAmount:       10000,
	}

	updatedCampaign := Campaign{
		ID:               1,
		Name:             inputData.Name,
		UserID:           1,
		ShortDescription: inputData.ShortDescription,
		Description:      inputData.Description,
		Perks:            inputData.Perks,
		GoalAmount:       inputData.GoalAmount,
	}

	mockRepo.On("FindByID", inputID.ID).Return(existingCampaign, nil)
	mockRepo.On("Update", mock.Anything).Return(updatedCampaign, nil)

	result, err := service.UpdateCampaign(inputID, inputData)

	assert.NoError(t, err)
	assert.Equal(t, updatedCampaign.Name, result.Name)
	assert.Equal(t, updatedCampaign.GoalAmount, result.GoalAmount)
	mockRepo.AssertExpectations(t)
}
