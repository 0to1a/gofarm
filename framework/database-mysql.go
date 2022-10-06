package framework

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// TODO: re-make with SQLBoiler

var (
	DatabaseMysql *sqlx.DB
	DialectMysql  goqu.DialectWrapper
)

type MysqlDatabase struct {
	Host     string
	Username string
	Password string
	Database string
	Dialect  goqu.DialectWrapper
	client   *sqlx.DB
}

func (w MysqlDatabase) Connect() *sqlx.DB {
	database, err := sqlx.Connect("mysql", w.Username+":"+w.Password+"@tcp("+w.Host+")/"+w.Database+"?multiStatements=true&parseTime=true&sql_mode='ANSI_QUOTES'")
	if err != nil {
		log.Fatalln(err)
	}
	w.client = database
	w.Dialect = goqu.Dialect("mysql")
	DatabaseMysql = database
	DialectMysql = goqu.Dialect("mysql")

	return database
}

func (w MysqlDatabase) CheckClient() *sqlx.DB {
	if DatabaseMysql == nil && w.client == nil {
		return nil
	} else if w.client != nil {
		return w.client
	} else {
		return DatabaseMysql
	}
}

func (w MysqlDatabase) TableCheck(table string) bool {
	_, err := w.CheckClient().Query("SHOW TABLES LIKE '" + table + "';")

	if err == nil {
		return true
	} else {
		return false
	}
}
