package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbSource = "postgres://root:123@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgx.Conn

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgx.Connect(context.Background(), dbSource)

	if err != nil {
		log.Fatal(err)
	}

	defer testDB.Close(context.Background())

	testQueries = New(testDB)

	os.Exit(m.Run())
}
