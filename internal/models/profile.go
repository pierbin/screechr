package models

import (
	"time"
)

type Profile struct {
	Id           int64     `db:"id" json:"id"`
	UserName     string    `db:"username" json:"username"`
	FirstName    string    `db:"firstname" json:"firstname"`
	LastName     string    `db:"lastname" json:"lastname"`
	Token        string    `db:"token" json:"token"`
	ProfileImage string    `db:"profileimage" json:"profileimage"`
	Created      time.Time `db:"created" json:"created" time_format:"2006-01-02T15:04:05Z07:00"`
	Updated      time.Time `db:"updated" json:"updated" time_format:"2006-01-02T15:04:05Z07:00"`
}
