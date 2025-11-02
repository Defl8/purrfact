package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	// Turso driver
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Database struct {
	FullPath string
	Name     string
	Driver   string
	IsRemote bool
	Token    string
}

func NewDatabase(name, driver string, isRemote bool) *Database {

	if isRemote {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}

		url := os.Getenv("TURSO_DATABASE_URL")
		token := os.Getenv("TURSO_AUTH_TOKEN")

		fullpath := fmt.Sprintf("%s?auth_token=%s", url, token)

		return &Database{
			Driver:   driver,
			Name:     name,
			FullPath: fullpath,
			IsRemote: isRemote,
			Token:    token,
		}
	}

	return &Database{
		Driver:   driver,
		Name:     name,
		IsRemote: isRemote,
	}
}

// Only really use for local db
func (database Database) SetDBPath(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Database at path '%s' does not exits. ERROR: %w", path, err)
	}

	return nil
}

func (database Database) connect() (*sql.DB, error) {
	dbConn, err := sql.Open(database.Driver, database.FullPath)
	if err != nil {
		return nil, fmt.Errorf("Connection to '%s' database could not be established. ERROR %w", database.Name, err)
	}

	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, fmt.Errorf("Could not ping '%s' database. ERROR %w", database.Name, err)
	}

	return dbConn, nil
}

func (db Database) ExecuteSelect(query string) (*sql.Rows, error) {
	conn, err := db.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		rows.Close()
		return nil, err
	}

	return rows, nil
}

func (db Database) ExecuteDUI(query string) (int64, error) {
	conn, err := db.connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(query)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil
}
