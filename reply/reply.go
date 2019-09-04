package reply

import (
	"encoding/json"
	"net/http"
)

/*Response is a standardized response for this API which
is encoded in the body of a HTTP response
*/
type Response struct {
	Ok        bool         `json:"ok"`
	ErrorCode int          `json:"errorCode"`
	Message   string       `json:"message"`
	Body      ResponseBody `json:"body"`
}

// ResponseBody is the body of a response belonging to Response
type ResponseBody struct {
	ValidationErrors []ValidationError `json:"validationErrors"`
	Data             interface{}       `json:"data"`
}

// ValidationError represents a form or field error
type ValidationError struct {
	Err  string `json:"error"`
	Path string `json:"path"`
}

// NewReply creates a new reply
func NewReply() Response {
	r := Response{}
	r.ErrorCode = ErrorNoError
	return r
}

// Error turns a Response into an error and then commits
func (r Response) Error(w http.ResponseWriter, message string, errorCode int) {
	r.Ok = false
	r.Message = message
	r.ErrorCode = errorCode
	r.Commit(w, nil)
}

// Success turns a Response into a successful response, and then commits
func (r Response) Success(w http.ResponseWriter, message string, data interface{}) {
	r.Ok = true
	r.Message = message
	r.ErrorCode = ErrorNoError
	r.Commit(w, data)
}

// Commit sends the Response as the HTTP body
func (r Response) Commit(w http.ResponseWriter, data interface{}) {
	r.Body.Data = data
	json.NewEncoder(w).Encode(r)
}

// AddValidationError adds a validation error to a given Response
func (r *Response) AddValidationError(path string, err string) {
	verr := ValidationError{Err: err, Path: path}
	r.Body.ValidationErrors = append(r.Body.ValidationErrors, verr)
}

// HasValidationErrors checks if a Response has any validation errors
func (r *Response) HasValidationErrors() bool {
	if len(r.Body.ValidationErrors) < 1 {
		return false
	}
	return true
}
