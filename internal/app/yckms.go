package app

import (
	"errors"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
)

type song struct {
	title  string
	artist string
}
type show struct {
	name     string
	playlist []song
}

// parse description and extract playlist
// Input exemple :
// 		<p>Au programme :</p>
// 		<p>- Revue de presse : Matthieu</p>
// 		<p>- Chronique Fidlar : Théo</p>
// 		<p>- Chronique Waste Of Space Orchestra : Eline</p>
// 		<p><br></p>
// 		<p>Playlist : Bus / I Buried Paul, Nails / Endless Resistance, Sepultura /
// 		Territory, Venom / Evilution Devilution, All Pigs Must Die / The Whip, Fidlar
// 		/ Too Real, Obituary / Slowly We Rot, Wayfarer / Catcher, Waste of Space
// 		Orchestra / Seeker's Reflection, Bat / Long Live the Lewd, Witchfinder /
// 		Ouija, Gadget /Choice of a Lost Generation</p>
// First step (1) :
// 		Bus / I Buried Paul, Nails / Endless Resistance, Sepultura /
// 		Territory, Venom / Evilution Devilution, All Pigs Must Die / The Whip, Fidlar
// 		/ Too Real, Obituary / Slowly We Rot, Wayfarer / Catcher, Waste of Space
// 		Orchestra / Seeker's Reflection, Bat / Long Live the Lewd, Witchfinder /
// 		Ouija, Gadget /Choice of a Lost Generation
// Second step (2) :
// 		An array containing each combo Artist / Song
// Third step (3) :
//		A song object
func parsePlaylist(desc string) ([]song, error) {

	var s []song

	// Split on carriage return
	split := strings.Split(desc, "\n")

	// pltf is the last element is the playlist, but not formated (1)
	plnf := split[len(split)-1]

	// remove trailing <p> and </p>
	// prepare regex
	reg, err := regexp.Compile(`<p>Playlist : (.+)</p>`)
	if err != nil {
		return nil, err
	}

	// pl contain the string playlist (1)
	pl := reg.FindSubmatch([]byte(plnf))
	if pl == nil {
		return nil, errors.New("No match found in show description")
	}

	// convert to string
	list := string(pl[1])

	// Split by ", " (2)
	songs := strings.Split(list, ", ")

	// for each song
	for _, e := range songs {
		elem := strings.Split(e, "/")
		// TRIM, just to be sure
		song := song{title: strings.Trim(elem[1], " "), artist: strings.Trim(elem[0], " ")}
		s = append(s, song)
	}

	return s, err

}

// SyncLast will sync last show
func SyncLast(url string) error {

	// get show
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return err
	}

	// pass last show as arg, extract songs from playlist
	songs, err := parsePlaylist(feed.Items[0].Description)
	if err != nil {
		return err
	}

	// create show stuct
	_ = show{name: feed.Title, playlist: songs}

	// auth to Spotify
	_, err = AuthToSpotify()

	return err
}
