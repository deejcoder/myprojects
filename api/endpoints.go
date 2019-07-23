package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dilicor/myprojects/config"
	log "github.com/sirupsen/logrus"
)

func getProjects(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)

	log.Infof("DB: %s", ac.Db.Name())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.GetConfig())
}

func getProject(w http.ResponseWriter, r *http.Request) {

}
