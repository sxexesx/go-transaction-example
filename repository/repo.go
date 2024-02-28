package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/sxexesx/go-transaction-example/postgre"
)

type Repo struct {
	db *postgre.DB
}

func NewRepo(db *postgre.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) WithTx(ctx context.Context, f func(context.Context, *sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "db begin tx")
	}

	if err := f(ctx, tx); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("db rollback:%w: %w", err2, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "tx commit")
	}
	return nil
}
