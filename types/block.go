package types

type Block struct {
	IdModel
	Type                string                      `json:"type"`
	OrgName             string                      `json:"orgName"`
	StackId             int64                       `json:"stackId"`
	Reference           string                      `json:"reference"`
	Name                string                      `json:"name"`
	IsShared            bool                        `json:"isShared"`
	DnsName             string                      `json:"dnsName,omitempty"`
	ModuleSource        string                      `json:"moduleSource"`
	ModuleSourceVersion string                      `json:"moduleSourceVersion"`
	Connections         map[string]ConnectionTarget `json:"connections"`
}

type Blocks []Block

func (b *Blocks) Find(stackId, blockId int64) *Block {
	for _, block := range *b {
		if block.StackId == stackId && block.Id == blockId {
			return &block
		}
	}
	return nil
}
