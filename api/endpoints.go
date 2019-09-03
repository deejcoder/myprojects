package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/Dilicor/myprojects/config"
	"github.com/Dilicor/myprojects/storage"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func getProjectList(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)

	json.NewEncoder(w).Encode(storage.GetProjects(ac.Db))
}

func getProject(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	params := mux.Vars(r)
	id := params["id"]

	json.NewEncoder(w).Encode(storage.GetProject(ac.Db, id))
}

func onUpdateProject(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)

	var response struct {
		FormErrors url.Values `json:"formErrors"`
	}

	var editedProject *storage.Project
	err := json.NewDecoder(r.Body).Decode(&editedProject)
	if err != nil {
		response.FormErrors.Add("none", "Internal error: unable to decode JSON request")
	}

	response.FormErrors = storage.UpdateProject(ac.Db, editedProject)

	if len(response.FormErrors) > 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(response)
	}
}

func onDeleteProject(w http.ResponseWriter, r *http.Request) {
	db := GetAppContext(r).Db
	params := mux.Vars(r)
	id := params["id"]

	if deleted := storage.DeleteProject(db, id); !deleted {
		w.WriteHeader(http.StatusNotFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {

	// get the provided secret key
	var content struct {
		SecretKey string `json:"secret_key"`
	}
	json.NewDecoder(r.Body).Decode(&content)

	// if the provided key matches the actual key produce a jwt token
	if content.SecretKey == config.GetConfig().AdminSecret {
		tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"timestamp": time.Now(),
		})
		tkString, _ := tk.SignedString([]byte(config.GetConfig().JwtSecret))

		// encode response as JSON
		var response struct {
			Token string `json:"token"`
		}

		response.Token = tkString
		json.NewEncoder(w).Encode(response)
		return
	}

	http.Error(w, "Authentication failed", http.StatusForbidden)
}

// validate checks if some token is valid and returns a JSON response
func validate(w http.ResponseWriter, r *http.Request) {
	var response struct {
		Validated bool `json:"validated"`
	}
	response.Validated = ValidateAuthorization(w, r)
	json.NewEncoder(w).Encode(response)
}
