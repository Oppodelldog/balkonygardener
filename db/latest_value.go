package db

import (
	"database/sql"
	"path"
)

func GetLatestValue(tableName string) (*SensorValue, error) {
	db, err := sql.Open("sqlite3", path.Join("./", dbFileName))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	createTableStatement := `SELECT t,v FROM ` + tableName + ` ORDER BY t DESC LIMIT 1;`
	rows, err := db.Query(createTableStatement)
	if err != nil {
		return nil, err
	}

	rows.Next()
	fv := SensorValue{}
	rows.Scan(&fv.T, &fv.V)

	return &fv, nil
}
