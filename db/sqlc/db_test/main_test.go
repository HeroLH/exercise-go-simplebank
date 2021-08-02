package db

import (
	"database/sql"
	"exercise-go-simplebank/db/sqlc/db"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver  = "postgres"
	dbSourcec = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSourcec)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
		return
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
