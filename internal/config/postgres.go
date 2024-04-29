package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitialDB(cfg *Global) *pgx.Conn {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("successfully connected to database")

	return db
}

//  migrate -database "postgres://postgres:postgres@localhost:5432/cats-sosial?sslmode=disable" -path db/migrations up
