package app

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// AuthToSpotify is used to auth user to his spotify account
func AuthToSpotify(SpotifyID string, SpotifySecret string) (spotify.Client, error) {

	// Create config object
	c := &clientcredentials.Config{
		ClientID:     SpotifyID,
		ClientSecret: SpotifySecret,
		TokenURL:     spotify.TokenURL,
	}

	// Token
	t, err := c.Token(context.Background())
	if err != nil {
		log.Fatalf("Can't get token: %v", err)
	}

	// New client, with token
	client := spotify.Authenticator{}.NewClient(t)

	return client, err

}
