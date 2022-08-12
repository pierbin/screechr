package repo

import "github.com/pierbin/screechr/internal/models"

type IScreechrRepo interface {
	GetProfile(id int64, token string) (*models.Profile, error)
	UpdateProfile(id int64, token string, profile *models.Profile) error
	CreateScreech(id int64, screech *models.Screech) error
	GetScreech(id int64) (*models.Screech, error)
	UpdateScreech(id int64, screech *models.Screech) error
	GetScreechList(creatorid, size int64, order string) ([]models.Screech, error)
}
