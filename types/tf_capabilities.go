package types

type TfCapabilities []TfCapability

func (s TfCapabilities) ExceptNeedsDestroyed() TfCapabilities {
	result := make(TfCapabilities, 0)
	for _, cur := range s {
		if !cur.NeedsDestroyed {
			result = append(result, cur)
		}
	}
	return result
}

func (s TfCapabilities) TfModuleAddrs() []string {
	result := make([]string, 0)
	for _, cur := range s {
		result = append(result, cur.TfModuleAddr())
	}
	return result
}
