package app

import (
	"testing"
)

func TestGetAlbumsFromCSV(t *testing.T) {
	artistExpected := "Black Bomb A"
	albumExpected := "Octantrion"

	s, err := getAlbumsFromCSV("20")
	if err != nil {
		t.Fail()
	}

	if s[0].artist != artistExpected {
		t.Errorf("Expected : %s, Get : %s", artistExpected, s[0].artist)
	}

	if s[3].album != albumExpected {
		t.Errorf("Expected : %s, Get : %s", albumExpected, s[3].album)
	}
}

func TestGet(t *testing.T) {

	artistExpected := "Midnight oil"

	c, err := download()
	if err != nil {
		t.Fail()
	}

	d, err := read(c)
	if err != nil {
		t.Fail()
	}

	r, err := get(d, "2")

	if r[0].artist != artistExpected {
		t.Errorf("Expected : %s, Get : %s", r[0].artist, artistExpected)
	}

}
