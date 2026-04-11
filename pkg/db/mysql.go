package db

import (
	"github.com/jmoiron/sqlx"
	"time"
)

const MysqlConnectionCnt = 50

func NewMySQL(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(MysqlConnectionCnt)
	db.SetMaxIdleConns(MysqlConnectionCnt / 2)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}
