package types

import (
	"github.com/google/uuid"
	"time"
)

type UidCreatedModel struct {
	Uid       uuid.UUID `json:"uid"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}
