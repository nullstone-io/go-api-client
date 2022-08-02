package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
)

type PromotionConfigs struct {
	Client *Client
}

func (s PromotionConfigs) path(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/promotion-config", s.Client.Config.OrgName, stackId, blockId, envId)
}

func (s PromotionConfigs) Get(stackId, blockId, envId int64, moduleSourceOverride string) (*types.RunConfig, error) {
	q := url.Values{}
	if moduleSourceOverride != "" {
		q.Set("module-source-override", moduleSourceOverride)
	}
	res, err := s.Client.Do(http.MethodGet, s.path(stackId, blockId, envId), q, nil, nil)
	if err != nil {
		return nil, err
	}

	var runConfig types.RunConfig
	if err := response.ReadJson(res, &runConfig); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &runConfig, nil
}
