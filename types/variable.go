package types

import (
	"github.com/nullstone-io/module/config"
	"math"
	"reflect"
)

type Variable struct {
	config.Variable `json:",inline"`

	// Value is the exact value set for this variable
	// This Value can be nearly any data type and is determined by Variable Type
	Value interface{} `json:"value"`

	// Redacted indicates that Value is a redacted value and not the real value
	// This happens when retrieving a Variable that has config.Variable Sensitive=true
	Redacted bool `json:"redacted"`

	// Unused signals that the variable is not used by the current module version
	// During promotion of a module into a new workspace, it's possible that the new version removes variables
	// If we removed those variables automatically, a user could face data loss that is unrecoverable
	// Instead, this field was added to signal to the user that they should remove the variable
	Unused bool `json:"unused"`
}

// HasValue determines whether the variable has a set value or if it's unused
// This *cannot* be a pointer receiver method
//
//goland:noinspection GoMixedReceiverTypes
func (v Variable) HasValue() bool {
	return v.Value != nil && !v.Unused
}

func (v *Variable) Redact() bool {
	if v == nil {
		return false
	}

	if v.Sensitive {
		v.Value = nil
		v.Redacted = true
		return true
	}
	return false
}

func (v *Variable) Equal(other Variable) bool {
	return v.SchemaEquals(other) &&
		v.Unused == other.Unused &&
		isVariableValueEqual(v.Type, v.Value, other.Value)
}

func (v *Variable) SchemaEquals(other Variable) bool {
	if v == nil {
		return false
	}
	s1 := v.Variable
	s2 := other.Variable
	return s1.Type == s2.Type &&
		s1.Sensitive == s2.Sensitive &&
		s1.Description == s2.Description &&
		isVariableValueEqual(s1.Type, s1.Default, s2.Default)
}

func (v *Variable) ValueEquals(other Variable) bool {
	val1 := v.Value
	if val1 == nil {
		val1 = v.Default
	}
	val2 := other.Value
	if val2 == nil {
		val2 = other.Default
	}
	return isVariableValueEqual(v.Type, val1, val2)
}

type VariableInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func isVariableValueEqual(varType string, val1, val2 any) bool {
	if val1 == nil {
		if val2 == nil {
			// both nil
			return true
		}
		// val1 is nil, val2 is not nil
		return false
	}
	if val2 == nil {
		// val1 is not nil, val2 is nil
		return false
	}

	// both are not nil
	switch varType {
	case "string":
		return val1 == val2
	case "bool":
		return val1 == val2
	default:
		return reflect.DeepEqual(val1, val2)
	case "number":
		return numericToFloat(val1) == numericToFloat(val2)
	}
}

func numericToFloat(v any) float64 {
	switch val := v.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	}
	return math.NaN()
}
