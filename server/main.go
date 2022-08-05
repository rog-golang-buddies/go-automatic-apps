package server

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/rog-golang-buddies/go-automatic-apps/config"
	"github.com/rog-golang-buddies/go-automatic-apps/server/httpd"
	"golang.org/x/sync/errgroup"
)

//go:embed web/dist
var webDistEmbed embed.FS

func Start(config config.ServerConfig) {
	// Set defaults
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.HttpPort == "" {
		config.HttpPort = "8080"
	}

	// NotifyContext for server graceful shutdown
	serverCtx, serverStop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer serverStop()

	controller := httpd.NewController(webDistEmbed)

	g, gCtx := errgroup.WithContext(serverCtx)
	g.Go(func() error {
		<-gCtx.Done()
		log.Println("Shutting down the server")
		return controller.Shutdown(context.Background())
	})

	err := controller.Start(serverCtx, config)
	if err != http.ErrServerClosed {
		log.Println("Shutting down the server has failed")
	}
}
