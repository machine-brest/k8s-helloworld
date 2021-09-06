package ifs

import "github.com/uptrace/bun"

type Database interface {
	GetConnection() (*bun.DB, error)
}
