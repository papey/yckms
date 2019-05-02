package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
)

// big thanks to
// https://github.com/zmb3/spotify/blob/master/examples/authenticate/authcode/authenticate.go

// callback uri
const uri = "http://localhost:8080/callback"

// Global vars user to auth
var (
	auth  = spotify.NewAuthenticator(uri, spotify.ScopeImageUpload, spotify.ScopePlaylistModifyPublic)
	ch    = make(chan *spotify.Client)
	state = "yckms"
)

// AuthToSpotify is used to auth user to his spotify account
func AuthToSpotify() (*spotify.Client, string, error) {

	err := ensureSpotifyCreds()
	if err != nil {
		return nil, "", err
	}

	// Create a simple http server
	http.HandleFunc("/callback", complete)
	// only respond on /callback
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Burp !")
	})
	// serve
	go http.ListenAndServe(":8080", nil)

	// try auth
	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page :", url)

	// wait for auth to complete
	client := <-ch

	// check if it's ok
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as :", user.ID)

	return client, user.ID, err

}

// complete auth is exec when /callback is call
func complete(w http.ResponseWriter, r *http.Request) {

	// get token
	token, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	// Get state
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// Token is used to create a new client
	client := auth.NewClient(token)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client

}

// getSpotifyFromEnv is used to check env and get Spotify Auth vars
func ensureSpotifyCreds() error {

	if os.Getenv("SPOTIFY_ID") == "" {
		return errors.New("SPOTIFY_ID env var not set")
	}

	if os.Getenv("SPOTIFY_SECRET") == "" {
		return errors.New("SPOTIFY_SECRET env var not set")
	}

	return nil

}
