package api

import (
	"context"
	"net/http"

	"github.com/Dilicor/myprojects/config"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
)

// AppContext stores shared application data for use within Requests
type AppContext struct {
	Db     *mongo.Database
	Config *config.Config
}

type appContextKey struct{}

// AppContextHandler adds the application context (AppContext) to a Request
func AppContextHandler(appContext *AppContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), appContextKey{}, appContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetAppContext returns the AppContext from a given request
func GetAppContext(r *http.Request) *AppContext {
	ac, ok := r.Context().Value(appContextKey{}).(*AppContext)

	if !ok {
		log.Fatal("AppContext was not found in the Request's context")
	}
	return ac
}
