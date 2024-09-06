package xtreamcodes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/charlieplate/rdmate/internal/assert"
)

type Action string

const (
	ActionVodMovieCategories  Action = "get_vod_categories"
	ActionVodMovieStreams     Action = "get_vod_streams"
	ActionVodSeriesCategories Action = "get_series_categories"
	ActionVodSeriesStreams    Action = "get_series"
	ActionLiveCategories      Action = "get_live_categories"
	ActionLiveStreams         Action = "get_live_streams"
	ActionGetVodInfo          Action = "get_vod_info"
	ActionNone                Action = ""
)

type Service struct {
	userFetcher       UserFetcher
	serverInfoFetcher ServerInfoFetcher
}

func NewService(uf UserFetcher, sif ServerInfoFetcher) *Service {
	return &Service{
		userFetcher:       uf,
		serverInfoFetcher: sif,
	}
}

func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	action := Action(r.FormValue("action"))
	handler := s.getHandlerFor(action)
	if handler == nil {
		return errors.New("no handler found for action")
	}

	return handler.HandleAction(w, r)
}

func (s *Service) getHandlerFor(action Action) ActionHandler {
	switch action {
	case ActionVodMovieCategories:
	case ActionVodMovieStreams:
	case ActionVodSeriesCategories:
	case ActionVodSeriesStreams:
	case ActionLiveCategories:
	case ActionLiveStreams:
	case ActionGetVodInfo:
		assert.TODO()
	default:
		return newNoneActionHandler(s)
	}

	return nil
}

type ActionHandler interface {
	HandleAction(w http.ResponseWriter, r *http.Request) error
	Action() Action
}

type NoneActionHandler struct {
	service *Service
}

func newNoneActionHandler(s *Service) *NoneActionHandler {
	return &NoneActionHandler{
		service: s,
	}
}

func (nah *NoneActionHandler) Action() Action {
	return ActionNone
}

func (nah *NoneActionHandler) HandleAction(w http.ResponseWriter, r *http.Request) error {
	up, err := usernamePasswordFromRequest(r)
	if err != nil {
		writeErrorJson(w, http.StatusBadRequest, err)
	}

	userInfo, err := nah.service.userFetcher.FetchUser(up.Username(), up.Password())
	if err != nil {
		status := 0
		if errors.Is(err, ErrUnauthorized) {
			status = http.StatusUnauthorized
		} else {
			status = http.StatusInternalServerError
		}

		writeErrorJson(w, status, err)
		return err
	}

	serverInfo, err := nah.service.serverInfoFetcher.FetchServerInfo()
	if err != nil {
		writeErrorJson(w, http.StatusInternalServerError, err)
		return err
	}

	resp := NoneActionResponse{
		UserInfo:   userInfo,
		ServerInfo: serverInfo,
	}

	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

type usernamePassword struct {
	username string
	password string
}

func usernamePasswordFromRequest(r *http.Request) (usernamePassword, error) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		return usernamePassword{}, errors.New("username and password are required")
	}

	return usernamePassword{
		username: username,
		password: password,
	}, nil
}

func (up *usernamePassword) Username() string {
	return up.username
}

func (up *usernamePassword) Password() string {
	return up.password
}

func writeErrorJson(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	// if this errors... eh...
	_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
