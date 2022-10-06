package framework

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// TODO: re-make with SQLBoiler

var (
	DatabaseMysql *sql.DB
	DialectMysql  goqu.DialectWrapper
)

type MysqlDatabase struct {
	Host     string
	Username string
	Password string
	Database string
	Dialect  goqu.DialectWrapper
	client   *sql.DB
}

func (w *MysqlDatabase) Connect() *sql.DB {
	database, err := sql.Open("mysql", w.Username+":"+w.Password+"@tcp("+w.Host+")/"+w.Database+"?multiStatements=true&parseTime=true&sql_mode='ANSI_QUOTES'")
	//database, err := sqlx.Connect("mysql", w.Username+":"+w.Password+"@tcp("+w.Host+")/"+w.Database+"?multiStatements=true&parseTime=true&sql_mode='ANSI_QUOTES'")
	if err != nil {
		log.Fatalln(err)
	}
	w.client = database
	w.Dialect = goqu.Dialect("mysql")
	DatabaseMysql = database
	DialectMysql = goqu.Dialect("mysql")

	return database
}

func (w *MysqlDatabase) CheckClient() *sql.DB {
	if DatabaseMysql == nil && w.client == nil {
		return nil
	} else if w.client != nil {
		return w.client
	} else {
		return DatabaseMysql
	}
}

func (w *MysqlDatabase) MigrateDatabase(data source.Driver) {
	url := w.Username + ":" + w.Password + "@tcp(" + w.Host + ")/" + w.Database + "?multiStatements=true&parseTime=true&sql_mode='ANSI_QUOTES'"
	m, err := migrate.NewWithSourceInstance("iofs", data, "mysql://"+url)
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Up()
	if err != nil {
		if err.Error() == okMigration1 {
			return
		}
		log.Fatalln(err)
	} else {
		log.Println(okMigration2)
	}
}

func (w *MysqlDatabase) TableCheck(table string) bool {
	_, err := w.CheckClient().Query("SHOW TABLES LIKE '" + table + "';")

	if err == nil {
		return true
	} else {
		return false
	}
}
