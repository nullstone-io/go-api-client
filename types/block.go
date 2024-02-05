package types

type Block struct {
	IdModel
	Type                string                      `json:"type"`
	OrgName             string                      `json:"orgName"`
	StackId             int64                       `json:"stackId"`
	Reference           string                      `json:"reference"`
	Name                string                      `json:"name"`
	IsShared            bool                        `json:"isShared"`
	OwningRepo          string                      `json:"owningRepo"`
	DnsName             string                      `json:"dnsName,omitempty"`
	ModuleSource        string                      `json:"moduleSource"`
	ModuleSourceVersion string                      `json:"moduleSourceVersion"`
	Connections         map[string]ConnectionTarget `json:"connections"`
	Capabilities        []Capability                `json:"capabilities"`
}

type Blocks []Block

func (s Blocks) Find(stackId, blockId int64) *Block {
	for _, block := range s {
		if block.StackId == stackId && block.Id == blockId {
			return &block
		}
	}
	return nil
}

func (s Blocks) FindByName(name string) *Block {
	if name == "" {
		return nil
	}
	for _, block := range s {
		if block.Name == name {
			return &block
		}
	}
	return nil
}
