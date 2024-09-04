package realdebrid

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type APIConnection struct {
	Client         *http.Client
	Key            string
	BaseURL        string
	TimeoutSeconds int
}

func NewAPIConnection(key string) *APIConnection {
	return &APIConnection{
		Key:            key,
		BaseURL:        "https://api.real-debrid.com/rest/1.0",
		Client:         http.DefaultClient,
		TimeoutSeconds: 5,
	}
}

func (api *APIConnection) User() (UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(api.TimeoutSeconds))
	defer cancel()

	res, err := api.get(ctx, "user")
	if err != nil {
		return UserResponse{}, fmt.Errorf("error making request: %w", err)
	}

	user, err := userFromJson(res)
	if err != nil {
		return UserResponse{}, fmt.Errorf("error unmarshalling user response: %w", err)
	}

	return user, nil
}

func (api *APIConnection) Torrents() (TorrentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(api.TimeoutSeconds))
	defer cancel()

	res, err := api.get(ctx, "torrents")
	if err != nil {
		return TorrentResponse{}, fmt.Errorf("error making request: %w", err)
	}

	torrents, err := torrentsFromJson(res)
	if err != nil {
		return TorrentResponse{}, fmt.Errorf("error unmarshalling torrent response: %w", err)
	}

	return torrents, nil
}

func (api *APIConnection) UnrestrictLink(linkId string) (UnrestrictResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data := url.Values{"link": {linkId}}

	res, err := api.post(ctx, "unrestrict/link", []byte(data.Encode()))
	if err != nil {
		return UnrestrictResponse{}, fmt.Errorf("error making request: %w", err)
	}

	unrestrict, err := unrestrictFromJson(res)
	if err != nil {
		return UnrestrictResponse{}, fmt.Errorf("error unmarshalling unrestrict response: %w", err)
	}

	return unrestrict, nil
}

func (api *APIConnection) baseRequest(ctx context.Context, method string, endpoint string, body []byte) ([]byte, error) {
	if api.Key == "" {
		return nil, errors.New("no API key provided")
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", api.BaseURL, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api.Key))

	if api.Client == nil {
		api.Client = http.DefaultClient
	}

	res, err := api.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making request: %s", res.Status)
	}

	j, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return j, nil
}

func (api *APIConnection) get(ctx context.Context, endpoint string) ([]byte, error) {
	return api.baseRequest(ctx, methodGet, endpoint, nil)
}

func (api *APIConnection) post(ctx context.Context, endpoint string, body []byte) ([]byte, error) {
	return api.baseRequest(ctx, methodPost, endpoint, body)
}
