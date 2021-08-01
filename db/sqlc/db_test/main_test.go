package db

import (
	"database/sql"
	"exercise-go-simplebank/db/sqlc/db"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

const (
	dbDriver  = "postgres"
	dbSourcec = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSourcec)
	if err != nil {
		return
	}
	testQueries = db.New(conn)

	os.Exit(m.Run())
}
