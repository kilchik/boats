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
	"log"
	"net/http"
	"os"
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
	db := sqlx.MustConnect("pgx", cfg.GetDSN())
	storage := storage2.NewStorageImpl(db)
	nausys := nausys.NewNausysClientImpl(cfg.GetNausysAddr(), cfg.GetNausysUser(), cfg.GetNausysPass())
	syncer := syncer.NewSyncerImpl(nausys, storage)
	if err := syncer.Sync(ctx, false); err != nil {
		logo.Error(ctx, "sync: %v", err)
		os.Exit(1)
	}

	srv := boats.NewBoatsServer(storage, syncer)
	http.Handle("/v1/boats/update", srv.Handler(srv.HandleUpdate))
	http.Handle("/v1/suggest", srv.Handler(srv.HandleSuggest))
	http.Handle("/v1/boats/find", srv.Handler(srv.HandleFind))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	if err := http.ListenAndServe(cfg.GetListenAddr(), nil); err != nil {
		logo.Error(ctx, "run server: %v", err)
		os.Exit(1)
	}
}
