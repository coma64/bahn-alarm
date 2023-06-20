package db

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strings"
)

var Db = sqlx.MustConnect("postgres", CreateConnectionStr())
var Sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func CreateConnectionStr() string {
	c := config.Conf.Db
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.User, strings.TrimSpace(c.Password), c.Host, c.DbName)
}
