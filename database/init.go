package database

import "errors"

var (
	ErrDbConnection = errors.New("BAĞLANTI KURULAMADI")
	ErrPing         = errors.New("PİNG ATILAMADI")
)

type Database interface {
	setup()
}

func SetupDB(db Database) {
	db.setup()
}
