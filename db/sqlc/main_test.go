package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	databasedriver = "postgres"
	datasource     = "postgresql://root:secret@localhost:5432/contoso_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(databasedriver, datasource)
	if err != nil {
		log.Fatal("error connecting database", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
