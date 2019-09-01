package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Dilicor/myprojects/config"
	"github.com/Dilicor/myprojects/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// AppContext stores shared application data for use within Requests
type AppContext struct {
	Db     *mongo.Database
	Config *config.Config
}

func configure(ac *AppContext) *http.Server {
	// cross origin resource sharing rules
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// endpoints
	router := mux.NewRouter()
	router.HandleFunc("/projects", getProjectList).Methods("GET")
	router.HandleFunc("/project/{id}/", getProject).Methods("GET")
	router.HandleFunc("/project/{id}/edit", editProject).Methods("GET")
	router.HandleFunc("/auth/login", login).Methods("POST")
	router.HandleFunc("/auth/validate", validate).Methods("GET")

	// configure server
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", ac.Config.API.Port),
		Handler: handler(ac, cors(router)),
	}

	return s
}

// Start starts the webserver, terminates on request
func Start(ctx context.Context) {

	db := storage.Connect()
	appContext := AppContext{
		Db:     db,
		Config: config.GetConfig(),
	}

	server := configure(&appContext)

	// listen for interupt signal
	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
		close(done)
	}()

	// start webserver
	log.Infof("Starting REST api on http://localhost:%d", appContext.Config.API.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Error(err)
	}

	<-done
}
