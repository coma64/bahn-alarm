package db

import (
	"github.com/Masterminds/squirrel"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db = sqlx.MustConnect("postgres", config.Conf.Db.Dsn)
var Sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
