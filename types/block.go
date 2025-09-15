package types

type BlockType string

const (
	BlockTypeApplication      BlockType = "Application"
	BlockTypeDatastore        BlockType = "Datastore"
	BlockTypeSubdomain        BlockType = "Subdomain"
	BlockTypeDomain           BlockType = "Domain"
	BlockTypeIngress          BlockType = "Ingress"
	BlockTypeClusterNamespace BlockType = "ClusterNamespace"
	BlockTypeCluster          BlockType = "Cluster"
	BlockTypeNetwork          BlockType = "Network"
	BlockTypeBlock            BlockType = "Block"
)

type Block struct {
	IdModel
	Type       string `json:"type"`
	OrgName    string `json:"orgName"`
	StackId    int64  `json:"stackId"`
	Reference  string `json:"reference"`
	Name       string `json:"name"`
	IsShared   bool   `json:"isShared"`
	OwningRepo string `json:"owningRepo"`
	Repo       string `json:"repo,omitempty"`
	RepoUrl    string `json:"repoUrl,omitempty"`
	Framework  string `json:"framework,omitempty"`
	DnsName    string `json:"dnsName,omitempty"`

	// ModuleSource
	// Deprecated - This has been moved to WorkspaceConfig
	ModuleSource string `json:"moduleSource"`
	// ModuleSourceVersion
	// Deprecated - This has been moved to WorkspaceConfig
	ModuleSourceVersion string `json:"moduleSourceVersion"`
	// Connections
	// Deprecated - This has been moved to WorkspaceConfig
	Connections map[string]ConnectionTarget `json:"connections"`
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
