package infra

import (
	"database/sql"
	"github.com/machine-brest/k8s-helloworld/internal/ifs"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Database struct{}

func NewDatabase() ifs.Database {
	return &Database{}
}

func (_ *Database) GetConnection() (*bun.DB, error) {
	dsn := "postgres://dbuser:dbpass@postgres:5432/sample?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	conn := bun.NewDB(sqlDb, pgdialect.New())

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
