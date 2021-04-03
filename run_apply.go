package api

import (
	"time"
)

type RunApply struct {
	CreatedAt time.Time `json:"createdAt"`
	Outputs   Outputs   `json:"outputs"`
	Resources Resources `json:"resources"`
}
