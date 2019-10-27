package storage

import (
	"boats/clients/nausys"
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type YachtInfo struct {
	Name          string    `db:"name"`
	BuilderName   string    `db:"builder_name"`
	ModelName     string    `db:"model_name"`
	CharterName   string    `db:"charter_name"`
	AvailableNow  bool      `db:"available_now"`
	AvailableFrom time.Time `db:"available_from"`
	AvailableTo   time.Time `db:"available_to"`
}

type Storage interface {
	ClearAll(ctx context.Context) error
	GetLastUpdateInfo(ctx context.Context) (time.Time, error)
	InsertBuilders(ctx context.Context, builders []nausys.Builder) error
	InsertModels(ctx context.Context, models []nausys.Model) error
	InsertCharters(ctx context.Context, charters []nausys.Charter) error
	InsertYachts(ctx context.Context, yachts []*nausys.Yacht) error
	InsertUpdateInfo(ctx context.Context) error
	//FindYachts(ctx context.Context, builderNamePrefix, modelNamePrefix string, limit, offset int) (yachts []YachtInfo, total int64, err error)
}

type StorageImpl struct {
	db *sqlx.DB
}

func NewStorageImpl(db *sqlx.DB) *StorageImpl {
	return &StorageImpl{db}
}

func (s *StorageImpl) ClearAll(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `
TRUNCATE models, charters, builders, yachts, update_info;`)
	return err
}

func (s *StorageImpl) InsertBuilders(ctx context.Context, builders []nausys.Builder) error {
	for _, builder := range builders {
		_, err := s.db.ExecContext(ctx, "INSERT INTO builders(id, name) VALUES ($1, $2)", builder.Id, builder.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StorageImpl) InsertModels(ctx context.Context, models []nausys.Model) error {
	for _, model := range models {
		_, err := s.db.ExecContext(ctx, "INSERT INTO models(id, name, builder_id) VALUES ($1, $2, $3)",
			model.Id, model.Name, model.BuilderId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StorageImpl) InsertCharters(ctx context.Context, charters []nausys.Charter) error {
	for _, charter := range charters {
		_, err := s.db.ExecContext(ctx, "INSERT INTO charters(id, name) VALUES ($1, $2)", charter.Id, charter.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StorageImpl) InsertYachts(ctx context.Context, yachts []*nausys.Yacht) error {
	for _, yacht := range yachts {
		_, err := s.db.ExecContext(ctx, `
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

func (s *StorageImpl) InsertUpdateInfo(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO update_info(moment) VALUES (NOW())`)
	return err
}

//
//func (s *StorageImpl) FindYachts(ctx context.Context, builderNamePrefix, modelNamePrefix string, limit, offset int) (yachts []YachtInfo, total int64, err error) {
//	err = s.db.SelectContext(ctx, &yachts, `
//SELECT
//`)
//	return
//}
