package syncer

import (
	"boats/clients/nausys"
	"boats/pkg/storage"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"sort"
	"time"
)

type Syncer interface {
	Sync(ctx context.Context, force bool)
}

type SyncerImpl struct {
	nausys  nausys.NausysClient
	storage storage.Storage
}

func NewSyncerImpl(nausys nausys.NausysClient, storage storage.Storage) *SyncerImpl {
	return &SyncerImpl{
		nausys:  nausys,
		storage: storage,
	}
}

func (s *SyncerImpl) Sync(ctx context.Context, force bool) error {
	if force {
		if err := s.storage.ClearAll(ctx); err != nil {
			return errors.Wrap(err, "clear all")
		}
	} else {
		_, err := s.storage.GetLastUpdateInfo(ctx)
		if err == nil {
			// Already synced and no force demanded
			return nil
		}
		if err != sql.ErrNoRows {
			return err
		}
	}

	buildersResp, err := s.nausys.GetBuilders(ctx)
	if err != nil || buildersResp.Status != "OK" {
		return errors.Wrapf(err, "get builders")
	}
	if err := s.storage.InsertBuilders(ctx, buildersResp.Builders); err != nil {
		return errors.Wrap(err, "insert builders into db")
	}

	modelsResp, err := s.nausys.GetModels(ctx)
	if err != nil || modelsResp.Status != "OK" {
		return errors.Wrapf(err, "get models")
	}
	if err := s.storage.InsertModels(ctx, modelsResp.Models); err != nil {
		return errors.Wrap(err, "insert models into db")
	}

	chartersResp, err := s.nausys.GetAllCharters(ctx)
	if err != nil || chartersResp.Status != "OK" {
		return errors.Wrapf(err, "get charters")
	}
	if err := s.storage.InsertCharters(ctx, chartersResp.Companies); err != nil {
		return errors.Wrap(err, "insert charters into db")
	}

	parseReservTime := func(str string) time.Time {
		res, err := time.Parse("02.01.2006", str)
		if err != nil {
			return time.Now()
		}
		return res
	}
	for _, charter := range chartersResp.Companies {
		occupRes, err := s.nausys.GetOccupancy(ctx, charter.Id, time.Now().Year())
		if err != nil {
			return errors.Wrap(err, "get occupancy")
		}
		reservs := make(map[int64]*struct {
			from []time.Time
			to   []time.Time
		})
		for _, r := range occupRes.Reservations {
			if _, ok := reservs[r.Id]; !ok {
				reservs[r.Id] = &struct {
					from []time.Time
					to   []time.Time
				}{}
			}
			reservs[r.Id].from = append(reservs[r.Id].from, parseReservTime(r.From))
			reservs[r.Id].to = append(reservs[r.Id].to, parseReservTime(r.To))
		}
		for _, r := range reservs {
			sort.Slice(r.from, func(i, j int) bool { return r.from[i].Unix() < r.from[j].Unix() })
			sort.Slice(r.to, func(i, j int) bool { return r.to[i].Unix() < r.to[j].Unix() })
		}

		chartersResp, err := s.nausys.GetCharterBoats(ctx, charter.Id)
		if err != nil || chartersResp.Status != "OK" {
			return errors.Wrapf(err, "get boats")
		}
		for _, y := range chartersResp.Yachts {
			if rlist, ok := reservs[y.Id]; ok {
				if len(rlist.from) == 0 {
					continue
				}

				if rlist.from[0].Unix() > time.Now().Unix() {
					y.AvailableFrom = sql.NullTime{time.Now(), true}
					y.AvailableTo = sql.NullTime{rlist.from[0], true}
				} else {
					y.AvailableFrom = sql.NullTime{rlist.to[0], true}
					if len(rlist.from) > 1 {
						y.AvailableTo = sql.NullTime{rlist.from[1], true}
					}
				}
			}
		}
		if err := s.storage.InsertYachts(ctx, chartersResp.Yachts); err != nil {
			return errors.Wrap(err, "insert yachts into db")
		}
	}

	if err := s.storage.InsertUpdateInfo(ctx); err != nil {
		return errors.Wrap(err, "insert update info")
	}

	return nil
}
