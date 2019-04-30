package app

import (
	"os"
	"testing"
)

func TestAuthToSpotify(t *testing.T) {

	user := os.Getenv("SPOTIFY_USER")
	if user == "" {
		t.Fatal("Error SPOTIFY_USER env var not set")
	}

	client, err := AuthToSpotify()
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
