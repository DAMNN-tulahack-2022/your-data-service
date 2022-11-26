package tools

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

const (
	_postgresGenericDriver = "postgres"
)

func BeginDbInstance(config *GenericEnvAppConfig, log *YourLogger) (err error) {
	log.log.Print(config.DbSoruceDSN)
	db, err := sqlx.Connect("postgres", config.DbSoruceDSN)
	if err != nil {
		return
	}

	DB = db 
	log.log.Print("db is init")
	log.log.Print("db is nil - [%v]", db == nil)

	return err
}

func Close() {
	if err := DB.Close(); err != nil {
		return
	}
}
