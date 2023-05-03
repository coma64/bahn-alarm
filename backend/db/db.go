package db

import (
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db = sqlx.MustConnect("postgres", config.Conf.Db.Dsn)
