package types

import (
	"math"
	"sort"
)

type EnvironmentType string

const (
	EnvTypePipeline EnvironmentType = "PipelineEnv"
	EnvTypePreview  EnvironmentType = "PreviewEnv"
	EnvTypeGlobal   EnvironmentType = "GlobalEnv"
)

type Environment struct {
	IdModel
	Type           EnvironmentType `json:"type"`
	Name           string          `json:"name"`
	Reference      string          `json:"reference"`
	OrgName        string          `json:"orgName"`
	StackId        int64           `json:"stackId"`
	ProviderConfig ProviderConfig  `json:"providerConfig"`
	PipelineOrder  *int            `json:"pipelineOrder"`
}

var _ sort.Interface = EnvsByPipelineOrder{}

type EnvsByPipelineOrder []*Environment

func (envs EnvsByPipelineOrder) Len() int { return len(envs) }

func (envs EnvsByPipelineOrder) Less(i, j int) bool {
	var first int
	if envs[i].PipelineOrder == nil {
		first = math.MaxInt
	} else {
		first = *envs[i].PipelineOrder
	}
	var second int
	if envs[j].PipelineOrder == nil {
		second = math.MaxInt
	} else {
		second = *envs[j].PipelineOrder
	}
	return first < second
}

func (envs EnvsByPipelineOrder) Swap(i, j int) {
	envs[i], envs[j] = envs[j], envs[i]
}
