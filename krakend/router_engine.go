package krakend

import (
	httpsecure "github.com/devopsfaith/krakend-httpsecure/gin"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/gin-gonic/gin"
)

func NewEngine(cfg config.ServiceConfig, logger logging.Logger) *gin.Engine {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: gin.DefaultWriter}), gin.Recovery())
	
	engine.RedirectTrailingSlash = true
	engine.RedirectFixedPath = true
	engine.HandleMethodNotAllowed = true
	
	if err := httpsecure.Register(cfg.ExtraConfig, engine); err != nil {
		logger.Warning(err)
	}
	
	return engine
}

type engineFactory struct{}

func (e engineFactory) NewEngine(cfg config.ServiceConfig, l logging.Logger) *gin.Engine {
	return NewEngine(cfg, l)
}
