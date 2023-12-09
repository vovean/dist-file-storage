package repository

import (
	"context"
	"fms/internal/domain"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (r *FMSRepository) GetStorages(ctx context.Context) ([]domain.Storage, error) {
	rows, err := r.db.QueryContext(ctx, selectStoragesQuery)
	if err != nil {
		return nil, errors.Wrap(err, "query db")
	}

	var res []domain.Storage
	for rows.Next() {
		var s domain.Storage
		if err := rows.Scan(&s.Id, &s.Address, &s.SpaceAvailableBytes); err != nil {
			return nil, errors.Wrap(err, "scan row")
		}

		res = append(res, s)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows error")
	}

	return res, nil
}

func (r *FMSRepository) GetStoragesTx(ctx context.Context, tx *sqlx.Tx) ([]domain.Storage, error) {
	rows, err := tx.QueryContext(ctx, selectStoragesQuery)
	if err != nil {
		return nil, errors.Wrap(err, "query db")
	}
	defer rows.Close()

	var res []domain.Storage
	for rows.Next() {
		var s domain.Storage
		if err := rows.Scan(&s.Id, &s.Address, &s.SpaceAvailableBytes); err != nil {
			return nil, errors.Wrap(err, "scan row")
		}

		res = append(res, s)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows error")
	}

	return res, nil
}

func (r *FMSRepository) AddStorage(ctx context.Context, storage domain.Storage) (domain.Storage, error) {
	var id int
	err := r.db.QueryRowContext(ctx, insertStorageQuery, storage.Address, storage.SpaceAvailableBytes).Scan(&id)
	if err != nil {
		return domain.Storage{}, errors.Wrap(err, "insert and scan")
	}

	storage.Id = id
	return storage, nil
}
