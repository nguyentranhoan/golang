package Media

import (
	"encoding/json"
	"net/http"
)

// Structure response
type Response struct {
	Message string    `json:"message"`
	Error   [] string `json:"error"`
}

var file_name = Gen_uuid()

func response(w http.ResponseWriter, message string, err error){

	if err != nil {
		var res = Response{Message: message, Error: []string{err.Error()}}
		http.Error(w, "", http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}
	var res = Response{Message: message, Error: nil}
	http.Error(w, "", http.StatusBadRequest)
	json.NewEncoder(w).Encode(res)
	return
}
