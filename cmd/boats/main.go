package main

import (
	"boats/clients/nausys"
	"boats/pkg/config"
	storage2 "boats/pkg/storage"
	"boats/pkg/syncer"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kilchik/logo/pkg/logo"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const updateInterval = 1*time.Second

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

	syncGuard := time.NewTimer(updateInterval)
	http.HandleFunc("/v1/boats/update", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-syncGuard.C:
		default:
			logo.Debug(ctx, "sync in progress")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		go func() {
			if err := syncer.Sync(ctx, true); err != nil {
				logo.Error(ctx, "sync: %v", err)
			}
			syncGuard.Reset(updateInterval)
		}()
	})

	http.HandleFunc("/v1/suggest", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("param")
		prefix := r.URL.Query().Get("prefix")
		if param == "" || prefix == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var names []string
		switch param {
		case "builders":
			names, err = storage.FindBuildersByPrefix(ctx, prefix, 5)
		case "models":
			names, err = storage.FindModelsByPrefix(ctx, prefix, 5)
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err != nil {
			logo.Error(ctx, "find names by prefix in db: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(names); err != nil {
			logo.Error(ctx, "send suggested names: %v", err)
		}
	})

	http.HandleFunc("/v1/boats/find", func(w http.ResponseWriter, r *http.Request) {
		builderName := r.URL.Query().Get("builder")
		modelName := r.URL.Query().Get("model")
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		if builderName == "" || modelName == "" || limit == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		yachts, total, err := storage.FindYachts(ctx, builderName, modelName, limit, offset)
		if err != nil {
			logo.Error(ctx, "find yachts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp, err := json.Marshal(map[string]interface{}{
			"total": total,
			"yachts": yachts,
		})
		fmt.Fprint(w, string(resp))
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	if err := http.ListenAndServe("localhost:9876", nil); err != nil {
		logo.Error(ctx, "run server: %v", err)
		os.Exit(1)
	}
}
