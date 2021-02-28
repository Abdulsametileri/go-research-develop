package database

import "errors"

var (
	ErrDbConnection = errors.New("VERİTABANI BAĞLANTISI KURULAMADI")
	ErrPing         = errors.New("VERİTABANINA PİNG ATILAMADI")
)

type Database interface {
	setup()
}

func SetupDB(db Database) {
	db.setup()
}
