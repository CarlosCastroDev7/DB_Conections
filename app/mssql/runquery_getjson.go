package mssql

import "fmt"

func Runquery_getjson(query, company string) (data string, err error) {
	db, err := connection_mssql(company)
	if err != nil {
		return "", err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var readRows string
		err := rows.Scan(&readRows)
		if err != nil {
			return "", err
		}
		data = fmt.Sprint(data, readRows)

	}

	return data, nil
}
