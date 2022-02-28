package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type PromotionConfigs struct {
	Client *Client
}

func (s PromotionConfigs) path(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/promotion-config", s.Client.Config.OrgName, stackId, blockId, envId)
}

func (s PromotionConfigs) Get(stackId, blockId, envId int64) (*types.RunConfig, error) {
	res, err := s.Client.Do(http.MethodGet, s.path(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var runConfig types.RunConfig
	if err := response.ReadJson(res, &runConfig); response.IsNotFoundError(err) {
		return nil, nil
	}
	return &runConfig, nil
}
