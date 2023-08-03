package api

import "gopkg.in/nullstone-io/go-api-client.v0/auth"

type ClientFactory struct {
	BaseAddress       string
	AccessTokenSource auth.AccessTokenSource
	IsTraceEnabled    bool
}

func (f *ClientFactory) Client(orgName string) *Client {
	return &Client{
		Config: Config{
			OrgName:           orgName,
			BaseAddress:       f.BaseAddress,
			IsTraceEnabled:    f.IsTraceEnabled,
			AccessTokenSource: f.AccessTokenSource,
		},
	}
}
