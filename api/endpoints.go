package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Dilicor/myprojects/config"
	"github.com/Dilicor/myprojects/storage"
	"github.com/dgrijalva/jwt-go"
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

func editProject(w http.ResponseWriter, r *http.Request) {

	if authorized := ValidateAuthorization(w, r); !authorized {
		return
	}

	fmt.Fprintln(w, "In progress")
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
