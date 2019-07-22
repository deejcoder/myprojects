package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Dilicor/myprojects/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Serve starts the webserver, terminates on request
func Serve(ctx context.Context) {
	cfg := config.GetConfig()

	// cross origin resource sharing restrictions
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// endpoints
	router := mux.NewRouter()
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/project/{slug}", getProject).Methods("GET")

	// init webserver
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.API.Port),
		Handler: cors(router),
	}

	// listen for interupt signal
	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
		close(done)
	}()

	// start webserver
	log.Infof("Serving web api on http://localhost:%d", cfg.API.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Error(err)
	}

	<-done
}
