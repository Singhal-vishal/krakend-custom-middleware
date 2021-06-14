package krakend

import (
	juju "github.com/devopsfaith/krakend-ratelimit/juju/router/gin"
	router "github.com/devopsfaith/krakend/router/gin"
)

// NewHandlerFactory returns a HandlerFactory with a rate-limit and a metrics collector middleware injected
func NewHandlerFactory() router.HandlerFactory {
	handlerFactory := juju.HandlerFactory
	return handlerFactory
}

type handlerFactory struct{}

func (h handlerFactory) NewHandlerFactory() router.HandlerFactory {
	return NewHandlerFactory()
}
