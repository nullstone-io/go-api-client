package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Organizations struct {
	Client *Client
}

func (o Organizations) basePath() string {
	return fmt.Sprintf("organizations")
}

// List - GET /organizations
func (o Organizations) List() ([]types.Organization, error) {
	res, err := o.Client.Do(http.MethodGet, o.basePath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var orgs []types.Organization
	if err := response.ReadJson(res, &orgs); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return orgs, nil
}
