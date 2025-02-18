package types

import (
	"encoding/json"
	"strings"
)

var (
	_ json.Unmarshaler = &ConnectionTargetString{}
	_ json.Marshaler   = &ConnectionTargetString{}
)

// ConnectionTargetString is represented as a ConnectionTarget in memory, but serialized to JSON as a string value
// IDs are ignored from the serialization; only StackName, EnvName, and BlockName are serialized
// Valid Formats:
// - `<block-name>`
// - `<stack-name>.<block-name>`
// - `<stack-name>.<env-name>.<block-name>`
type ConnectionTargetString struct {
	ConnectionTarget `json:",inline"`
}

// UnmarshalJSON is used to customize the parsing of an existing connection target
// Deprecated: This is temporary as we transition Connection.Target from a string to a ConnectionTarget
func (t *ConnectionTargetString) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	var tmp struct {
		StackId   int64  `json:"stackId"`
		StackName string `json:"stackName"`
		BlockId   int64  `json:"blockId"`
		BlockName string `json:"blockName"`
		EnvId     *int64 `json:"envId"`
		EnvName   string `json:"envName"`
	}
	if err := json.Unmarshal(data, &tmp); err == nil {
		t.StackId = tmp.StackId
		t.StackName = tmp.StackName
		t.BlockId = tmp.StackId
		t.BlockName = tmp.BlockName
		t.EnvId = tmp.EnvId
		t.EnvName = tmp.EnvName
		return nil
	}
	// It must be a legacy string, we're going to parse the format: `[[stack.]env.]block`
	tokens := strings.Split(string(data), ".")
	switch len(tokens) {
	case 3:
		t.StackName = tokens[0]
		t.EnvName = tokens[1]
		t.BlockName = tokens[2]
	case 2:
		t.StackName = tokens[0]
		t.BlockName = tokens[1]
	case 1:
		t.BlockName = tokens[0]
	}
	return nil
}

func (t *ConnectionTargetString) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}

	tokens := make([]string, 0)
	if t.StackName != "" {
		tokens = append(tokens, t.StackName)
	}
	if t.EnvName != "" {
		tokens = append(tokens, t.EnvName)
	}
	if t.BlockName != "" {
		tokens = append(tokens, t.BlockName)
	}
	if len(tokens) > 0 {
		return json.Marshal(strings.Join(tokens, "."))
	}
	return json.Marshal("")
}
