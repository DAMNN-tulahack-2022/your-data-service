package tools

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type DBConnectionPool struct {
	log *YourLogger

	Db *sqlx.DB
}

const (
	_postgresGenericDriver = "postgres"
)



func BeginDbInstance(config *GenericEnvAppConfig, log *YourLogger) (db *DBConnectionPool, err error) {
	db = new(DBConnectionPool)
	db.log = log
	db.Db, err = sqlx.Connect(_postgresGenericDriver, config.DbSoruceDSN)
	if err != nil {
		return 
	}

	db.Db.SetMaxIdleConns(config.MaxIdleConns)
	db.Db.SetMaxOpenConns(config.MaxOpenConns)
	db.Db.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second)
	db.Db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)

	return
}
