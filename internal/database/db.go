package database

import (
	"database/sql"
	"fmt"
	"os"
	"errors"
)



type Database struct {
	FullPath string
	Name string
	Driver string
}

func (database Database) SetDBPath(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Database at path '%s' does not exits. ERROR: %w", path, err)
	}
}


func (database Database) connect() (*sql.DB, error) {
	dbConn, err := sql.Open(database.Driver, database.Name)
	if err != nil {
		return nil, fmt.Errorf("Connection to '%s' database could not be established. ERROR %w", database.Name, err)
	}
	
	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, fmt.Errorf("Could not ping '%' database", database.Name, err)
	}
}

func (db Database) ExecuteQuery() {

}
