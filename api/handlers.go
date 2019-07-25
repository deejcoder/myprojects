package api

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type appContextKey struct{}

func handler(appContext *AppContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// simple logging
		startTime := time.Now()
		defer func() {
			log.WithFields(log.Fields{
				"remote":   r.RemoteAddr,
				"duration": time.Since(startTime),
			}).Infof("%s %s", r.Method, r.URL.RequestURI())

		}()

		// only return JSON responses
		w.Header().Set("Content-Type", "application/json")

		// add AppContext to request context for access within requests
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
