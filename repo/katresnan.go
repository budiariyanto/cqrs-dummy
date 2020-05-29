package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/budiariyanto/cqrs-dummy/model"
)

type DonationRepo struct {
	db *sqlx.DB
}

func NewDonationRepo(db *sqlx.DB) *DonationRepo {
	return &DonationRepo{
		db: db,
	}
}

func (d *DonationRepo) Create(dn model.Donation) (err error) {
	_, err = d.db.NamedExec(`INSERT INTO donation (id, campaign_id, donor_name, amount)
					VALUES (:id, :campaign_id, :donor_name, :amount)`, dn)
	return
} 

func (d *DonationRepo) SumDonation(campaignID int64) (sum int64, err error) {
	row := d.db.QueryRowx("SELECT SUM(amount) FROM donation WHERE campaign_id = ?", campaignID)
	err = row.Scan(&sum)
	return
}