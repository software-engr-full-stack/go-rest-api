package sqlite_test

import (
	"testing"

	"go-rest-api/sqlite"
	"go-rest-api/lib"
)

// Ensure the test database can open & close.
func TestDB(t *testing.T) {
	db := MustOpenDB(t)
	MustCloseDB(t, db)
}

// MustOpenDB returns a new, open DB. Fatal on error.
func MustOpenDB(tb testing.TB) *sqlite.DB {
	tb.Helper()

	cfg, err := lib.TestConfig()
	if err != nil {
		tb.Fatal(err)
	}

	// Write to an in-memory database by default.
	// If the -dump flag is set, generate a temp file for the database.
	db := sqlite.NewDB(cfg)
	if err := db.Open(); err != nil {
		tb.Fatal(err)
	}

	return db
}

// MustCloseDB closes the DB. Fatal on error.
func MustCloseDB(tb testing.TB, db *sqlite.DB) {
	tb.Helper()
	if err := db.Close(); err != nil {
		tb.Fatal(err)
	}
}
