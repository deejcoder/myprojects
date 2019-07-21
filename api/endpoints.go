package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dilicor/myprojects/config"
)

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.GetConfig())
}

func getProject(w http.ResponseWriter, r *http.Request) {

}
