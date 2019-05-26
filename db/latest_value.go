package db

import (
	"database/sql"
	"path"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/Oppodelldog/balkonygardener/log"
)

func GetLatestValue(tableName string) (*SensorValue, error) {
	db, err := sql.Open("sqlite3", path.Join("./", config.Db.Filename))
	if err != nil {
		return nil, err
	}
	defer log.Error(db.Close())

	createTableStatement := `SELECT t,v FROM ` + tableName + ` ORDER BY t DESC LIMIT 1;`
	rows, err := db.Query(createTableStatement)
	if err != nil {
		return nil, err
	}

	rows.Next()
	fv := SensorValue{}
	log.Error(rows.Scan(&fv.T, &fv.V))

	return &fv, nil
}
