package types

type DependencyConfigs []DependencyConfig

func (d DependencyConfigs) FindByReference(reference WorkspaceTarget) *DependencyConfig {
	for _, c := range d {
		if c.Reference.Id() == reference.Id() {
			return &c
		}
	}
	return nil
}
