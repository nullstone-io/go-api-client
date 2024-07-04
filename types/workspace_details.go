package types

import (
	"encoding/json"
)

type WorkspaceDetails struct {
	OrgName string      `json:"orgName"`
	Stack   Stack       `json:"stack"`
	Env     Environment `json:"env"`
	// BlockRaw is a json.RawMessage because it is a polymorphic type
	// Use Block() to get a concrete type, but returned as "any" value
	// Use BlockType() to get the string type
	BlockRaw json.RawMessage `json:"block"`
}

func (d WorkspaceDetails) BlockType() BlockType {
	var tmp struct {
		Type BlockType `json:"type"`
	}
	json.Unmarshal(d.BlockRaw, &tmp)
	if tmp.Type == "" {
		tmp.Type = BlockTypeBlock
	}
	return tmp.Type
}

func (d WorkspaceDetails) AsBlock() Block {
	var val Block
	json.Unmarshal(d.BlockRaw, &val)
	return val
}

// Block parses the block into the concrete type (e.g. Block, Application, etc.)
// You can use `block, ok := result.(Block)` or `switch result.(type)`
func (d WorkspaceDetails) Block() any {
	switch d.BlockType() {
	case BlockTypeApplication:
		var val Application
		json.Unmarshal(d.BlockRaw, &val)
		return val
	case BlockTypeSubdomain:
		var val Subdomain
		json.Unmarshal(d.BlockRaw, &val)
		return val
	case BlockTypeDomain:
		var val Domain
		json.Unmarshal(d.BlockRaw, &val)
		return val
	default:
		var val Block
		json.Unmarshal(d.BlockRaw, &val)
		return val
	}
}
