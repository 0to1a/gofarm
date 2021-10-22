package model

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

var (
	Database *sqlx.DB
	Dialect  goqu.DialectWrapper

	RedisDB *redis.Client
)
