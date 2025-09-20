package types

// NamedCapabilityConfigs is a map of CapabilityConfig indexed by Capability.Name
type NamedCapabilityConfigs map[string]CapabilityConfig

type CapabilityConfigs []CapabilityConfig

func (s CapabilityConfigs) FindByName(name string) *CapabilityConfig {
	for _, c := range s {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

// FindById
// Deprecated
func (s CapabilityConfigs) FindById(id int64) *CapabilityConfig {
	for _, c := range s {
		if c.Id == id {
			return &c
		}
	}
	return nil
}

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
