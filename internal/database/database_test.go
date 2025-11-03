package database

import "testing"


func TestConnect(t *testing.T) {
	db := NewDatabase("purrfact", "libsql", true)

	conn, err := db.connect()
	if err != nil {
		conn.Close()
		t.Fatalf("Failed to connect to database: %v", err)
	}

	if err = conn.Ping(); err != nil {
		conn.Close()
		t.Fatalf("Failed to ping database: %v", err)
	}
	defer conn.Close()
}
