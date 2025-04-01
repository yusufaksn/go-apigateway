package db

import (
	"context"
	"fmt"
	"log"

	//"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// PostgreSQL bağlantısını başlatan fonksiyon
func ConnectDB() {
	//dsn := os.Getenv("DATABASE_URL")
	dsn := "postgres://databasepostgresql_user:databasepostgresql_password@localhost:5432/databasepostgresql?sslmode=disable"
	if dsn == "" {
		log.Fatal("DATABASE_URL ortam değişkeni ayarlanmamış.")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}
	fmt.Println("✅ Veritabanı bağlantısı başarılı!")
}
