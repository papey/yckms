package app

import (
	"os"
	"testing"
)

func TestAuthToSpotify(t *testing.T) {

	id := os.Getenv("SPOTIFY_ID")
	if id == "" {
		t.Fatal("Error SPOTIFY_ID env var not set")
	}

	secret := os.Getenv("SPOTIFY_SECRET")
	if secret == "" {
		t.Fatal("Error SPOTIFY_SECRET env var not set")
	}

	user := os.Getenv("SPOTIFY_USER")
	if user == "" {
		t.Fatal("Error SPOTIFY_USER env var not set")
	}

	client, err := AuthToSpotify(id, secret)
	if err != nil {
		t.Fatal(err)
	}

	page, err := client.GetPlaylistsForUser(user)
	if err != nil {
		t.Fatal(err)
	}

	for _, elem := range page.Playlists {
		t.Logf("Playlist name: %s", elem.Name)
	}

}
