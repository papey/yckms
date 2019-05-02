package app

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// Auth contains everything needed to auth on spotify
type auth struct {
	secret string
	id     string
	user   string
}

// AuthToSpotify is used to auth user to his spotify account
func AuthToSpotify() (*spotify.Client, string, error) {

	creds, err := getSpotifyFromEnv()
	if err != nil {
		return nil, "", err
	}

	// Create config object
	c := &clientcredentials.Config{
		ClientID:     creds.id,
		ClientSecret: creds.secret,
		TokenURL:     spotify.TokenURL,
	}

	// Token
	t, err := c.Token(context.Background())
	if err != nil {
		log.Fatalf("Can't get token: %v", err)
	}

	// New client, with token
	client := spotify.Authenticator{}.NewClient(t)

	return &client, creds.user, err

}

// getSpotifyFromEnv is used to check env and get Spotify Auth vars
func getSpotifyFromEnv() (*auth, error) {

	id := os.Getenv("SPOTIFY_ID")
	if id == "" {
		return nil, errors.New("SPOTIFY_ID env var not set")
	}

	secret := os.Getenv("SPOTIFY_SECRET")
	if secret == "" {
		return nil, errors.New("SPOTIFY_SECRET env var not set")
	}

	user := os.Getenv("SPOTIFY_USER")
	if user == "" {
		return nil, errors.New("SPOTIFY_USER env var not set")
	}

	return &auth{id: id, secret: secret, user: user}, nil

}
