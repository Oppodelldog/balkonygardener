package db

import (
	"database/sql"
	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/Oppodelldog/balkonygardener/log"
	"path"
)

func ListValues(tableName string) (SensorValues, error) {
	db, err := sql.Open("sqlite3", path.Join("./", config.Db.Filename))
	if err != nil {
		return nil, err
	}
	defer log.Error(db.Close())

	createTableStatement := `SELECT t,v FROM ` + tableName + ` ORDER BY t;`
	rows, err := db.Query(createTableStatement)
	if err != nil {
		return nil, err
	}

	floatValues := SensorValues{}

	for rows.Next() {
		fv := SensorValue{}
		log.Error(rows.Scan(&fv.T, &fv.V))
		floatValues = append(floatValues, fv)
	}

	return floatValues, nil
}
