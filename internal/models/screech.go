package models

import (
	"time"
)

type Screech struct {
	Id        int64     `db:"id" json:"id"`
	Content   string    `db:"content" json:"content" binding:"required,max=1024"`
	CreatorId int64     `db:"creatorid" json:"creatorid"`
	Created   time.Time `db:"created" json:"created" time_format:"2006-01-02T15:04:05Z07:00"`
	Updated   time.Time `db:"updated" json:"updated" time_format:"2006-01-02T15:04:05Z07:00"`
}
