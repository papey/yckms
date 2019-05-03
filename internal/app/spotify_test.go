package app

import (
	"os"
	"testing"
)

func TestAuthToSpotify(t *testing.T) {

	if os.Getenv("DRONE") == "true" {
		t.Skip("Skipping test in CI environment")
	}

	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	client, user, err := AuthToSpotify()
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
