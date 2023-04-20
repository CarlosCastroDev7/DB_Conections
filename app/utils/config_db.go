package utils

import "github.com/carlos/SQLserver_connection/app/models"

func Config_mssql(company string) (config *models.Config_model_mssql) {

	config = &models.Config_model_mssql{
		Server:   "Server",
		Port:     1433,
		User:     "User",
		Password: "Password",
		Database: company,
	}
	return

}
