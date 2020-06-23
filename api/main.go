package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"encoding/json"
	"errors"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/handlers"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joleal/imperial/api/controllers"
)

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func main() {

	r := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			aud := "https://imperialapi.leal.im"
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("Invalid audience")
			}
			// Verify 'iss' claim
			iss := "https://dev-1vkyztq3.eu.auth0.com/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer")
			}

			cert, err := getPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	secureHandle := func(path string, handler http.Handler) *mux.Route {
		return r.Handle(path, jwtMiddleware.Handler(handler))
	}

	//r.HandleFunc("/info", u.HelloHandler)
	//secureHandle("/", http.HandlerFunc(u.HelloHandler))

	u := &controllers.User{}
	secureHandle("/user/register", http.HandlerFunc(u.RegisterHandler)).Methods("POST")
	secureHandle("/users", http.HandlerFunc(u.Users)).Methods("GET")

	g := &controllers.Game{}
	secureHandle("/games/active", http.HandlerFunc(g.GetActiveGames)).Methods("GET")
	secureHandle("/games/open", http.HandlerFunc(g.GetOpenGames)).Methods("GET")
	secureHandle("/game", http.HandlerFunc(g.CreateGame)).Methods("POST")

	log.Println("Listing for requests at http://localhost:3333/")
	log.Fatal(http.ListenAndServe(":3333", corsWrapper.Handler(r)))

}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://dev-1vkyztq3.eu.auth0.com/.well-known/jwks.json")

	if err != nil {
		log.Println("couldn't get resp")
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	//err = json.NewDecoder(resp.Body).Decode(&jwks)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &jwks)
	if err != nil {
		log.Println("couldn't decode body")
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
