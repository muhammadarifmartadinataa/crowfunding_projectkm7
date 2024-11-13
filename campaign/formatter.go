package campaign

type CampaignFormater struct {
	ID               int    `json:id`
	UserID           int    `json:user_id`
	Name             string `json:name`
	ShortDescription string `json:short_description`
	ImageUrl         string `json:image_url`
	GoalAmount       int    `json:goal_amount`
	CurrentAmount    int    `json:current_amount`
}

func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormatter := CampaignFormater{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.ImageUrl = ""
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormater {
	campaignsFormatter := []CampaignFormater{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}
