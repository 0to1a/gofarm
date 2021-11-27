package model

import (
	"database/sql"
	"framework/framework/database"
	"framework/framework/utils"
	"github.com/doug-martin/goqu/v9"
)

func ExampleGetApikey(apikey string) (id *int, username string) {
	sqlQuery, _, _ := database.Dialect.
		Select("id", "username").
		From("access_list").
		Where(goqu.Ex{"api_key": apikey}).
		ToSQL()
	row := database.Database.QueryRow(sqlQuery)

	userId := 0
	err := row.Scan(&userId, &username)
	if err == sql.ErrNoRows {
		return nil, ""
	} else if err != nil {
		utils.LogError("ExampleGetApikey: SELECT " + err.Error())
		return nil, ""
	}

	return &userId, username
}

func ExampleGetUsername(username string) (id *int) {
	sqlQuery, _, _ := database.Dialect.
		Select("id").
		From("access_list").
		Where(goqu.Ex{"username": username}).
		ToSQL()
	row := database.Database.QueryRow(sqlQuery)

	userId := 0
	err := row.Scan(&userId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		utils.LogError("ExampleGetUsername: SELECT " + err.Error())
		return nil
	}

	return &userId
}
