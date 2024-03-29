package httpd

import (
	"context"
	"io/fs"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rog-golang-buddies/go-automatic-apps/config"
	"github.com/rs/cors"
)

type controller struct {
	mux       *chi.Mux
	server    *http.Server
	webDistFS fs.FS
}

func NewController(webDist fs.FS) *controller {
	c := &controller{
		mux:       chi.NewMux(),
		webDistFS: webDist,
	}

	return c
}

// Start starts the server and blocks until shutdown
func (c *controller) Start(ctx context.Context, config config.ServerConfig) error {

	// Migrate the schema
	for _, model := range config.Models {
		err := config.DB.AutoMigrate(model.Model)
		if err != nil {
			log.Fatalf("Error on AutoMigrate model %s: %s", model.Name, err.Error())
		}
	}

	c.mux.Use(middleware.Recoverer, middleware.Logger)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler(c.mux)

	webRoot, err := fs.Sub(c.webDistFS, "web/dist")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	c.mux.Handle("/*", http.FileServer(http.FS(webRoot)))

	//define endpoints
	c.mux.Get("/api/models", CreateGetModelsHandler(config))
	c.mux.Get("/api/models/{model}/rows", CreateGetModelRows(config))

	c.server = &http.Server{
		Addr:         config.Host + ":" + config.HttpPort,
		Handler:      handler,
		ErrorLog:     log.Default(),     // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	// Create fake data
	p, err := config.FindModel("Product")
	if err != nil {
		log.Fatal(err)
	}

	m := config.DB.Model(p.Model)
	m.Create(map[string]interface{}{
		"Code":  "ABC123",
		"Price": 18,
	})
	m.Create(map[string]interface{}{
		"Code":  "DEF567",
		"Price": 9090,
	})

	log.Println("Starting server")
	return c.server.ListenAndServe()
}

func (c *controller) Shutdown(ctx context.Context) error {
	return c.server.Shutdown(ctx)
}
