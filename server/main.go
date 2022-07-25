package server

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/rog-golang-buddies/go-automatic-apps/rpc"
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

	fmt.Println("Starting server...")

	webRoot, err := fs.Sub(webDistEmbed, "web/dist")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(webRoot)))

	grpcServer := grpc.NewServer()
	databaseService := rpc.NewDatabaseService()
	rpc.RegisterDatabaseServiceServer(grpcServer, databaseService)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		var grpcServerUrl = config.Host + ":" + config.GRPCPort
		log.Println("gRPC server started at " + grpcServerUrl)
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("RpcServe: ", err)
		}
	}()

	var httpServerUrl = config.Host + ":" + config.HttpPort
	log.Println("http server started at " + httpServerUrl)
	err = http.ListenAndServe(httpServerUrl, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
