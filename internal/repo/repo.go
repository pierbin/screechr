package repo

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pierbin/screechr/internal/config"
)

var db *sql.DB

func conn(cfg *config.Config) (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	db, err := sql.Open(cfg.Driver, cfg.ConnString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(cfg *config.Config) *sql.DB {
	db, err := conn(cfg)
	if err != nil {
		panic("connect db error")
	}

	_, err = db.Exec(getProfileScript())
	if err != nil {
		panic("create db error")
	}
	_, err = db.Exec(getScreechScript())
	if err != nil {
		panic("create db error")
	}

	nowTime := time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")

	statement := "insert into profile (username,firstname,lastname,token,profileimage,created,updated) values('piershen','Pier','Shen','xYz123','/images/profile1.png',$1,$2)"
	_, err = db.Exec(statement, nowTime, nowTime)
	if err != nil {
		panic("create data error")
	}

	statement = "insert into profile (username,firstname,lastname,token,profileimage,created,updated) values('jasonshen','Jason','Shen','aBc123','/images/profile2.png',$1,$2)"
	_, err = db.Exec(statement, nowTime, nowTime)
	if err != nil {
		panic("create data error")
	}

	return db
}

func getProfileScript() string {
	return `CREATE TABLE "profile" (
		"id"	INTEGER NOT NULL UNIQUE,
		"username"	TEXT NOT NULL UNIQUE,
		"firstname"	TEXT NOT NULL,
		"lastname"	TEXT NOT NULL,
		"token"	TEXT NOT NULL UNIQUE,
		"profileimage"	TEXT,
		"created"	datetime NOT NULL,
		"updated"	datetime NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`
}

func getScreechScript() string {
	return `CREATE TABLE "screech" (
		"id"	INTEGER NOT NULL UNIQUE,
		"content"	TEXT NOT NULL,
		"creatorid"	INTEGER NOT NULL,
		"created"	datetime NOT NULL,
		"updated"	datetime NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT),
		CONSTRAINT "fk_creatorid_profile_id" FOREIGN KEY("creatorid") REFERENCES "profile"("id")
	)`
}
