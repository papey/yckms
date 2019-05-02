package app

import (
	"testing"
)

func TestAuthToSpotify(t *testing.T) {

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
