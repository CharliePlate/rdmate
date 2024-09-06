package xtreamcodes

import "errors"

var (
	ErrUnauthorized        = errors.New("unauthenticated")
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
)

type UserInfo struct {
	Username             string   `json:"username"`
	Password             string   `json:"password"`
	Message              string   `json:"message"`
	Auth                 int      `json:"auth"`
	Status               string   `json:"status"`
	ExpDate              string   `json:"exp_date"`
	IsTrial              string   `json:"is_trial"`
	ActiveCons           string   `json:"activeCons"`
	CreatedAt            string   `json:"created_at"`
	MaxConnections       string   `json:"max_connections"`
	AllowedOutputFormats []string `json:"allowed_output_formats"`
}

type ServerInfo struct {
	Xui            bool   `json:"xui"`
	Version        string `json:"version"`
	Revision       int    `json:"revision"`
	URL            string `json:"url"`
	Port           string `json:"port"`
	HTTPSPort      string `json:"https_port"`
	ServerProtocol string `json:"server_protocol"`
	RtmpPort       string `json:"rtmp_port"`
	TimestampNow   int    `json:"timestamp_now"`
	TimeNow        string `json:"time_now"`
	Timezone       string `json:"timezone"`
}

type NoneActionResponse struct {
	UserInfo   UserInfo   `json:"userInfo"`
	ServerInfo ServerInfo `json:"serverInfo"`
}

type UserFetcher interface {
	FetchUser(username, password string) (UserInfo, error)
}

type ServerInfoFetcher interface {
	FetchServerInfo() (ServerInfo, error)
}
