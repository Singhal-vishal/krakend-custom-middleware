package krakend

import (
	"github.com/devopsfaith/krakend/logging"
	client "github.com/devopsfaith/krakend/transport/http/client/plugin"
	server "github.com/devopsfaith/krakend/transport/http/server/plugin"
)

func LoadPlugins(folder, pattern string, logger logging.Logger) {
	n, err := client.Load(
		folder,
		pattern,
		client.RegisterClient,
	)
	if err != nil {
		logger.Warning("loading plugins:", err)
	}
	logger.Info("total http executor plugins loaded:", n)
	
	n, err = server.Load(
		folder,
		pattern,
		server.RegisterHandler,
	)
	if err != nil {
		logger.Warning("loading plugins:", err)
	}
	logger.Info("total http handler plugins loaded:", n)
}

type pluginLoader struct{}

func (d pluginLoader) Load(folder, pattern string, logger logging.Logger) {
	LoadPlugins(folder, pattern, logger)
}
