package mssql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/carlos/SQLserver_connection/app/utils"
	_ "github.com/denisenkom/go-mssqldb"
)

func connection_mssql(company string) (db *sql.DB, err error) {

	config := utils.Config_mssql(company)

	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s",
		config.Server,
		config.User,
		config.Password,
		config.Port,
		config.Database)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil

}
