package types

type Variables map[string]Variable

func (v Variables) Redact() {
	for k, variable := range v {
		if cur := variable; cur.Redact() {
			v[k] = cur
		}
	}
}

func (v Variables) Equal(other Variables) bool {
	if v == nil {
		return other == nil
	}
	if other == nil {
		return false
	}
	if len(v) != len(other) {
		return false
	}
	for k, variable := range v {
		if otherVar, ok := other[k]; !ok {
			return false
		} else if !variable.Equal(otherVar) {
			return false
		}
	}
	return true
}
