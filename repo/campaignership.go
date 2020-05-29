package repo

import (
	"github.com/jmoiron/sqlx"
)

type CampaignershipRepo struct {
	db *sqlx.DB
}

func NewCampaignershipRepo(db *sqlx.DB) *CampaignershipRepo {
	return &CampaignershipRepo{
		db: db,
	}
}

func (c *CampaignershipRepo) CloseCampaign(id int64, reason string) (err error) {
	_, err = c.db.Exec("UPDATE campaign SET is_closed=true, reason=? WHERE id=?", id, reason)
	return
}
