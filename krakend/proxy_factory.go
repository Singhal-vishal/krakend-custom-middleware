package krakend

import (
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
)

func NewProxyFactory(logger logging.Logger, backendFactory proxy.BackendFactory) proxy.Factory {
	proxyFactory := proxy.NewDefaultFactory(backendFactory, logger)
	proxyFactory = proxy.NewShadowFactory(proxyFactory)
	return proxyFactory
}

type proxyFactory struct{}

func (p proxyFactory) NewProxyFactory(logger logging.Logger, backendFactory proxy.BackendFactory) proxy.Factory {
	return NewProxyFactory(logger, backendFactory)
}
