package xtreamcodes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/charlieplate/rdmate/xtreamcodes"
	"github.com/stretchr/testify/require"
)

var player_api = "player_api.php"

func TestActions_NoneActionHandler(t *testing.T) {
	type test struct {
		name           string
		userInfo       xtreamcodes.UserInfo
		serverInfo     xtreamcodes.ServerInfo
		expectedError  error
		expectedStatus int
	}

	tests := []test{
		{
			name:           "Test Valid Non-Action Handler",
			userInfo:       xtreamcodes.UserInfo{Username: "test", Password: "pass", Auth: 1},
			serverInfo:     xtreamcodes.ServerInfo{Xui: true, Port: "8080", URL: "localhost"},
			expectedError:  nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Test Unauthorized",
			userInfo:       xtreamcodes.UserInfo{Username: "Unauthorized", Password: "pass", Auth: 0},
			serverInfo:     xtreamcodes.ServerInfo{Xui: true, Port: "8080", URL: "localhost"},
			expectedError:  xtreamcodes.ErrUnauthorized,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Test Internal Server Error",
			userInfo:       xtreamcodes.UserInfo{Username: "Internal Server Error", Password: "pass", Auth: 1},
			serverInfo:     xtreamcodes.ServerInfo{},
			expectedError:  fmt.Errorf("server info fetch error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			uf := new(MockUserFetcher)
			sif := new(MockServerInfoFetcher)

			uf.On("FetchUser", tc.userInfo.Username, tc.userInfo.Password).Return(tc.userInfo)
			sif.On("FetchServerInfo").Return(tc.serverInfo)

			var expected []byte
			rec, req := setupHttpTest("GET", []string{player_api}, map[string]string{"username": tc.userInfo.Username, "password": tc.userInfo.Password}, nil)
			svc := xtreamcodes.NewService(uf, sif)

			err := svc.Handle(rec, req)
			require.ErrorIs(t, err, tc.expectedError)
			require.Equal(t, tc.expectedStatus, rec.Code)
			require.Equal(t, "application/json", rec.Header().Get("Content-Type"))

			if tc.expectedError == nil {
				expected, err = json.Marshal(xtreamcodes.NoneActionResponse{ServerInfo: tc.serverInfo, UserInfo: tc.userInfo})
			} else {
				expected, err = json.Marshal(map[string]string{"error": tc.expectedError.Error()})
			}

			require.NoError(t, err)
			require.JSONEq(t, string(expected), rec.Body.String())
		})
	}
}

func setupHttpTest(method string, path []string, params map[string]string, body []byte) (*httptest.ResponseRecorder, *http.Request) {
	p := url.Values{}
	for k, v := range params {
		p.Add(k, v)
	}

	req := httptest.NewRequest(method, fmt.Sprintf("/%s?%s", strings.Join(path, "/"), p.Encode()), bytes.NewReader(body))
	rec := httptest.NewRecorder()

	return rec, req
}
