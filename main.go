package main

import (
	"context"
	krakend2 "github.com/devopsfaith/krakend-ce/krakend"
	"github.com/devopsfaith/krakend-cobra"
	flexibleconfig "github.com/devopsfaith/krakend-flexibleconfig"
	"github.com/devopsfaith/krakend-viper"
	"github.com/devopsfaith/krakend/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	fcPartials  = "FC_PARTIALS"
	fcTemplates = "FC_TEMPLATES"
	fcSettings  = "FC_SETTINGS"
	fcPath      = "FC_OUT"
	fcEnable    = "FC_ENABLE"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case sig := <-sigs:
			log.Println("Signal intercepted:", sig)
			cancel()
		case <-ctx.Done():
		}
	}()
	
	var cfg config.Parser
	cfg = viper.New()
	if os.Getenv(fcEnable) != "" {
		cfg = flexibleconfig.NewTemplateParser(flexibleconfig.Config{
			Parser:    cfg,
			Partials:  os.Getenv(fcPartials),
			Settings:  os.Getenv(fcSettings),
			Path:      os.Getenv(fcPath),
			Templates: os.Getenv(fcTemplates),
		})
	}

	cmd.Execute(cfg, krakend2.NewExecutor(ctx))
	
	log.Println("started KrakeD")
}
