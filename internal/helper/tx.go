package helper

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func CommitOrRollback(tx pgx.Tx) {
	err := recover()
	if err != nil {
		if err := tx.Rollback(context.Background()); err != nil {
			log.Printf("failed to rollback: %v", err.Error())
			return
		}
	} else {
		if err := tx.Commit(context.Background()); err != nil {
			log.Printf("failed to commit: %v", err.Error())
			return
		}
	}
}
