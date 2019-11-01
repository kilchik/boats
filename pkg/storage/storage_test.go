package storage

import (
	"boats/clients/nausys"
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)


type TestStorageSuite struct {
	suite.Suite

	db *sqlx.DB
	store  *StorageImpl
}

func (s *TestStorageSuite) SetupTest() {
	s.db = s.initDB()
	s.store = NewStorageImpl(s.db)
}

func (s *TestStorageSuite) initDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		s.FailNow("pass DSN in env var")
	}

	sqlDB := sqlx.MustConnect("pgx", dsn)

	err := goose.Up(sqlDB.DB, "../../migrations")
	if err != nil {
		s.FailNow("migrate", err)
	}

	return sqlDB
}

func (s *TestStorageSuite) TearDownTest() {
	s.store.ClearAll(context.Background(), s.db)
	s.db.Close()
}

func (s *TestStorageSuite) TestStorageImpl_FindYachts() {
	ctx := context.Background()
	s.store.InsertBuilders(ctx, s.db, []nausys.Builder{
		{
			Id:1,
			Name:"builder abc",
		},
		{
			Id:2,
			Name:"builder def",
		},
	})
	s.store.InsertModels(ctx, s.db, []nausys.Model{
		{
			Id:1,
			Name:"model ghi",
			BuilderId:2,
		},
		{
			Id:2,
			Name:"model jkl",
			BuilderId:1,
		},
	})
	s.store.InsertCharters(ctx, s.db, []nausys.Charter{
		{
			Id:1,
			Name:"charter mno",
		},
		{
			Id:2,
			Name:"charter pqr",
		},
	})
	s.store.InsertYachts(ctx, s.db, []*nausys.Yacht{
		{
			Id:1,
			Name:"yacht1",
			CharterId:2,
			ModelId:1,
		},
		{
			Id:2,
			Name:"yacht2",
			CharterId:2,
			ModelId:1,
		},
		{
			Id:3,
			Name:"yacht3",
			CharterId:1,
			ModelId:2,
		},
	})

	for _, tcase := range []struct{
		builderPrefix string
		modelPrefix string
		expectedYachtName string
	} {
		{
			builderPrefix:"builder ab",
			modelPrefix:"model jk",
			expectedYachtName:"yacht3",
		},
	} {
		yachts, total, err := s.store.FindYachts(ctx, tcase.builderPrefix, tcase.modelPrefix, 100, 0)
		s.NoError(err)
		s.Equal(int64(1), total)
		s.Equal(1, len(yachts))
		s.Equal(tcase.expectedYachtName, yachts[0].Name)
	}
}

func TestStorageTestSuite(t *testing.T) {
	suite.Run(t, new(TestStorageSuite))
}
