package main

import (
	"boats/clients/nausys"
	"boats/internal/app/boats"
	"boats/pkg/config"
	storage2 "boats/pkg/storage"
	"boats/pkg/syncer"
	"context"
	"flag"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kilchik/logo/pkg/logo"
	"github.com/pressly/goose"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	confPath := flag.String("conf", "boats.toml", "specify path to configuration file")
	flag.Parse()

	cfg, err := config.Init(*confPath)
	if err != nil {
		log.Fatalf("initialize config: %v", err)
	}

	logo.Init(cfg.GetEnableDebugLogs())

	ctx := context.Background()
	logo.Info(ctx, "starting boats...")

	// Init storage
	db := initDB(cfg.GetDSN())
	storage := storage2.NewStorageImpl(db)

	// Sync nausys db with storage
	nausys := nausys.NewNausysClientImpl(cfg.GetNausysAddr(), cfg.GetNausysUser(), cfg.GetNausysPass())
	syncer := syncer.NewSyncerImpl(nausys, storage)
	if err := syncer.Sync(ctx, false); err != nil {
		logo.Error(ctx, "sync: %v", err)
		os.Exit(1)
	}

	// Init server
	srv := boats.NewBoatsServer(storage, syncer)
	http.Handle("/v1/boats/update", srv.Handler(srv.HandleUpdate))
	http.Handle("/v1/suggest", srv.Handler(srv.HandleSuggest))
	http.Handle("/v1/boats/find", srv.Handler(srv.HandleFind))
	http.Handle("/", http.FileServer(http.Dir("/static")))

	logo.Info(ctx, "start listening %q", cfg.GetListenAddr())
	if err := http.ListenAndServe(cfg.GetListenAddr(), nil); err != nil {
		logo.Error(ctx, "run server: %v", err)
		os.Exit(1)
	}
}

func initDB(dsn string) *sqlx.DB {
	db := sqlx.MustOpen("pgx", dsn)

	for {
		if err := db.Ping(); err != nil {
			log.Printf("ping db failed: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	if err := goose.Up(db.DB, "/migrations"); err != nil {
		log.Fatalf("migrate", err)
	}

	return db
}