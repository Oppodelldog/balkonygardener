package db

import (
	"database/sql"
	"path"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func SaveFloat(name string, value float64) error{
	db, err := sql.Open("sqlite3", path.Join("./",dbFileName))
	if err != nil {
		return err
	}
	defer db.Close()

	createTableStatement := `create table if not exists ` + name + ` (t datetime, v float)`
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
	defer stmt.Close()
	_, err = stmt.Exec(time.Now(), value)
	if err != nil {
		return err
	}
	err = tx.Commit()

	return err
}
