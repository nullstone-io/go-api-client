package types

import "time"

type IdModel struct {
	Id        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
