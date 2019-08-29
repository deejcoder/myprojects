package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dilicor/myprojects/storage"
)

func getProjectList(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	col := ac.Db.Collection("projects")

	json.NewEncoder(w).Encode(storage.GetProjects(col))
}

func getProject(w http.ResponseWriter, r *http.Request) {

}
