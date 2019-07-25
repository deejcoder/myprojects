package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dilicor/myprojects/storage"
)

func getProjects(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	col := ac.Db.Collection("projects")

	storage.AddProject(col, "asyncbots", "completed", make([]string, 0), "something special")
	json.NewEncoder(w).Encode(storage.GetProjects(ac.Db.Collection("projects")))
}

func getProject(w http.ResponseWriter, r *http.Request) {

}
