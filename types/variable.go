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

func (v *Variable) ValueEqualsDefault() bool {
	return isVariableValueEqual(v.Type, v.Value, v.Default)
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
	case "number":
		return numericToFloat(val1) == numericToFloat(val2)
	default:
		return deepValueEqual(val1, val2)
	}
}

// deepValueEqual compares two arbitrary values (objects, maps, lists, etc.) for equality.
// Unlike reflect.DeepEqual, it treats numeric values as equal regardless of their concrete
// type, so float64(60) == int(60). Values decoded from JSON (e.g. the effective config from
// the API) are float64, while values parsed from IaC may be int; without this, identical
// configs for object/map/list variables would appear as spurious changes.
func deepValueEqual(val1, val2 any) bool {
	if f1, ok := tryNumericToFloat(val1); ok {
		f2, ok2 := tryNumericToFloat(val2)
		return ok2 && f1 == f2
	}
	switch v1 := val1.(type) {
	case map[string]any:
		v2, ok := val2.(map[string]any)
		if !ok || len(v1) != len(v2) {
			return false
		}
		for k, a := range v1 {
			b, ok := v2[k]
			if !ok || !deepValueEqual(a, b) {
				return false
			}
		}
		return true
	case []any:
		v2, ok := val2.([]any)
		if !ok || len(v1) != len(v2) {
			return false
		}
		for i := range v1 {
			if !deepValueEqual(v1[i], v2[i]) {
				return false
			}
		}
		return true
	}
	return reflect.DeepEqual(val1, val2)
}

func numericToFloat(v any) float64 {
	if f, ok := tryNumericToFloat(v); ok {
		return f
	}
	return math.NaN()
}

func tryNumericToFloat(v any) (float64, bool) {
	switch val := v.(type) {
	case int:
		return float64(val), true
	case int8:
		return float64(val), true
	case int16:
		return float64(val), true
	case int32:
		return float64(val), true
	case int64:
		return float64(val), true
	case uint:
		return float64(val), true
	case uint8:
		return float64(val), true
	case uint16:
		return float64(val), true
	case uint32:
		return float64(val), true
	case uint64:
		return float64(val), true
	case float32:
		return float64(val), true
	case float64:
		return val, true
	}
	return 0, false
}
