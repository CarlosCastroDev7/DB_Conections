package mssql

import (
	"encoding/json"
)

func Runquery_getdata(query, company string) (data string, err error) {

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

	cols, err := rows.Columns()
	if err != nil {
		return "", err
	}

	fields := make([]interface{}, len(cols))
	for index := range fields {
		fields[index] = new(interface{})
	}

	data_list := []map[string]interface{}{}

	for rows.Next() {

		err := rows.Scan(fields...)
		if err != nil {
			return "", err
		}

		data_map := map[string]interface{}{}
		for index, value := range fields {
			value := value.(*interface{})
			switch data := (*value).(type) {
			case nil:
			case int64:
				data_map[cols[index]] = int(data)
			default:
				data_map[cols[index]] = data
			}
		}
		data_list = append(data_list, data_map)
	}

	final_Data, err := json.Marshal(data_list)
	if err != nil {
		return "", err
	}
	return string(final_Data), nil
}
