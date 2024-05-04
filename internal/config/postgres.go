package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

//  migrate -database "postgres://postgres:postgres@localhost:5432/cats-sosial?sslmode=disable" -path db/migrations up
//  migrate -database "postgres://postgres:postgres@localhost:5432/cats-sosial?sslmode=disable" -path db/migrations down

func NewDB(cfg *Global) *pgxpool.Pool {
	DATABASE_URL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)

	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	fmt.Println("Connected to the database!!")

	return connPool
}
