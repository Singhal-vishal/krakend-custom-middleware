package krakend

import (
	"context"
	cb "github.com/devopsfaith/krakend-circuitbreaker/gobreaker/proxy"
	"github.com/devopsfaith/krakend-martian"
	juju "github.com/devopsfaith/krakend-ratelimit/juju/proxy"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/transport/http/client"
	httprequestexecutor "github.com/devopsfaith/krakend/transport/http/client/plugin"
)

// NewBackendFactory creates a BackendFactory by stacking all the available middlewares:
// - oauth2 client credentials
// - http cache
// - martian
// - pubsub
// - amqp
// - cel
// - lua
// - rate-limit
// - circuit breaker
// - metrics collector
// - opencensus collector
func NewBackendFactory(logger logging.Logger) proxy.BackendFactory {
	return NewBackendFactoryWithContext(context.Background(), logger)
}

// NewBackendFactory creates a BackendFactory by stacking all the available middlewares and injecting the received context
func NewBackendFactoryWithContext(ctx context.Context, logger logging.Logger) proxy.BackendFactory {
	requestExecutorFactory := func(cfg *config.Backend) client.HTTPRequestExecutor {
		return client.DefaultHTTPRequestExecutor(client.NewHTTPClient)
	}
	requestExecutorFactory = httprequestexecutor.HTTPRequestExecutor(logger,requestExecutorFactory)
	backendFactory := martian.NewConfiguredBackendFactory(logger, requestExecutorFactory)
	backendFactory = juju.BackendFactory(backendFactory)
	backendFactory = cb.BackendFactory(backendFactory, logger)
	return backendFactory
}

type backendFactory struct{}

func (b backendFactory) NewBackendFactory(ctx context.Context, l logging.Logger) proxy.BackendFactory {
	return NewBackendFactoryWithContext(ctx, l)
}
