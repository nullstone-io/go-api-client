package types

type ConnectionTarget struct {
	StackId   int64  `json:"stackId"`
	BlockId   int64  `json:"blockId,omitempty"`
	BlockName string `json:"blockName"`
	EnvId     *int64 `json:"envId"`
}
