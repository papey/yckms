package app

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// AuthToSpotify is used to auth user to his spotify account
func AuthToSpotify() (*spotify.Client, error) {

	id, secret, err := getSpotifyFromEnv()
	if err != nil {
		return nil, err
	}

	// Create config object
	c := &clientcredentials.Config{
		ClientID:     id,
		ClientSecret: secret,
		TokenURL:     spotify.TokenURL,
	}

	// Token
	t, err := c.Token(context.Background())
	if err != nil {
		log.Fatalf("Can't get token: %v", err)
	}

	// New client, with token
	client := spotify.Authenticator{}.NewClient(t)

	return &client, err

}

// getSpotifyFromEnv is used to check env and get Spotify Auth vars
func getSpotifyFromEnv() (string, string, error) {

	id := os.Getenv("SPOTIFY_ID")
	if id == "" {
		return "", "", errors.New("SPOTIFY_ID env var not set")
	}

	secret := os.Getenv("SPOTIFY_SECRET")
	if secret == "" {
		return "", "", errors.New("SPOTIFY_SECRET env var not set")
	}

	return id, secret, nil

}
