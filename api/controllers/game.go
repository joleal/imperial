package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/joleal/imperial/api/models"
	"io/ioutil"
	"net/http"
)

//Game contains the Game controller
type Game struct {
	ID string
}

//GetOpenGames gets active games for logged user
func (g *Game) GetOpenGames(w http.ResponseWriter, r *http.Request) {
	// user
	games, err := models.GetOpenGames()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonResponse))
	w.Header().Set("Content-Type", "application/json")

}

//GetActiveGames gets active games for logged user
func (g *Game) GetActiveGames(w http.ResponseWriter, r *http.Request) {
	// user
	user := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	games, err := models.GetActiveGames(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonResponse))
	w.Header().Set("Content-Type", "application/json")

}

//Creates a game in the database
func (g *Game) CreateGame(w http.ResponseWriter, r *http.Request) {
	// user
	var game models.SaveGame

	//Read all the data in r.Body from a byte[], convert it to a string, and assign store it in 's'.
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
	}

	// use the built in Unmarshal function to put the string we got above into the empty page we created at the top.  Notice the &p.  The & is important, if you don't understand it go and do the 'Tour of Go' again.
	err = json.Unmarshal(s, &game)
	if err != nil {
		panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
	}

	err = game.CreateGame()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "ok"}`))

}
