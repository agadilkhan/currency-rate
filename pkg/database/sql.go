package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strings"
)

type SQLX struct {
	Client *sqlx.DB
}

// postgres://username:password@localhost:5432/db_name?sslmode=disable&search_path=public

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

func (s *SQLX) Close() error {
	return s.Client.Close()
}
