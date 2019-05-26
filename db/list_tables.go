package db

import (
	"database/sql"
	"path"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/Oppodelldog/balkonygardener/log"
)

func ListTables() ([]string, error) {
	db, err := sql.Open("sqlite3", path.Join("./", config.Db.Filename))
	if err != nil {
		return nil, err
	}
	defer func() { log.Error(db.Close()) }()

	createTableStatement := `SELECT name FROM sqlite_master WHERE type='table';`
	rows, err := db.Query(createTableStatement)
	if err != nil {
		return nil, err
	}

	var tableNames []string
	var tableName string
	for rows.Next() {
		_ = rows.Scan(&tableName)
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}
