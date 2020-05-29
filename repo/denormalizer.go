package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/budiariyanto/cqrs-dummy/model"
)

type DenormalizeRepo struct {
	db *sqlx.DB
}

func NewDenormalizeRepo(db *sqlx.DB) *DenormalizeRepo {
	return &DenormalizeRepo{
		db: db,
	}
}

func (d *DenormalizeRepo) GetCampaignDetail(id int64) (cd model.CampaignDetail, err error) {
	row := d.db.QueryRowx("SELECT * FROM campaign_detail WHERE id = ?", id)
	err = row.StructScan(&cd)
	return
}

func (d *DenormalizeRepo) UpdateDonatonStatus(campaignID int64, status bool, reason string) (err error) {
	_, err = d.db.Exec(`UPDATE campaign_detail 
					SET is_closed=?,
						reason=?
					WHERE id=?
					`, status, reason, campaignID)

	return
}

func (d *DenormalizeRepo) UpdateTotalDonation(campaignID, totalDonation int64) (err error) {
	_, err = d.db.Exec(`UPDATE campaign_detail SET total_donation = ? WHERE campaign_id = ?`, totalDonation, campaignID)
	return
}