package payloads

type CampaignDetailPayload struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	TotalDonation int64  `json:"total_donation"`
	TargetDonation int64 `json:"target_donation"`
	IsClosed      bool   `json:"is_closed"`
	Reason        string `json:"reason"`
}

type DonationReq struct {
	CampaignID int64 `json:"campaign_id"`
	DonorName string `json:"donor_name"`
	Amount int `json:"amount"`
}

type CreateDonationMessage struct {
	CampaignID int64 `json:"campaign_id"`
	DonorName string `json:"donor_name"`
	Amount int `json:"amount"`
	TargetDonation int64 `json:"target_donation"`
	TotalDonation int64 `json:"total_donation"`
}