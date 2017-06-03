package db

import (
	"database/sql"
	"path"
)

func ListValues(tableName string) (SensorValues, error) {
	db, err := sql.Open("sqlite3", path.Join("./", dbFileName))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	createTableStatement := `SELECT t,v FROM ` + tableName + ` ORDER BY t;`
	rows, err := db.Query(createTableStatement)
	if err != nil {
		return nil, err
	}

	floatValues := SensorValues{}

	for rows.Next() {
		fv := SensorValue{}
		rows.Scan(&fv.T, &fv.V)
		floatValues = append(floatValues, fv)
	}

	return floatValues, nil
}
