package syncer

import (
	"boats/clients/nausys"
	"boats/pkg/storage"
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/kilchik/logo/pkg/logo"
	"github.com/pkg/errors"
	"sort"
	"time"
)

type Syncer interface {
	Sync(ctx context.Context, force bool) error
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

type reservSlots struct {
	from []time.Time
	to   []time.Time
}

func (s *SyncerImpl) Sync(ctx context.Context, force bool) error {
	if !force {
		_, err := s.storage.GetLastUpdateInfo(ctx)
		if err == nil {
			// Already synced and no force demanded
			return nil
		}
		if err != sql.ErrNoRows {
			return err
		}
	}

	// Retrieve data from Nausys
	logo.Info(ctx, "retrieving data from nausys...")
	buildersResp, err := s.nausys.GetBuilders(ctx)
	if err != nil || buildersResp.Status != "OK" {
		return errors.Wrapf(err, "get builders")
	}

	modelsResp, err := s.nausys.GetModels(ctx)
	if err != nil || modelsResp.Status != "OK" {
		return errors.Wrapf(err, "get models")
	}

	chartersResp, err := s.nausys.GetAllCharters(ctx)
	if err != nil || chartersResp.Status != "OK" {
		return errors.Wrapf(err, "get charters")
	}

	parseReservTime := func(str string) time.Time {
		res, err := time.Parse("02.01.2006", str)
		if err != nil {
			return time.Now()
		}
		return res
	}

	var yachts []*nausys.Yacht
	for _, charter := range chartersResp.Companies {
		// Get reservations for this and next years
		var charterReservs []nausys.Reservation
		yearCur := time.Now().Year()
		occupResYearCur, err := s.nausys.GetOccupancy(ctx, charter.Id, yearCur)
		if err != nil {
			return errors.Wrapf(err, "get occupancy for charter %d, year %d", charter.Id, yearCur)
		}
		charterReservs = append(charterReservs, occupResYearCur.Reservations...)
		occupResYearNext, err := s.nausys.GetOccupancy(ctx, charter.Id, yearCur+1)
		if err != nil {
			return errors.Wrapf(err, "get occupancy for charter %d, year %d", charter.Id, yearCur+1)
		}
		charterReservs = append(charterReservs, occupResYearNext.Reservations...)
		reservs := make(map[int64]*reservSlots)
		for _, r := range charterReservs {
			if _, ok := reservs[r.Id]; !ok {
				reservs[r.Id] = &reservSlots{}
			}
			reservs[r.Id].from = append(reservs[r.Id].from, parseReservTime(r.From))
			reservs[r.Id].to = append(reservs[r.Id].to, parseReservTime(r.To))
		}

		// Get boats
		boatsResp, err := s.nausys.GetCharterBoats(ctx, charter.Id)
		if err != nil || boatsResp.Status != "OK" {
			return errors.Wrapf(err, "get boats")
		}

		// Set free slot for every yacht
		for _, y := range boatsResp.Yachts {
			if _, ok := reservs[y.Id]; ok {
				y.AvailableFrom, y.AvailableTo = findNextFreeSlot(reservs[y.Id])
			}
		}

		yachts = append(yachts, boatsResp.Yachts...)
	}

	// Put all into storage
	logo.Info(ctx, "putting data into storage...")
	if err := s.storage.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		if err := s.storage.ClearAll(ctx, tx); err != nil {
			return errors.Wrap(err, "clear all")
		}

		if err := s.storage.InsertBuilders(ctx, tx, buildersResp.Builders); err != nil {
			return errors.Wrap(err, "insert builders into db")
		}

		if err := s.storage.InsertModels(ctx, tx, modelsResp.Models); err != nil {
			return errors.Wrap(err, "insert models into db")
		}

		if err := s.storage.InsertCharters(ctx, tx, chartersResp.Companies); err != nil {
			return errors.Wrap(err, "insert charters into db")
		}

		if err := s.storage.InsertYachts(ctx, tx, yachts); err != nil {
			return errors.Wrap(err, "insert yachts into db")
		}

		if err := s.storage.InsertUpdateInfo(ctx, tx); err != nil {
			return errors.Wrap(err, "insert update info")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "process sync transaction")
	}

	logo.Info(ctx, "synchronized with nausys")

	return nil
}

func findNextFreeSlot(reservs *reservSlots) (from, to sql.NullTime) {
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	// Skip reservations in past
	var reservsFuture reservSlots
	for i := 0; i < len(reservs.from); i++ {
		if reservs.to[i].Before(today) {
			continue
		}
		reservsFuture.from = append(reservsFuture.from, reservs.from[i])
		reservsFuture.to = append(reservsFuture.to, reservs.to[i])
	}

	if len(reservsFuture.from) == 0 && len(reservsFuture.to) == 0 {
		return
	}
	sort.Slice(reservsFuture.to, func(i, j int) bool { return reservsFuture.to[i].Before(reservsFuture.to[j]) })
	sort.Slice(reservsFuture.from, func(i, j int) bool { return reservsFuture.from[i].Before(reservsFuture.from[j]) })

	if today.Add(24 * time.Hour).Before(reservsFuture.from[0]) {
		from = sql.NullTime{today, true}
		to = sql.NullTime{reservsFuture.from[0].Add(-24 * time.Hour), true}
		return
	}

	for i := 0; i < len(reservsFuture.from)-1; i++ {
		if reservsFuture.to[i].Add(48 * time.Hour).Before(reservsFuture.from[i+1]) {
			from = sql.NullTime{reservsFuture.to[i].Add(24 * time.Hour), true}
			to = sql.NullTime{reservsFuture.from[i+1].Add(-24 * time.Hour), true}
			return
		}
	}

	from = sql.NullTime{reservsFuture.to[len(reservsFuture.to)-1].Add(24 * time.Hour), true}

	return
}
