package transaction

import (
	"campaign/campaign"
	"campaign/user"
	"time"
)

type Transaction struct {
	ID         int    `json:"id" gorm:"primary_key; autoincrement"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	User       user.User
	Campaign   campaign.Campaign
	PaymentURL string    `json:"payment_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
