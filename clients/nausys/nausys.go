package nausys

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const timeoutDefault = 10 * time.Second

type nausysCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Charter struct {
	Id   int64  `json:"id",db:"id"`
	Name string `json:"name",db:"name"`
}

type GetAllChartersResponse struct {
	Status    string    `json:"status"`
	Companies []Charter `json:"companies"`
}

type Yacht struct {
	Id            int64        `json:"id",db:"id"`
	Name          string       `json:"name",name:"name"`
	CharterId     int64        `json:"companyId",db:"charter_id"`
	ModelId       int64        `json:"yachtModelId",db:"model_id"`
	AvailableFrom sql.NullTime `db:"available_from"`
	AvailableTo   sql.NullTime `db:"available_to"`
}

type GetCharterBoatsResponse struct {
	Status string   `json:"status"`
	Yachts []*Yacht `json:"yachts"`
}

type Model struct {
	Id        int64  `json:"id",db:"id"`
	Name      string `json:"name",db:"name"`
	BuilderId int64  `json:"yachtBuilderId",db:"builder_id"`
}

type GetModelsResponse struct {
	Status string  `json:"status"`
	Models []Model `json:"models"`
}

type Builder struct {
	Id   int64  `json:"id",db:"id"`
	Name string `json:"name",db:"name"`
}

type GetBuildersResponse struct {
	Status   string    `json:"status"`
	Builders []Builder `json:"builders"`
}

type Reservation struct {
	Id   int64  `json:"yachtId"`
	From string `json:"periodFrom"`
	To   string `json:"periodTo"`
}

type GetOccupancyResponse struct {
	Reservations []Reservation `json:"reservations"`
}

type NausysClient interface {
	GetAllCharters(ctx context.Context) (*GetAllChartersResponse, error)
	GetCharterBoats(ctx context.Context, charterId int64) (*GetCharterBoatsResponse, error)
	GetModels(ctx context.Context) (*GetModelsResponse, error)
	GetBuilders(ctx context.Context) (*GetBuildersResponse, error)
	GetOccupancy(ctx context.Context, charterId int64, year int) (*GetOccupancyResponse, error)
}

type NausysClientImpl struct {
	addr       string
	creds      *nausysCreds
	httpClient *http.Client
}

func NewNausysClientImpl(addr, userName, userPass string) *NausysClientImpl {
	return &NausysClientImpl{
		addr: addr,
		creds: &nausysCreds{
			Username: userName,
			Password: userPass,
		},
		httpClient: &http.Client{
			Timeout: timeoutDefault,
		},
	}
}

func (nc *NausysClientImpl) GetAllCharters(ctx context.Context) (*GetAllChartersResponse, error) {
	resp := &GetAllChartersResponse{}
	if err := nc.processPost(ctx, "/catalogue/v6/charterCompanies", resp); err != nil {
		return nil, errors.Wrap(err, "get all charters from nausys")
	}
	return resp, nil
}

func (nc *NausysClientImpl) GetCharterBoats(ctx context.Context, charterId int64) (*GetCharterBoatsResponse, error) {
	resp := &GetCharterBoatsResponse{}
	if err := nc.processPost(ctx, fmt.Sprintf("/catalogue/v6/yachts/%d", charterId), resp); err != nil {
		return nil, errors.Wrap(err, "get charter boats from nausys")
	}
	return resp, nil
}

func (nc *NausysClientImpl) GetModels(ctx context.Context) (*GetModelsResponse, error) {
	resp := &GetModelsResponse{}
	if err := nc.processPost(ctx, "/catalogue/v6/yachtModels", resp); err != nil {
		return nil, errors.Wrap(err, "get yacht models from nausys")
	}
	return resp, nil
}

func (nc *NausysClientImpl) GetBuilders(ctx context.Context) (*GetBuildersResponse, error) {
	resp := &GetBuildersResponse{}
	if err := nc.processPost(ctx, "/catalogue/v6/yachtBuilders", resp); err != nil {
		return nil, errors.Wrap(err, "get yacht builders from nausys")
	}
	return resp, nil
}

func (nc *NausysClientImpl) GetOccupancy(ctx context.Context, charterId int64, year int) (*GetOccupancyResponse, error) {
	resp := &GetOccupancyResponse{}
	if err := nc.processPost(ctx, fmt.Sprintf("/yachtReservation/v6/occupancy/%d/%d", charterId, year), resp); err != nil {
		return nil, errors.Wrap(err, "get yacht reservations from nausys")
	}
	return resp, nil
}

func (nc *NausysClientImpl) processPost(ctx context.Context, ep string, respBody interface{}) error {
	req, _ := json.Marshal(nc.creds)
	resp, err := nc.httpClient.Post(nc.addr+ep, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return errors.Wrap(err, "http post to nausys")
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return errors.Wrap(err, "decode charters-boats response body")
	}
	return nil
}
