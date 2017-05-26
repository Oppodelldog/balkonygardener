package db

import (
	"database/sql"
	"path"
)

func ListTables() ([]string, error) {
	db, err := sql.Open("sqlite3", path.Join("./", dbFileName))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	createTableStatement := `SELECT name FROM sqlite_master WHERE type='table';`
	rows, err := db.Query(createTableStatement)
	if err != nil {
		return nil, err
	}

	tableNames := []string{}
	var tableName string
	for rows.Next() {
		rows.Scan(&tableName)
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}
