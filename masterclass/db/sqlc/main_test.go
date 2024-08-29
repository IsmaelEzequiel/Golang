package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgres://root:123@localhost:5432/simple_bank?sslmode=disable"
)

var storeDB Store

func TestMain(m *testing.M) {
	poolConn, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal(err)
	}

	defer poolConn.Close()

	storeDB = *NewStore(poolConn)

	os.Exit(m.Run())
}
