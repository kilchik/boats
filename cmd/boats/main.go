package main

import (
	"boats/clients/nausys"
	"boats/pkg/config"
	storage2 "boats/pkg/storage"
	"boats/pkg/syncer"
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kilchik/logo/pkg/logo"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	cfg, err := config.Init("/Users/akilchik/Documents/go-pets/boats/boats.toml")
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

	//http.HandleFunc("/v1/boats/update", func(w http.ResponseWriter, r *http.Request) {
	//
	//}

	//http.HandleFunc("/v1/boats/suggest", func(w http.ResponseWriter, r *http.Request) {
	//
	//}

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
