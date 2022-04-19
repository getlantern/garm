package common

import (
	"context"
	"runner-manager/params"
)

type Provider interface {
	// CreateInstance creates a new compute instance in the provider.
	CreateInstance(ctx context.Context, bootstrapParams params.BootstrapInstance) (params.Instance, error)
	// Delete instance will delete the instance in a provider.
	DeleteInstance(ctx context.Context, instance string) error
	// GetInstance will return details about one instance.
	GetInstance(ctx context.Context, instance string) (params.Instance, error)
	// ListInstances will list all instances for a provider.
	ListInstances(ctx context.Context) ([]params.Instance, error)
	// RemoveAllInstances will remove all instances created by this provider.
	RemoveAllInstances(ctx context.Context) error
	// Stop shuts down the instance.
	Stop(ctx context.Context, instance string, force bool) error
	// Start boots up an instance.
	Start(ctx context.Context, instance string) error
}
