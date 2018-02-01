package main

import (
	"encoding/json"
	"github.com/thaihuynhxyz/go-jwt-middleware"
	"github.com/go-martini/martini"
	"net/http"
	"gopkg.in/square/go-jose.v2"
)

func main() {

	StartServer()

}

func StartServer() {
	m := martini.Classic()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func() (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jose.HS256,
	})

	m.Get("/ping", PingHandler)
	m.Get("/secured/ping", jwtMiddleware.CheckJWT, SecuredPingHandler)

	m.Run()
}

type Response struct {
	Text string `json:"text"`
}

func respondJson(text string, w http.ResponseWriter) {
	response := Response{text}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	respondJson("All good. You don't need to be authenticated to call this", w)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	respondJson("All good. You only get this message if you're authenticated", w)
}
