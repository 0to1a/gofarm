package framework

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bmizerany/assert"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"testing"
	"testing/fstest"
)

func TestConnectMysql(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	t.Run("Connect to mock", func(t *testing.T) {
		dbMysql.connectSuccess(db)
		assert.NotEqual(t, nil, dbMysql.client)
	})
	t.Run("Connect mysql", func(t *testing.T) {
		dbMysql.Host = ":33306"
		dbMysql.Username = ""
		dbMysql.Password = ""
		dbMysql.Database = ""
		assert.Panic(t, nil, func() {
			dbMysql.Connect()
		})
	})
}

func TestCheckClientMysql(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	t.Run("Error not connect", func(t *testing.T) {
		dbMysql.client = nil
		DatabaseMysql = nil
		assert.Equal(t, (*sql.DB)(nil), dbMysql.CheckClient())
	})
	t.Run("With global connection", func(t *testing.T) {
		dbMysql.client = nil
		DatabaseMysql = db
		assert.NotEqual(t, (*sql.DB)(nil), dbMysql.CheckClient())
	})
	t.Run("With local connection", func(t *testing.T) {
		dbMysql.client = db
		DatabaseMysql = nil
		assert.NotEqual(t, (*sql.DB)(nil), dbMysql.CheckClient())
	})
}

func TestMigrateDatabase(t *testing.T) {
	fsSample := fstest.MapFS{
		"migration/1_test.up.sql": {
			Data: []byte("SELECT 1;"),
		},
		"migration/1_test.down.sql": {
			Data: []byte("SELECT 1;"),
		},
	}
	d, _ := iofs.New(fsSample, "migration")

	t.Run("Database connection fail", func(t *testing.T) {
		dbMysql.Host = "localhost:33306"
		dbMysql.Username = "test"
		dbMysql.Password = "test"
		dbMysql.Database = "test"
		assert.Panic(t, "dial tcp [::1]:33306: connect: connection refused", func() {
			dbMysql.Connect()
			dbMysql.MigrateDatabase(d)
		})
	})
	t.Run("Migrate success", func(t *testing.T) {
		assert.Panic(t, nil, func() {
			dbMysql.migrateHandling(nil)
		})
	})
	t.Run("Migrate no change", func(t *testing.T) {
		assert.Panic(t, nil, func() {
			dbMysql.migrateHandling(migrate.ErrNoChange)
		})
	})
	t.Run("Migrate Fail", func(t *testing.T) {
		assert.Panic(t, "database locked", func() {
			dbMysql.migrateHandling(migrate.ErrLocked)
		})
	})
}

func TestTableCheck(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := mock.NewRows([]string{"Tables_in"}).
		AddRow("one")

	mock.ExpectQuery("SHOW TABLES LIKE 'test'").WillReturnRows(rows)

	t.Run("Failure", func(t *testing.T) {
		dbMysql.client = nil
		DatabaseMysql = db
		dbMysql.TableCheck("test")
		assert.Equal(t, false, dbMysql.TableCheck("test"))
	})
}
