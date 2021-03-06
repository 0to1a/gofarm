package migration

import (
	"framework/framework/database"
	"github.com/doug-martin/goqu/v9"
	"log"
)

func migrate211020example() {
	ret := database.MysqlSeedCheck(database.Database)
	if ret != database.ErrSeedNoRows {
		return
	}

	_, err := database.Database.Query(`
	CREATE TABLE access_list (
		id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
		username VARCHAR(50) NOT NULL DEFAULT '' COLLATE 'utf8mb4_general_ci',
		api_key VARCHAR(50) NOT NULL DEFAULT '' COLLATE 'utf8mb4_general_ci',
		
		PRIMARY KEY (id) USING BTREE
	) COLLATE='utf8mb4_general_ci' ENGINE=InnoDB;
	`)
	if err != nil {
		log.Fatalln(ErrorMigration, err.Error())
	}

	sqlQuery, _, _ := database.Dialect.
		Insert("access_list").
		Cols("username", "api_key").
		Vals(
			goqu.Vals{"example1", "123456789"},
			goqu.Vals{"example2", "987654321"},
		).
		ToSQL()
	_, err = database.Database.Exec(sqlQuery)
	if err != nil {
		log.Fatalln(ErrorMigration, err.Error())
	}

	database.MysqlInsertMigration(database.Database)
}
