package xtreamcodes_test

import (
	"github.com/charlieplate/rdmate/xtreamcodes"
	"github.com/stretchr/testify/mock"
)

type MockUserFetcher struct {
	mock.Mock
}

func (m *MockUserFetcher) FetchUser(username, password string) (xtreamcodes.UserInfo, error) {
	args := m.Called(username, password)
	if args.Get(0).(xtreamcodes.UserInfo).Auth == 0 {
		return xtreamcodes.UserInfo{}, xtreamcodes.ErrUnauthorized
	}

	if args.Get(0).(xtreamcodes.UserInfo).Username == "Bad Request" {
		return xtreamcodes.UserInfo{}, xtreamcodes.ErrBadRequest
	}

	if args.Get(0).(xtreamcodes.UserInfo).Username == "Internal Server Error" {
		return xtreamcodes.UserInfo{}, xtreamcodes.ErrInternalServerError
	}

	return args.Get(0).(xtreamcodes.UserInfo), nil
}

type MockServerInfoFetcher struct {
	mock.Mock
}

func (m *MockServerInfoFetcher) FetchServerInfo() (xtreamcodes.ServerInfo, error) {
	args := m.Called()
	return args.Get(0).(xtreamcodes.ServerInfo), nil
}
