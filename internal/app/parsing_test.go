package app

import (
	"testing"
)

func TestParseYCKMPlaylist(t *testing.T) {
	input := `Au programme :
- Revue de presse : Matthieu
- Chronique Fidlar : Théo
- Chronique Waste Of Space Orchestra : Eline

Playlist : Bus / I Buried Paul, Nails / Endless Resistance, Sepultura / Territory, Venom / Evilution Devilution, All Pigs Must Die / The Whip, Fidlar / Too Real, Obituary / Slowly We Rot, Wayfarer / Catcher, Waste of Space Orchestra / Seeker's Reflection, Bat / Long Live the Lewd, Witchfinder / Ouija, Gadget /Choice of a Lost Generationi`

	artistExpected := "Bus"
	songExpected := "Territory"

	s, err := parseYCKMPlaylist(input)
	if err != nil {
		t.FailNow()
	}

	if s[0].artist != artistExpected {
		t.Errorf("Expected : %s, Get : %s", artistExpected, s[0].artist)
	}

	if s[2].title != songExpected {
		t.Errorf("Expected : %s, Get : %s", songExpected, s[2].title)
	}

}
