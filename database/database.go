package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"path/filepath"
)

var DB *leveldb.DB

func Init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Error reading home directory.")
	}

	path := filepath.Join(configDir, "bjcli")

	if err := initDb(path); err != nil {
		log.Fatal("Error opening database.")
	}
}

func Shutdown() {
	if err := closeDb(); err != nil {
		log.Fatal("Error closing database.")
	}
}

func initDb(path string) error {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func closeDb() error {
	return DB.Close()
}
