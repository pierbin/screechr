package controller

import (
	"github.com/pierbin/screechr/internal/models"
	"github.com/pierbin/screechr/internal/repo"
	"go.uber.org/zap"
)

type ScreechrCtl struct {
	scRepo repo.IScreechrRepo
	logger *zap.Logger
}

func NewScreechrCtl(scRepo repo.IScreechrRepo) *ScreechrCtl {
	return &ScreechrCtl{
		scRepo: scRepo,
		logger: zap.NewExample(),
	}
}

func (scCtl *ScreechrCtl) GetProfile(id int64, token string) (*models.Profile, error) {
	profile, err := scCtl.scRepo.GetProfile(id, token)
	if err != nil {
		scCtl.logger.Error("GetProfile", zap.Error(err))
		return nil, err
	}

	return profile, nil
}

func (scCtl *ScreechrCtl) UpdateProfile(id int64, token string, profile *models.Profile) error {
	err := scCtl.scRepo.UpdateProfile(id, token, profile)
	if err != nil {
		scCtl.logger.Error("UpdateProfile", zap.Error(err))
		return err
	}

	return nil
}

func (scCtl *ScreechrCtl) CreateScreech(id int64, screech *models.Screech) error {
	err := scCtl.scRepo.CreateScreech(id, screech)
	if err != nil {
		scCtl.logger.Error("CreateScreech", zap.Error(err))
		return err
	}

	return nil
}

func (scCtl *ScreechrCtl) GetScreech(id int64) (*models.Screech, error) {
	screech, err := scCtl.scRepo.GetScreech(id)
	if err != nil {
		scCtl.logger.Error("GetScreech", zap.Error(err))
		return nil, err
	}

	return screech, nil
}

func (scCtl *ScreechrCtl) UpdateScreech(id int64, screech *models.Screech) error {
	err := scCtl.scRepo.UpdateScreech(id, screech)
	if err != nil {
		scCtl.logger.Error("UpdateScreech", zap.Error(err))
		return err
	}

	return nil
}

func (scCtl *ScreechrCtl) GetScreechList(creatorid, size int64, order string) ([]models.Screech, error) {
	screechs, err := scCtl.scRepo.GetScreechList(creatorid, size, order)
	if err != nil {
		scCtl.logger.Error("GetScreechList", zap.Error(err))
		return nil, err
	}

	return screechs, nil
}
