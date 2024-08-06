package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secretPass123@localhost:5432/go-bank?sslmode=disable"
)

var testQueries *Queries


func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}