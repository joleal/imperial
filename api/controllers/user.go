package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/joleal/imperial/api/models"
	"net/http"
)

//User contains the User controller
type User struct {
	ID string
}

//RegisterHandler registers a user in the database
func (u *User) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Declare a new user to be saved
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.ID = r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)

	err = (&newUser).RegisterUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "ok"}`))
}

//HelloHandler handles hello world request
func (u *User) HelloHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
	w.WriteHeader(http.StatusOK)
	/*
		user := r.Context().Value("user")
		// message := fmt.Sprintf("This is an authenticated request")
		// message += fmt.Sprintf("Claim content:\n")
		// for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		// 	message += fmt.Sprintf("%s :\t%#v\n", k, v)
		// }

		jsonResponse, err := json.Marshal(user.(*jwt.Token).Claims.(jwt.MapClaims)["sub"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		message := jsonResponse
		w.Write([]byte(message))
	*/
}
