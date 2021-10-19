package database

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"log"
	"runtime"
)

var (
	dialectMysql goqu.DialectWrapper
)

func MysqlTableCheck(database *sqlx.DB, table string) bool {
	_, tableCheck := database.Query("select * from " + table + ";")

	if tableCheck == nil {
		return true
	} else {
		return false
	}
}

func MysqlSeedCheck(database *sqlx.DB) int64 {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	seedName := ""
	if ok && details != nil {
		seedName = SeedName(details.Name())
	} else {
		log.Fatalln(errorMigration1)
	}

	if database == nil {
		log.Fatalln(errorMigration2)
	}

	if !MysqlTableCheck(database, "migration") {
		_, err := database.Query(`
		CREATE TABLE migration (
			id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
			seedName VARCHAR(100) NOT NULL DEFAULT '',
			PRIMARY KEY (id)
		) COLLATE='utf8mb4_general_ci';
		`)

		if err != nil {
			log.Fatalln(errorMigration3)
		}
	}

	dialectMysql = goqu.Dialect("mysql")
	sqlQuery, _, _ := dialectMysql.
		Select("id").
		From("migration").
		Where(goqu.Ex{"seedName": seedName}).
		ToSQL()
	row := database.QueryRow(sqlQuery)
	trxId := ""

	err := row.Scan(&trxId)
	if err == sql.ErrNoRows {
		return ErrSeedNoRows
	}

	return SeedOK
}

func MysqlInsertMigration(database *sqlx.DB) int64 {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	seedName := ""
	if ok && details != nil {
		seedName = SeedName(details.Name())
	} else {
		log.Fatalln(errorMigration1)
	}

	dialectMysql = goqu.Dialect("mysql")
	sqlQuery, _, _ := dialectMysql.
		Insert("migration").
		Cols("seedName").
		Vals(goqu.Vals{seedName}).
		ToSQL()
	res, err := database.Exec(sqlQuery)

	if err != nil {
		log.Println(errorMigration4)
		return ErrMigration
	} else {
		targetId, err := res.LastInsertId()
		if err != nil {
			log.Println(errorMigration4)
			return ErrMigration
		} else {
			return targetId
		}
	}
}

func MysqlConnect(databaseUsername string, databasePassword string, databaseHost string, databaseName string) *sqlx.DB {
	database, err := sqlx.Connect("mysql", databaseUsername+":"+databasePassword+"@tcp("+databaseHost+")/"+databaseName+"?multiStatements=true&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	return database
}
