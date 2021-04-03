package api

import (
	"github.com/google/uuid"
	"github.com/nullstone-io/module/config"
)

type ModuleVersion struct {
	UidCreatedModel
	ModuleUid uuid.UUID       `json:"moduleUid"`
	Version   string          `json:"version"`
	Manifest  config.Manifest `json:"manifest"`
}
