package cluster

import (
	"github.com/docker/engine-api/types/network"
	"golang.org/x/net/context"
)

// Provider provides clustering config details
type Provider interface {
	IsManager() bool
	IsAgent() bool
	GetLocalAddress() string
	GetAdvertiseAddress() string
	GetRemoteAddress() string
	ListenClusterEvents() <-chan struct{}
	AttachNetwork(string, string, []string) (*network.NetworkingConfig, error)
	DetachNetwork(string, string) error
	UpdateAttachment(string, string, *network.NetworkingConfig) error
	WaitForDetachment(context.Context, string, string, string, string) error
}
