package controller

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pierbin/screechr/internal/models"
	"github.com/pierbin/screechr/internal/repo"
)

func initProfile() (*models.Profile, *models.Profile) {
	return &models.Profile{
			Id:           1,
			UserName:     "piershen",
			FirstName:    "Pier",
			LastName:     "Shen",
			Token:        "xYz123",
			ProfileImage: "/images/profile1.png",
			Created:      time.Now().UTC(),
			Updated:      time.Now().UTC(),
		},
		&models.Profile{
			Id:           2,
			UserName:     "jasonshen",
			FirstName:    "Jason",
			LastName:     "Shen",
			Token:        "aBc123",
			ProfileImage: "/images/profile2.png",
			Created:      time.Now().UTC(),
			Updated:      time.Now().UTC(),
		}
}

func initScreech() (*models.Screech, *models.Screech) {
	return &models.Screech{
			Id:        1,
			Content:   "content 1",
			CreatorId: 1,
			Created:   time.Now().UTC(),
			Updated:   time.Now().UTC(),
		},
		&models.Screech{
			Id:        2,
			Content:   "content 2",
			CreatorId: 2,
			Created:   time.Now().UTC(),
			Updated:   time.Now().UTC(),
		}
}

func TestGetProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id1 int64 = 1
	var id2 int64 = 3
	profile1, _ := initProfile()

	mockIScreechrRepo := repo.NewMockIScreechrRepo(ctl)
	gomock.InOrder(
		mockIScreechrRepo.EXPECT().GetProfile(gomock.Eq(id1), gomock.Eq("xYz123")).Return(profile1, nil),
		mockIScreechrRepo.EXPECT().GetProfile(gomock.Eq(id1), gomock.Not("xYz123")).Return(nil, errors.New("Unauthorized")),
		mockIScreechrRepo.EXPECT().GetProfile(gomock.Eq(id2), gomock.Any()).Return(nil, nil),
	)

	screechrCtl := NewScreechrCtl(mockIScreechrRepo)
	profile, err := screechrCtl.GetProfile(id1, "xYz123")
	if err != nil || profile == nil {
		t.Errorf("err %v", err)
	}

	if profile.Id != profile1.Id {
		t.Error("profile should be matched")
	}

	profile, err = screechrCtl.GetProfile(id1, "")
	if profile != nil {
		t.Error("profile should be null")
	}
	if err != nil && err.Error() != "Unauthorized" {
		t.Error("err should be Unauthorized")
	}

	profile, err = screechrCtl.GetProfile(id2, "xYz123")
	if profile != nil || err != nil {
		t.Error("profile and err should be null")
	}
}

func TestUpdateProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id1 int64 = 1
	var updatedProfile = &models.Profile{
		ProfileImage: "/images/profilex.png",
	}
	profile1, _ := initProfile()

	mockIScreechrRepo := repo.NewMockIScreechrRepo(ctl)

	mockIScreechrRepo.EXPECT().UpdateProfile(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(id int64, token string, profile *models.Profile) error {
		if id == 1 && token == "xYz123" {
			profile1.ProfileImage = profile.ProfileImage
		}
		return nil
	}).Times(2)

	screechrCtl := NewScreechrCtl(mockIScreechrRepo)

	err := screechrCtl.UpdateProfile(id1, "xYz123", updatedProfile)
	if err != nil {
		t.Errorf("err %v", err)
	}
	if profile1.ProfileImage != updatedProfile.ProfileImage {
		t.Error("profile image should be updated")
	}

	profile1, _ = initProfile()

	err = screechrCtl.UpdateProfile(id1, "", updatedProfile)
	if err != nil {
		t.Errorf("err %v", err)
	}
	if profile1.ProfileImage == updatedProfile.ProfileImage {
		t.Error("profile image should not be updated")
	}
}

func TestCreateScreech(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id1 int64 = 1
	var id2 int64 = 3
	screech1, _ := initScreech()
	newScreech := new(models.Screech)

	mockIScreechrRepo := repo.NewMockIScreechrRepo(ctl)

	mockIScreechrRepo.EXPECT().CreateScreech(gomock.Any(), gomock.Any()).DoAndReturn(func(id int64, screech *models.Screech) error {
		if id == 3 {
			return errors.New("not exist profile")
		}

		if id == 1 {
			newScreech = &models.Screech{
				Id:        screech.Id,
				Content:   screech.Content,
				CreatorId: screech.CreatorId,
				Created:   screech.Created,
				Updated:   screech.Updated,
			}
		}
		return nil
	}).Times(2)

	screechrCtl := NewScreechrCtl(mockIScreechrRepo)

	err := screechrCtl.CreateScreech(id1, screech1)
	if err != nil {
		t.Errorf("err %v", err)
	}
	if newScreech.Id == 0 {
		t.Error("screech should be created")
	}

	newScreech = new(models.Screech)

	err = screechrCtl.CreateScreech(id2, screech1)
	if err == nil {
		t.Error("profile image should not be created")
	}
}

func TestGetScreech(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id1 int64 = 1
	var id2 int64 = 10

	screech1, _ := initScreech()

	mockIScreechrRepo := repo.NewMockIScreechrRepo(ctl)
	mockIScreechrRepo.EXPECT().GetScreech(gomock.Eq(id1)).Return(screech1, nil)
	mockIScreechrRepo.EXPECT().GetScreech(gomock.Eq(id2)).Return(nil, nil)

	screechrCtl := NewScreechrCtl(mockIScreechrRepo)
	screech, err := screechrCtl.GetScreech(id1)
	if err != nil || screech == nil {
		t.Errorf("err %v", err)
	}

	if screech.Id != screech1.Id {
		t.Error("screech should be matched")
	}

	screech, err = screechrCtl.GetScreech(id2)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("err %v", err)
	}
	if screech != nil {
		t.Error("screech should be null")
	}
}

func TestUpdateScreech(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id1 int64 = 1
	var id2 int64 = 10
	var updatedScreech = &models.Screech{
		Content: "content x",
	}
	screech1, _ := initScreech()

	mockIScreechrRepo := repo.NewMockIScreechrRepo(ctl)

	mockIScreechrRepo.EXPECT().UpdateScreech(gomock.Any(), gomock.Any()).DoAndReturn(func(id int64, screech *models.Screech) error {
		if id == 1 {
			screech1.Content = screech.Content
		}
		return nil
	}).Times(2)

	screechrCtl := NewScreechrCtl(mockIScreechrRepo)

	err := screechrCtl.UpdateScreech(id1, updatedScreech)
	if err != nil {
		t.Errorf("err %v", err)
	}
	if screech1.Content != updatedScreech.Content {
		t.Error("content should be updated")
	}

	screech1, _ = initScreech()

	err = screechrCtl.UpdateScreech(id2, updatedScreech)
	if err != nil {
		t.Errorf("err %v", err)
	}
	if screech1.Content == updatedScreech.Content {
		t.Error("content should not be updated")
	}
}
