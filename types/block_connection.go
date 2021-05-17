package types

type BlockConnection struct {
	BlockId int64  `json:"blockId"`
	EnvId   *int64 `json:"envId"`
}
