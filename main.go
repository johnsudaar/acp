package main

import (
	"context"

	"github.com/Scalingo/go-utils/logger"

	"github.com/johnsudaar/acp/config"
	"github.com/johnsudaar/acp/devices/drivers"
	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/webserver"
	"github.com/pkg/errors"
)

func main() {
	// ------------ Initialization -------------------
	// Load App config
	err := config.Init()
	if err != nil {
		panic(errors.Wrap(err, "fail to init config"))
	}
	// Logger init
	log := logger.Default()
	ctx := logger.ToCtx(context.Background(), log)
	log.Info("Config initialized")
	// Load devices drivers
	drivers.LoadDrivers()
	log.Info("Drivers initialized")

	// Load current graph
	graph, err := graph.Load(ctx)
	if err != nil {
		panic(err)
	}
	log.Info("Graph loaded")

	log.Info("Init phase done.")

	log.Info("Starting services")
	// ------------ Start ----------------------------
	err = webserver.Start(ctx, graph)
	if err != nil {
		panic(err)
	}
}
