package db

import "golang.org/x/tools/go/cfg"

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error)
