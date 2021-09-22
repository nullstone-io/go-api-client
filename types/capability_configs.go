package types

type CapabilityConfigs []CapabilityConfig

func (s CapabilityConfigs) ExceptNeedsDestroyed() CapabilityConfigs {
	result := make(CapabilityConfigs, 0)
	for _, cur := range s {
		if !cur.NeedsDestroyed {
			result = append(result, cur)
		}
	}
	return result
}

func (s CapabilityConfigs) TfModuleAddrs() []string {
	result := make([]string, 0)
	for _, cur := range s {
		result = append(result, cur.TfModuleAddr())
	}
	return result
}
