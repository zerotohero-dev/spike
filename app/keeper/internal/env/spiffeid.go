package env

// TODO: configurable from environment.
const trustRoot = "spike.ist"

func SpikeKeeperSpiffeId() string {
	return "spiffe://" + trustRoot + "/spike/keeper"
}

func SpikeNexusSpiffeId() string {
	return "spiffe://" + trustRoot + "/spike/nexus"
}
