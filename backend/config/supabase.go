package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB holds the shared connection pool for database operations.
var DB *pgxpool.Pool

// ConnectDB initializes the global connection pool using the SUPABASE_DB_URL env var.
func ConnectDB() {
	dsn := os.Getenv("SUPABASE_DB_URL")
	if dsn == "" {
		log.Fatal("SUPABASE_DB_URL belum diset di .env")
	}

	ctx := context.Background()

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Gagal parse DB config: %v", err)
	}

	// Supabase uses PgBouncer transaction pooling; simple protocol avoids prepared statements.
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("Gagal konek database: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Gagal ping database: %v", err)
	}

	fmt.Println("Berhasil konek ke Supabase")
	DB = pool
}
