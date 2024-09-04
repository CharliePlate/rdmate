package realdebrid

import (
	"encoding/json"
	"fmt"
)

type UserResponse struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Locale     string `json:"locale"`
	Avatar     string `json:"avatar"`
	Type       string `json:"type"`
	Expiration string `json:"expiration"`
	ID         int    `json:"id"`
	Points     int    `json:"points"`
	Premium    int    `json:"premium"`
}

func userFromJson(data []byte) (UserResponse, error) {
	var user UserResponse
	err := json.Unmarshal(data, &user)
	if err != nil {
		return user, fmt.Errorf("error unmarshalling user response: %w", err)
	}

	return user, nil
}

type TorrentResponse []struct {
	ID       string   `json:"id"`
	Filename string   `json:"filename"`
	Hash     string   `json:"hash"`
	Host     string   `json:"host"`
	Status   string   `json:"status"`
	Added    string   `json:"added"`
	Ended    string   `json:"ended"`
	Links    []string `json:"links"`
	Bytes    int64    `json:"bytes"`
	Split    int      `json:"split"`
	Progress int      `json:"progress"`
}

func torrentsFromJson(data []byte) (TorrentResponse, error) {
	var torrents TorrentResponse
	err := json.Unmarshal(data, &torrents)
	if err != nil {
		return torrents, fmt.Errorf("error unmarshalling torrent response: %w", err)
	}

	return torrents, nil
}

type UnrestrictResponse struct {
	ID         string `json:"id"`
	Filename   string `json:"filename"`
	MimeType   string `json:"mimeType"`
	Link       string `json:"link"`
	Host       string `json:"host"`
	HostIcon   string `json:"host_icon"`
	Download   string `json:"download"`
	Filesize   int64  `json:"filesize"`
	Chunks     int    `json:"chunks"`
	Crc        int    `json:"crc"`
	Streamable int    `json:"streamable"`
}

func unrestrictFromJson(data []byte) (UnrestrictResponse, error) {
	var unrestrict UnrestrictResponse
	err := json.Unmarshal(data, &unrestrict)
	if err != nil {
		return unrestrict, fmt.Errorf("error unmarshalling unrestrict response: %w", err)
	}

	return unrestrict, nil
}

const (
	methodGet  = "GET"
	methodPost = "POST"
)
