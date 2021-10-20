package model

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

var (
	Database *sqlx.DB
	Dialect  goqu.DialectWrapper
)
