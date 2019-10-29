package app

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bluest-eel/cli/common"
	"github.com/bluest-eel/cli/components"
	log "github.com/sirupsen/logrus"
)

// Application ...
type Application struct {
	components.CLI
}

// ProcessCLIOptions ...
func (app *Application) ProcessCLIOptions() {
	log.Debug("Processing CLI options ...")
	// The ability to delay application execution is added here to assist in
	// local development environment setups.
	delayPtr := flag.Int("delay", 0, "Delay application startup, in milliseconds")
	versionPtr := flag.Bool("version", false, "Display version/build info and exit")
	flag.Parse()

	if *versionPtr {
		println("Version: ", common.VersionString())
		println("Build: ", common.BuildString())
		os.Exit(0)
	}

	if delay := *delayPtr; delay > 0 {
		log.Infof("Delaying application startup for %d milliseconds ...", delay)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		log.Debug("Application execution resuming ...")
	} else {
		log.Debug("No application delay option was provided; continuing ...")
	}
	log.Info("CLI options processed.")
}

// HandleSignals ...
func (app *Application) HandleSignals() {
	log.Debug("Setting up signal handler ...")
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		s := <-signalChan
		log.Debugf("Signal: %#v", s)
		switch s {
		case syscall.SIGINT: // ^C
			log.Debug("Received SIGNINT signal; quitting ...")
			os.Exit(0)
		case syscall.SIGTERM:
			log.Debug("Received SIGTERM signal; quitting ...")
			os.Exit(0)
		case syscall.SIGQUIT:
			log.Debug("Received SIGQUIT signal; quitting ...")
			os.Exit(1)
		default:
			log.Debugf("Received unexpected signal %#v", s)
		}
	}()
	log.Info("Signal handler is set up.")
}

// // SetupDB ...
// func (app *Application) SetupDB() db.Database {
// 	log.Debug("Setting up database ...")
// 	conn, err := db.Open(app.Config)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Info("Database is set up.")
// 	return conn
// }

// // SetupCache ...
// func (app *Application) SetupCache() cache.Cache {
// 	log.Debug("Setting up cache ...")
// 	c, err := cache.Open(app.Config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Info("Cache is set up.")
// 	return c
// }

// // SetHTTPDRoutes ...
// func (app *Application) SetHTTPDRoutes() {
// 	log.Debug("Setting up HTTPD routes ...")
// 	s := handlers.NewHTTPHandlerServer(&app.DB)

// 	// Basic routes
// 	app.HTTPD.GET("/", s.HomeRoute)
// 	app.HTTPD.GET("/health", s.HealthRoute, middlewareutil.VersionInfo())

// 	log.Info("HTTPD routes are set up.")
// }

// // SetHTTPDMiddleware ...
// func (app *Application) SetHTTPDMiddleware() {
// 	log.Debug("Setting up HTTPD middleware ...")
// 	app.HTTPD.Pre(middleware.RemoveTrailingSlash())
// 	app.HTTPD.Use(middleware.Logger())
// 	app.HTTPD.Use(middleware.Recover())
// 	log.Info("HTTPD middleware set up.")
// }

// // SetupgRPCImplementation ...
// func (app *Application) SetupgRPCImplementation(r *reverb.Reverb) {
// 	log.Debug("Setting up gRPC implementation ...")
// 	s := handlers.NewGRPCHandlerServer(&app.DB, app.Config)
// 	s.RegisterServer(r.GRPCServer)
// 	log.Info("gRPC implementation is set up.")
// }

// // StartgRPCD ...
// func (app *Application) StartgRPCD() {
// 	log.Debug("Starting gRPC daemon ...")
// 	serverOpts := app.Config.GraphGRPCConnectionString()
// 	server := app.GRPCD.Start(serverOpts)
// 	app.SetupgRPCImplementation(server)
// 	go server.Serve()
// 	log.Infof("gRPC daemon started on %s.", serverOpts)
// }

// // StartHTTPD ...
// func (app *Application) StartHTTPD() {
// 	log.Debug("Starting HTTP daemon ...")
// 	server := app.HTTPD.Start(app.Config.HTTPConnectionString())
// 	app.HTTPD.Logger.Fatal(server)
// }

// Start ...
func (app *Application) Start() {
	// app.StartgRPCD()
	// The HTTPD component should be the last one started, since we're just
	// going to use its daemon for ours.
	// app.StartHTTPD()
}
