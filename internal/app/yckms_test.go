package app

import (
	"testing"
)

func TestParsePlaylist(t *testing.T) {
	input := `<p>Au programme :</p>
<p>- Revue de presse : Matthieu</p>
<p>- Chronique Fidlar : Théo</p>
<p>- Chronique Waste Of Space Orchestra : Eline</p>
<p><br></p>
<p>Playlist : Bus / I Buried Paul, Nails / Endless Resistance, Sepultura / Territory, Venom / Evilution Devilution, All Pigs Must Die / The Whip, Fidlar / Too Real, Obituary / Slowly We Rot, Wayfarer / Catcher, Waste of Space Orchestra / Seeker's Reflection, Bat / Long Live the Lewd, Witchfinder / Ouija, Gadget /Choice of a Lost Generation</p>`

	s, err := parsePlaylist(input)
	if err != nil {
		t.Fail()
	}

	if s[0].artist != "Bus" {
		t.Errorf("Expected : Bus, Get : %s", s[0].artist)
		t.Fail()
	}

	if s[2].title != "Territory" {
		t.Errorf("Expected : Territory, Get : %s", s[2].title)
		t.Fail()
	}

}
