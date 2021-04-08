package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Scalingo/go-utils/logger"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/urfave/cli"

	"github.com/johnsudaar/acp/config"
	"github.com/johnsudaar/acp/devices/drivers"
	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/realtime"
	"github.com/johnsudaar/acp/scenes"
	"github.com/johnsudaar/acp/tests/proxy"
	"github.com/johnsudaar/acp/timer"
	"github.com/johnsudaar/acp/webserver"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	Version = "dev"
)

func main() {
	if len(os.Args) == 1 {
		StartServer(true)
		return
	}
	app := &cli.App{
		Name:    "acp",
		Version: Version,
		Commands: []cli.Command{
			{
				Name: "start",
				Action: func(c *cli.Context) error {
					StartServer(false)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func StartServer(gui bool) {
	jww.SetStdoutThreshold(jww.LevelDebug)

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

	// Init centrifuge
	realtime := realtime.New()

	// Load current graph
	graph, err := graph.Load(ctx, realtime)
	if err != nil {
		panic(err)
	}
	log.Info("Graph loaded")

	// Init scenes
	scenes := scenes.New(realtime)

	// Init timers
	timers, err := timer.Load(ctx, realtime, graph)
	if err != nil {
		panic(err)
	}

	err = realtime.Start(ctx, graph)
	if err != nil {
		panic(err)
	}

	log.Info("Init phase done.")

	log.Info("Starting services")
	// ------------ Start ----------------------------
	go proxy.Start()
	if !gui {
		err := webserver.Start(ctx, graph, realtime, timers, scenes)
		if err != nil {
			panic(err)
		}
		return
	} else {
		go func() {
			err := webserver.Start(ctx, graph, realtime, timers, scenes)
			if err != nil {
				panic(err)
			}
		}()

		// Initialize astilectron
		var a, _ = astilectron.New(log, astilectron.Options{
			AppName: "ACP",
		})
		defer a.Close()

		// Start astilectron
		a.Start()

		config := config.Get()
		serverURL := fmt.Sprintf("http://localhost:%v/index.html", config.Server.Port)
		fmt.Println(serverURL)

		var w, _ = a.NewWindow(serverURL, &astilectron.WindowOptions{
			Center: astikit.BoolPtr(true),
			Height: astikit.IntPtr(1024),
			Width:  astikit.IntPtr(1024),
		})
		w.Create()
		a.Wait()
	}
}
