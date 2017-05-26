package db

import (
	"database/sql"
	"path"
	"time"
)

type FloatValues []FloatValue
type FloatValue struct{
	T time.Time
	V interface{}
}
func ListFloatValues(tableName string) (FloatValues, error) {
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

	floatValues := FloatValues{}

	for rows.Next() {
		fv := FloatValue{}
		rows.Scan(&fv.T,&fv.V)
		floatValues = append(floatValues, fv)
	}

	return floatValues, nil
}
