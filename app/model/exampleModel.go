package model

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"log"
)

func ExampleGetApikey(apikey string) (id *int, username string) {
	sqlQuery, _, _ := Dialect.
		Select("id", "username").
		From("access_list").
		Where(goqu.Ex{"api_key": apikey}).
		ToSQL()
	row := Database.QueryRow(sqlQuery)

	userId := 0
	err := row.Scan(&userId, &username)
	if err == sql.ErrNoRows {
		return nil, ""
	} else if err != nil {
		log.Println("ExampleGetApikey: SELECT", err)
		return nil, ""
	}

	return &userId, username
}

func ExampleGetUsername(username string) (id *int) {
	sqlQuery, _, _ := Dialect.
		Select("id").
		From("access_list").
		Where(goqu.Ex{"username": username}).
		ToSQL()
	row := Database.QueryRow(sqlQuery)

	userId := 0
	err := row.Scan(&userId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		log.Println("ExampleGetUsername: SELECT", err)
		return nil
	}

	return &userId
}
