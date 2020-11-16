package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func Handler(r chi.Router) {

}

func Signin(w http.ResponseWriter, r *http.Request) {
	var request User
	json.NewDecoder(r.Body).Decode(&request)
	service := NewUser()
	token, err := service.Signin(request)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(map[string]string{"token": token})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
