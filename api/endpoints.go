package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/deejcoder/myprojects/config"
	reply "github.com/deejcoder/myprojects/reply"
	"github.com/deejcoder/myprojects/storage"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

/* getProjectList
GET /projects, returns all projects
*/
func getProjectList(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	response := reply.NewReply()

	projects := storage.GetProjects(ac.Db)
	response.Success(w, "Obtained project list", projects)
}

/* getProject
GET /project/{id}/ returns data for a particular project
*/
func getProject(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	params := mux.Vars(r)
	id := params["id"]
	response := reply.NewReply()

	project := storage.GetProject(ac.Db, id)
	response.Success(w, "Obtained project information", project)
}

/* onUpdateProject
POST /project/{id}/update, with new project data in the request body,
replaces an existing project
*/
func onUpdateProject(w http.ResponseWriter, r *http.Request) {
	ac := GetAppContext(r)
	resp := reply.NewReply()

	// get the project
	var newProject *storage.Project
	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		resp.Error(w, "The provided JSON document contained syntax errors", reply.ErrorValidationError)
		return
	}

	// check if project is valid
	if valid := newProject.Validate(&resp); !valid {
		resp.Error(w, "There were validation error(s) when updating the project", reply.ErrorValidationError)
		return
	}

	// update the database
	if updated := storage.UpdateProject(ac.Db, newProject); !updated {
		resp.Error(w, "Internal server error", reply.ErrorInternalError)
		return
	}

	resp.Success(w, "The project was successfully updated", nil)
}

/* onDeleteProject
DELETE /project/{id}/delete, deletes a project by ID
*/
func onDeleteProject(w http.ResponseWriter, r *http.Request) {
	db := GetAppContext(r).Db
	params := mux.Vars(r)
	id := params["id"]
	resp := reply.NewReply()

	// delete the project
	if deleted := storage.DeleteProject(db, id); !deleted {
		resp.Error(w, "Failed to delete the project (internal error)", reply.ErrorInternalError)
		return
	}

	resp.Success(w, "Project was deleted successfully", nil)
}

/* login
POST /auth/login, client provides secret_key (passphase), returns token
*/
func login(w http.ResponseWriter, r *http.Request) {
	resp := reply.NewReply()

	// get the secret key
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

		var response struct {
			Token string `json:"token"`
		}
		response.Token = tkString
		resp.Success(w, "Login validated", response)
		return
	}

	resp.Error(w, "Authentication failed", reply.ErrorNotAuthorized)
}

/* validate
GET /auth/validate returns true/false if client has valid jwt token
*/
func validate(w http.ResponseWriter, r *http.Request) {
	resp := reply.NewReply()
	if validated := ValidateAuthorization(w, r); !validated {
		resp.Error(w, "Token not OK", reply.ErrorNotAuthorized)
		return
	}
	resp.Success(w, "Token OK", nil)
}
