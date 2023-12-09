package repository

import (
	"context"
	"fms/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (r *FMSRepository) AddFile(ctx context.Context, file domain.FullFileInfo, parts []domain.FilePart) error {
	//tx, err := r.db.BeginTx(ctx, nil)
	return nil
}

func (r *FMSRepository) AddFileTx(ctx context.Context, tx *sqlx.Tx, file domain.FullFileInfo, parts []domain.FilePart) error {
	var fileId int64
	err := tx.QueryRowContext(ctx, insertFileQuery, file.Filename, file.Size).Scan(&fileId)
	if err != nil {
		return errors.Wrap(err, "insert file")
	}

	builder := squirrel.Insert("fileparts").Columns("part_no", "file", "storage", "size", "path_in_storage")
	for _, p := range parts {
		builder = builder.Values(p.PartId, fileId, p.Storage.Id, p.Size, p.Path)
	}

	query, args, err := builder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "build query")
	}

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return errors.Wrap(err, "insert fileparts")
	}

	return nil
}

func (r *FMSRepository) GetFileParts(ctx context.Context, filename string) ([]domain.FilePart, error) {
	rows, err := r.db.QueryContext(ctx, getFilePartsQuery, filename)
	if err != nil {
		return nil, errors.Wrap(err, "query db")
	}
	defer rows.Close()

	var res []domain.FilePart
	for rows.Next() {
		var p domain.FilePart
		err := rows.Scan(
			&p.PartId,
			&p.Storage.Address,
			&p.Size,
			&p.Path,
			&p.IsStored,
		)
		if err != nil {
			return nil, errors.Wrap(err, "scan value")
		}

		res = append(res, p)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows err")
	}

	return res, nil
}

func (r *FMSRepository) DeleteFile(ctx context.Context, filename string) error {
	//TODO implement me
	panic("implement me")
}

func (r *FMSRepository) DeleteFileTx(ctx context.Context, tx *sqlx.Tx, filename string) error {
	if _, err := tx.ExecContext(ctx, deleteFilePartsQuery, filename); err != nil {
		return errors.Wrap(err, "delete file parts")
	}

	if _, err := tx.ExecContext(ctx, deleteFileQuery, filename); err != nil {
		return errors.Wrap(err, "delete file")
	}

	return nil
}

func (r *FMSRepository) MarkPartStored(ctx context.Context, filename string, partId int) error {
	res, err := r.db.ExecContext(ctx, markPartStoredQuery, filename, partId)
	if err != nil {
		return errors.Wrap(err, "query db")
	}
	if rows, err := res.RowsAffected(); err != nil || rows != 1 {
		return errors.Wrap(err, "cannot check affected rows")
	}

	return nil
}
