package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SQLX struct {
	Client *sqlx.DB
}

func NewSQL(dataSourceName string) (*SQLX, error) {
	if !strings.Contains(dataSourceName, "://") {
		return nil, fmt.Errorf("undefined data source name " + dataSourceName)
	}
	driverName := strings.ToLower(strings.Split(dataSourceName, "://")[0])

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("db connection failed: %v", err)
	}

	return &SQLX{Client: db}, nil
}
