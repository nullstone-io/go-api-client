package find

import (
	"context"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func PublicUrls(ctx context.Context, cfg api.Config, stackId int64, wsUid uuid.UUID) ([]string, error) {
	client := api.Client{Config: cfg}
	outputs, err := client.WorkspaceOutputs().GetCurrent(ctx, stackId, wsUid, false)
	if err != nil {
		return nil, err
	} else if outputs == nil {
		return nil, nil
	}
	return ExtractPublicUrlsFromOutputs(outputs), nil
}

func ExtractPublicUrlsFromOutputs(outputs types.Outputs) []string {
	pu, _ := outputs["public_urls"]
	if pu.Value == nil {
		return nil
	}
	rawPublicUrls, _ := pu.Value.([]interface{})
	publicUrls := make([]string, 0)
	for _, cur := range rawPublicUrls {
		publicUrls = append(publicUrls, cur.(string))
	}
	return publicUrls
}
