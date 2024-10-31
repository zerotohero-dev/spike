package config

import "os"

const NexusVersion = "0.1.0"
const PilotVersion = "0.1.0"
const KeeperVersion = "0.1.0"

func SpiffeEndpointSocket() string {
	p := os.Getenv("SPIFFE_ENDPOINT_SOCKET")
	if p != "" {
		return p
	}

	// The socket address below is good for development purposes;
	// however, configuring a more restricted socket like the following
	// is recommended; and it's what SPIRE Helm Charts use by default:
	// "unix:///run/spire/agent/sockets/spire.sock"
	//
	// If We implement a Kubernetes operation mode for SPIKE, we might
	// change this return value based on the system SPIKE is deployed.
	// As in:
	// for dev linux mode: unix:///tmp/spire-agent/public/api.sock
	// for k8s: unix:///run/spire/agent/sockets/spire.sock"
	return "unix:///tmp/spire-agent/public/api.sock"
}
