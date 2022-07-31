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
	"github.com/rog-golang-buddies/go-automatic-apps/rpc"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type ServerConfig struct {
	Host     string
	HttpPort string
	GRPCPort string
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
	if config.GRPCPort == "" {
		config.GRPCPort = "8081"
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

	grpcServer := grpc.NewServer()
	databaseService := rpc.NewDatabaseService()
	rpc.RegisterDatabaseServiceServer(grpcServer, databaseService)
	g.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort))
		if err != nil {
			log.Panicf("grpc listen: %v", err.Error())
			return err
		}

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Panicf("grpc serve: %v", err.Error())
			return err
		}

		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()
		grpcServer.GracefulStop()
		return httpServer.Shutdown(context.Background())
	})
}
