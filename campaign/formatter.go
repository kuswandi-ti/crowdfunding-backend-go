package campaign

import (
	"strings"
)

type CampaignFormatter struct {
	ID         			int    	`json:"id"`
	UserID				int   	`json:"user_id"`
	Name       			string 	`json:"name"`
	ShortDescription	string 	`json:"short_description"`
	ImageUrl      		string 	`json:"image_url"`
	GoalAmount      	int 	`json:"goal_amount"`
	CurrentAmount 		int   	`json:"current_amount"`
	Slug 				string 	`json:"slug"`
}

type CampaignDetailFormatter struct {
	ID 					int 						`json:"id"`
	Name 				string 						`json:"name"`
	ShortDescription 	string 						`json:"short_description"`
	Description 		string 						`json:"description"`
	ImageURL 			string 						`json:"image_url"`
	GoalAmount 			int 						`json:"goal_amount"`
	CurrentAmount 		int 						`json:"current_amount"`
	UserID 				int 						`json:"user_id"`
	Slug 				string 						`json:"slug"`
	Perks 				[]string 					`json:"perks"`
	User				CampaignUserFormatter 		`json:"user"`
	Images				[]CampaignImageFormatter	`json:"images"`
}

type CampaignUserFormatter struct {
	Name 		string 		`json:"name"`
	ImageURL	string 		`json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL	string 		`json:"image_url"`
	IsPrimary	bool 		`json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{}
	formatter.ID = campaign.ID
	formatter.UserID = campaign.UserID
	formatter.Name = campaign.Name
	formatter.ShortDescription = campaign.ShortDescription
	formatter.ImageUrl = ""
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}

	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageURL = ""
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserID = campaign.UserID

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, perk)
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName
	campaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImagesFormatter := CampaignImageFormatter{}
		campaignImagesFormatter.ImageURL = image.FileName
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImagesFormatter.IsPrimary = isPrimary
		images = append(images, campaignImagesFormatter)
	}
	campaignDetailFormatter.Images = images
	return campaignDetailFormatter
}