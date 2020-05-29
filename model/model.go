package model

type Donation struct {
	ID int64 `db:"id" json:"id"`
	CampaignID int64 `db:"campaign_id" json:"campaign_id"`
	DonorName string `db:"donor_name" json:"donor_name"`
	Amount int64 `db:"amount" json:"amount"`
}

type Campaign struct {
	ID int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	TargetDonation int64 `db:"target_donation" json:"target_donation"`
	IsClosed bool `db:"is_closed" json:"is_closed"`
	Reason string `db:"reason" json:"reason"`
}

type CampaignDetail struct {
	ID            int64  `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	TotalDonation int64  `db:"total_donation" json:"total_donation"`
	TargetDonation int64 `db:"target_donation" json:"target_donation"`
	IsClosed      bool   `db:"is_closed" json:"is_closed"`
	Reason        string `db:"reason" json:"reason"`
}