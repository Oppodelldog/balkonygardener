package db

import (
	"database/sql"
	"path"
	"time"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/Oppodelldog/balkonygardener/log"
)

func SaveString(name string, value string) error {
	db, err := sql.Open("sqlite3", path.Join("./", config.Db.Filename))
	if err != nil {
		return err
	}
	defer func() { log.Error(db.Close()) }()

	createTableStatement := `create table if not exists ` + name + ` (t datetime, v string)`
	_, err = db.Exec(createTableStatement)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into " + name + "(t, v) values(?, ?)")
	if err != nil {
		return err
	}
	defer func() { log.Error(stmt.Close()) }()
	_, err = stmt.Exec(time.Now(), value)
	if err != nil {
		return err
	}
	err = tx.Commit()

	return err
}
