package main

import (
	"fmt"

	mssql "github.com/carlos/SQLserver_connection/app/mssql"
)

func main() {
	data, err := mssql.Runquery_getjson("SELECT * FROM [use].usuarios FOR JSON PATH", "ENFETER")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
