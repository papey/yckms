package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbumsFromCSV(t *testing.T) {
	artistExpected := "Black Bomb A"
	albumExpected := "Octantrion"

	s, err := getAlbumsFromCSV("20")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, artistExpected, s[0].artist)

	assert.Equal(t, albumExpected, s[3].album)
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

	assert.Equal(t, artistExpected, r[0].artist)

}
