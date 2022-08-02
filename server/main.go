package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

type ServerConfig struct {
	Host     string
	HttpPort string
	Tables   []*schema.Table
}

//go:embed web/dist
var webDistEmbed embed.FS

func Start(config ServerConfig) {

	// Set defauls
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.HttpPort == "" {
		config.HttpPort = "8080"
	}

	// NotifyContext for server graceful shutdown
	serverCtx, serverStop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer serverStop()

	// create a new serve mux to attach handlers to it and enable Cors policy
	mux := http.NewServeMux()

	c := cors.Options{
		AllowedOrigins: []string{"*"},
	}
	handler := cors.New(c).Handler(mux)

	httpServer := http.Server{
		Addr:         config.Host + ":" + config.HttpPort,
		Handler:      handler,
		ErrorLog:     log.Default(),     // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		BaseContext: func(_ net.Listener) context.Context {
			return serverCtx
		},
	}

	fmt.Println("Starting server...")

	webRoot, err := fs.Sub(webDistEmbed, "web/dist")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	mux.Handle("/", http.FileServer(http.FS(webRoot)))

	g, gCtx := errgroup.WithContext(serverCtx)
	g.Go(func() error {
		// Run the server
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})
}
