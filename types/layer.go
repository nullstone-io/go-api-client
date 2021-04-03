package types

type Layer string

const (
	LayerNetwork     Layer = "network"
	LayerCluster     Layer = "cluster"
	LayerDatabase    Layer = "database"
	LayerService     Layer = "service"
	LayerPublicEntry Layer = "public-entry"
)
