package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pierbin/screechr/internal/config"
	"github.com/pierbin/screechr/internal/models"
	"github.com/pierbin/screechr/internal/utils"
)

type ScreechrRepo struct {
	db *sql.DB
}

func NewScreechrRepo(cfg *config.Config) *ScreechrRepo {
	return &ScreechrRepo{db: InitDB(cfg)}
}

func (scRepo *ScreechrRepo) GetProfile(id int64, token string) (*models.Profile, error) {
	var profile models.Profile
	sqlStatement := `select id,username,firstname,lastname,token,profileimage,created,updated from profile where id = $1`

	err := scRepo.db.QueryRow(sqlStatement, id).Scan(&profile.Id, &profile.UserName, &profile.FirstName, &profile.LastName, &profile.Token, &profile.ProfileImage, &profile.Created, &profile.Updated)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if profile.Token != token {
		return nil, errors.New("Unauthorized")
	}

	return &profile, nil
}

func (scRepo *ScreechrRepo) UpdateProfile(id int64, token string, profile *models.Profile) error {
	var builderSql strings.Builder
	params := make([]interface{}, 0)

	builderSql.WriteString("update profile set ")
	if profile.UserName != "" {
		builderSql.WriteString("username=$1,")
		params = append(params, profile.UserName)
	}
	if profile.FirstName != "" {
		builderSql.WriteString("firstname=$2,")
		params = append(params, profile.FirstName)
	}
	if profile.LastName != "" {
		builderSql.WriteString("lastname=$3,")
		params = append(params, profile.LastName)
	}
	if profile.Token != "" {
		builderSql.WriteString("token=$4,")
		params = append(params, profile.Token)
	}
	if profile.ProfileImage != "" {
		builderSql.WriteString("profileimage=$5,")
		params = append(params, profile.ProfileImage)
	}

	builderSql.WriteString("updated=$6 ")
	params = append(params, utils.TimeNowUtcStr())

	builderSql.WriteString("where id=$7 and token=$8")
	params = append(params, id)
	params = append(params, token)

	_, err := scRepo.db.Exec(builderSql.String(), params...)
	if err != nil {
		return err
	}

	return nil
}

func (scRepo *ScreechrRepo) CreateScreech(id int64, screech *models.Screech) error {
	nowTime := time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")
	sql := "insert into screech(content,creatorid,created,updated) values($1,$2,$3,$4)"

	_, err := scRepo.db.Exec(sql, screech.Content, id, nowTime, nowTime)
	if err != nil {
		return err
	}

	return nil
}

func (scRepo *ScreechrRepo) GetScreech(id int64) (*models.Screech, error) {
	var screech models.Screech
	sqlStatement := `select id,content,creatorid,created,updated from screech where id = $1`

	err := scRepo.db.QueryRow(sqlStatement, id).Scan(&screech.Id, &screech.Content, &screech.CreatorId, &screech.Created, &screech.Updated)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &screech, nil
}

func (scRepo *ScreechrRepo) UpdateScreech(id int64, screech *models.Screech) error {
	var builderSql strings.Builder
	params := make([]interface{}, 0)

	builderSql.WriteString("update screech set ")
	if screech.Content != "" {
		builderSql.WriteString("content=$1,")
		params = append(params, screech.Content)
	}

	builderSql.WriteString("updated=$2 ")
	params = append(params, utils.TimeNowUtcStr())

	builderSql.WriteString("where id=$3")
	params = append(params, id)

	_, err := scRepo.db.Exec(builderSql.String(), params...)
	if err != nil {
		return err
	}

	return nil
}

func (scRepo *ScreechrRepo) GetScreechList(creatorid, size int64, order string) ([]models.Screech, error) {
	screeches := make([]models.Screech, 0)

	var builderSql strings.Builder

	builderSql.WriteString("select id,content,creatorid,created,updated from screech where 0=0 ")

	if creatorid != 0 {
		builderSql.WriteString("and creatorid=$1 ")
	}

	builderSql.WriteString(fmt.Sprintf("order by created %s ", order))
	builderSql.WriteString(fmt.Sprintf("limit %d ", size))

	rows, err := scRepo.db.Query(builderSql.String(), creatorid)
	if err != nil {
		if err == sql.ErrNoRows {
			return screeches, nil
		}
		return nil, err
	}
	defer rows.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for rows.Next() {
		var screech models.Screech
		err = rows.Scan(&screech.Id, &screech.Content, &screech.CreatorId, &screech.Created, &screech.Updated)
		if err != nil {
			return nil, err
		}
		screeches = append(screeches, screech)
	}

	return screeches, nil
}
