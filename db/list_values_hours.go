package db

import (
	"database/sql"
	"path"
	"time"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/Oppodelldog/balkonygardener/log"

	"github.com/sirupsen/logrus"
)

func GetValuesForHours(tableName string, hours int) (SensorValues, error) {
	db, err := sql.Open("sqlite3", path.Join("./", config.Db.Filename))
	if err != nil {
		return nil, err
	}
	defer log.Error(db.Close())

	dateFrom := time.Now().Add(time.Hour * time.Duration(-hours))
	dateTo := time.Now().Add(time.Hour)

	createTableStatement := `SELECT t,v FROM ` + tableName + ` WHERE t BETWEEN "` + dateFrom.Format("2006-01-02 15") + `" AND "` + dateTo.Format("2006-01-02 15") + `" ORDER BY t;`

	rows, err := db.Query(createTableStatement)
	if err != nil {
		logrus.Info(createTableStatement)
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
