package storage

import (
	"boats/clients/nausys"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type YachtInfo struct {
	Name          string `db:"name"`
	BuilderName   string `db:"builder_name"`
	ModelName     string `db:"model_name"`
	CharterName   string `db:"charter_name"`
	AvailableNow  bool   `db:"available_now"`
	AvailableFrom string `db:"available_from"`
	AvailableTo   string `db:"available_to"`
}

type Storage interface {
	WithTransaction(ctx context.Context, f func(tx *sqlx.Tx) error) (err error)

	ClearAll(ctx context.Context, querier sqlx.ExecerContext) error
	GetLastUpdateInfo(ctx context.Context) (time.Time, error)
	InsertBuilders(ctx context.Context, querier sqlx.ExecerContext, builders []nausys.Builder) error
	InsertModels(ctx context.Context, querier sqlx.ExecerContext, models []nausys.Model) error
	InsertCharters(ctx context.Context, querier sqlx.ExecerContext, charters []nausys.Charter) error
	InsertYachts(ctx context.Context, querier sqlx.ExecerContext, yachts []*nausys.Yacht) error
	InsertUpdateInfo(ctx context.Context, querier sqlx.ExecerContext) error

	FindYachts(ctx context.Context, builderNamePrefix, modelNamePrefix string, limit, offset int) (yachts []YachtInfo, total int64, err error)
	FindBuildersByPrefix(ctx context.Context, prefix string, limit int) ([]string, error)
	FindModelsByPrefix(ctx context.Context, prefix string, limit int) ([]string, error)
}

type StorageImpl struct {
	db *sqlx.DB
}

func NewStorageImpl(db *sqlx.DB) *StorageImpl {
	return &StorageImpl{db}
}

func (s *StorageImpl) WithTransaction(ctx context.Context, f func(tx *sqlx.Tx) error) (err error) {
	var trx *sqlx.Tx
	trx, err = s.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recover after panic: %v", r)
		}
		for {
			if err != nil {
				if errRollback := trx.Rollback(); errRollback != nil {
					err = errors.Wrapf(errRollback, fmt.Sprintf("rollback after error: %v", err))
				}
				return
			}
			if err := trx.Commit(); err != nil {
				err = errors.Wrapf(err, "commit transaction")
			} else {
				return
			}
		}
	}()
	err = f(trx)
	return err
}

func (s *StorageImpl) ClearAll(ctx context.Context, querier sqlx.ExecerContext) error {
	_, err := querier.ExecContext(ctx, `
TRUNCATE models, charters, builders, yachts, update_info;`)
	return err
}

func (s *StorageImpl) InsertBuilders(ctx context.Context, querier sqlx.ExecerContext, builders []nausys.Builder) error {
	for _, builder := range builders {
		_, err := querier.ExecContext(ctx, "INSERT INTO builders(id, name) VALUES ($1, $2)", builder.Id, builder.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StorageImpl) InsertModels(ctx context.Context, querier sqlx.ExecerContext, models []nausys.Model) error {
	for _, model := range models {
		_, err := querier.ExecContext(ctx, "INSERT INTO models(id, name, builder_id) VALUES ($1, $2, $3)",
			model.Id, model.Name, model.BuilderId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StorageImpl) InsertCharters(ctx context.Context, querier sqlx.ExecerContext, charters []nausys.Charter) error {
	for _, charter := range charters {
		_, err := querier.ExecContext(ctx, "INSERT INTO charters(id, name) VALUES ($1, $2)", charter.Id, charter.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StorageImpl) InsertYachts(ctx context.Context, querier sqlx.ExecerContext, yachts []*nausys.Yacht) error {
	for _, yacht := range yachts {
		_, err := querier.ExecContext(ctx, `
INSERT INTO yachts(id, name, model_id, charter_id, available_from, available_to) VALUES ($1, $2, $3, $4, $5, $6)`,
			yacht.Id, yacht.Name, yacht.ModelId, yacht.CharterId, yacht.AvailableFrom, yacht.AvailableTo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StorageImpl) GetLastUpdateInfo(ctx context.Context) (time.Time, error) {
	var updatedAt time.Time
	err := s.db.GetContext(ctx, &updatedAt, `SELECT moment FROM update_info ORDER BY moment DESC LIMIT 1;`)
	return updatedAt, err
}

func (s *StorageImpl) InsertUpdateInfo(ctx context.Context, querier sqlx.ExecerContext) error {
	_, err := querier.ExecContext(ctx, `INSERT INTO update_info(moment) VALUES (NOW())`)
	return err
}

func (s *StorageImpl) FindYachts(ctx context.Context, builderNamePrefix, modelNamePrefix string, limit, offset int) (yachts []YachtInfo, total int64, err error) {
	queryCommon := `
SELECT %s
FROM yachts Y JOIN charters C on Y.charter_id = C.id JOIN models M on Y.model_id = M.id JOIN builders B on M.builder_id = B.id
WHERE B.name LIKE $1 AND M.name LIKE $2`
	err = s.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		if err = tx.GetContext(ctx, &total, fmt.Sprintf(queryCommon, "COUNT(*)"), builderNamePrefix+"%", modelNamePrefix+"%"); err != nil {
			return errors.Wrap(err, "select total number of yachts")
		}

		query := fmt.Sprintf(queryCommon, "Y.name AS name, B.name AS builder_name, M.name AS model_name, C.name AS charter_name, COALESCE(Y.available_from < NOW(), TRUE) AS available_now, COALESCE(to_char(Y.available_from, 'DD Mon YYYY'), '') AS available_from, COALESCE(to_char(Y.available_to, 'DD Mon YYYY'), '') AS available_to") +
			"OFFSET $3 LIMIT $4;"
		if err = tx.SelectContext(ctx, &yachts, query, builderNamePrefix+"%", modelNamePrefix+"%", offset, limit); err != nil {
			return errors.Wrap(err, "select yachts")
		}

		return nil
	})

	return
}

func (s *StorageImpl) FindBuildersByPrefix(ctx context.Context, prefix string, limit int) ([]string, error) {
	var res []string
	if err := s.db.SelectContext(ctx, &res, `SELECT name FROM builders WHERE name LIKE $1 LIMIT $2;`, prefix+"%", limit); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StorageImpl) FindModelsByPrefix(ctx context.Context, prefix string, limit int) ([]string, error) {
	var res []string
	if err := s.db.SelectContext(ctx, &res, `SELECT name FROM models WHERE name LIKE $1 LIMIT $2;`, prefix+"%", limit); err != nil {
		return nil, err
	}
	return res, nil
}
