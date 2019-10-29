package boats

import (
	"boats/pkg/storage"
	"boats/pkg/syncer"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kilchik/logo/pkg/logo"
	"net/http"
	"strconv"
	"time"
)

const updateInterval = 1 * time.Second

type BoatsServer struct {
	storage   storage.Storage
	syncer    syncer.Syncer
	syncGuard *time.Timer
}

func NewBoatsServer(storage storage.Storage, syncer syncer.Syncer) *BoatsServer {
	return &BoatsServer{
		syncGuard: time.NewTimer(updateInterval),
		syncer:    syncer,
		storage:   storage,
	}
}

func (s *BoatsServer) HandleUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	select {
	case <-s.syncGuard.C:
	default:
		logo.Debug(ctx, "sync in progress")
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

	go func() {
		if err := s.syncer.Sync(ctx, true); err != nil {
			logo.Error(ctx, "sync: %v", err)
		}
		s.syncGuard.Reset(updateInterval)
	}()
}

func (s *BoatsServer) HandleSuggest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	prefix := r.URL.Query().Get("prefix")
	if param == "" || prefix == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var names []string
	var err error
	switch param {
	case "builders":
		names, err = s.storage.FindBuildersByPrefix(ctx, prefix, 5)
	case "models":
		names, err = s.storage.FindModelsByPrefix(ctx, prefix, 5)
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
}

func (s *BoatsServer) HandleFind(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	builderName := r.URL.Query().Get("builder")
	modelName := r.URL.Query().Get("model")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if builderName == "" || modelName == "" || limit == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	yachts, total, err := s.storage.FindYachts(ctx, builderName, modelName, limit, offset)
	if err != nil {
		logo.Error(ctx, "find yachts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(map[string]interface{}{
		"total":  total,
		"yachts": yachts,
	})
	fmt.Fprint(w, string(resp))
}
