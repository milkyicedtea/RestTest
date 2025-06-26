package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

var (
	pool *pgxpool.Pool = nil
	once sync.Once
)

func InitPgxPool(ctx context.Context, cfg *Config) error {
	var initErr error
	once.Do(func() {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%d&pool_min_conns=%d",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
			4, // max connections, can also be from config
			1, // min connections, can also be from config
		)

		//log.Printf("PG connStr: %s", connStr)

		pgxConfig, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			initErr = fmt.Errorf("failed to parse pgxpool config: %w", err)
			return
		}

		p, err := pgxpool.NewWithConfig(ctx, pgxConfig)
		if err != nil {
			initErr = fmt.Errorf("failed to create pgxpool: %w", err)
			return
		}

		// ping database to ensure connectivity
		if err := p.Ping(ctx); err != nil {
			p.Close() // close pool if ping fails
			initErr = fmt.Errorf("failed to ping database with pgxpool: %w", err)
			return
		}
		pool = p // assign to the global variable
		log.Println("Successfully connected to PostgreSQL.")
	})

	if initErr != nil {
		return initErr
	}
	if pool == nil && initErr == nil {
		return fmt.Errorf("pgxpool initialization failed silently")
	}
	return nil
}

func GetPool() *pgxpool.Pool {
	if pool == nil {
		log.Panic("pgxpool has not been initialized. Call database.InitPgxPool first.")
	}
	return pool
}

func ClosePgxPool() {
	if pool != nil {
		pool.Close()
		log.Println("PostgreSQL connection pool closed.")
	}
}

func GetPgConn() (*pgxpool.Conn, error) {
	conn, err := GetPool().Acquire(context.Background())
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return conn, nil
}
