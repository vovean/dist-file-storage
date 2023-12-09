package repository

import (
	"context"
	"database/sql"
	goerrors "errors"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type FMSRepository struct {
	db *sqlx.DB
}

func NewFMSRepository(db *sqlx.DB) *FMSRepository {
	return &FMSRepository{db: db}
}

func (r *FMSRepository) Serializable(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return errors.Wrap(err, "begin tx")
	}

	err = f(ctx, tx)
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = goerrors.Join(err, rErr)
		}
		return errors.Wrap(err, "error during transaction")
	}

	return errors.Wrap(tx.Commit(), "commit tx")
}
