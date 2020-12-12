package duros

import (
	v2 "github.com/metal-stack/duros-go/api/duros/v2"
)

type MockClient struct {
	c v2.DurosAPIClient
}

func NewMock(c v2.DurosAPIClient) *MockClient {
	return &MockClient{
		c: c,
	}
}
