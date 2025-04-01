package db

import (
	"context"
	"fmt"
	"log"

	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {

	var pgConnection = os.Getenv("PG_CONN")
	if pgConnection == "" {
		log.Fatal("connection error")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), pgConnection)
	if err != nil {
		log.Fatal("Do not connection", err)
	}
	fmt.Println("✅ Connected")
}
