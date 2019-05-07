package app

import (
	"regexp"
	"strings"
)

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
	if len(split) == 0 {
		return nil, nil
	}

	plnf := split[len(split)-1]

	// remove trailing <p> and </p>
	// prepare regex
	reg, err := regexp.Compile(`(?:Playlist|PLAYLIST|Setlist) : (.+)`)
	if err != nil {
		return nil, err
	}

	// pl contain the string playlist (1)
	pl := reg.FindSubmatch([]byte(plnf))
	if pl == nil {
		return nil, nil
	}

	// convert to string
	list := string(pl[1])

	// Split by ", " (2)
	songs := strings.Split(list, ", ")

	// if playlist no so long, do not add anything, it's a special episode
	// see https://podcast.ausha.co/yckm/yckm-beer-x-metal-burp-666-burp
	if len(songs) <= 1 {
		return nil, nil
	}

	// for each song
	for _, e := range songs {
		elem := strings.Split(e, "/")
		// TRIM, just to be sure
		if len(elem) >= 2 {
			song := song{title: strings.Trim(elem[1], " "), artist: strings.Trim(elem[0], " ")}
			s = append(s, song)
		}
	}

	return s, err

}
