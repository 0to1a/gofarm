package database

import (
	"crypto/sha1"
	"encoding/base64"
)

const (
	SeedOK        = 200
	ErrSeedNoRows = -1
	ErrMigration  = -2
)

const (
	okMigration1    = "No Migration Unit: "
	errorMigration1 = "Err Migration #S0001: no indexing function"
	errorMigration2 = "Err Migration #S0002: database not init"
	errorMigration3 = "Err Migration #S0003: database can't create table"
	errorMigration4 = "Err Migration #S0004: problem to insert data"
)

func SeedName(functionName string) string {
	hashes := sha1.New()
	hashes.Write([]byte(functionName))
	seedName := base64.URLEncoding.EncodeToString(hashes.Sum(nil))

	return seedName
}
