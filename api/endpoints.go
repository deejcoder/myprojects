package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dilicor/myprojects/storage"
	"github.com/gorilla/mux"
)

func getProjectList(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	col := ac.Db.Collection("projects")

	json.NewEncoder(w).Encode(storage.GetProjects(col))
}

func getProject(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	col := ac.Db.Collection("projects")
	params := mux.Vars(r)

	id := params["id"]

	json.NewEncoder(w).Encode(storage.GetProject(col, id))
}
