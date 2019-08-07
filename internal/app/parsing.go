package app

import (
	"regexp"
	"strings"
)

// parse description and extract playlist, YCKM edition
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
func parseYCKMPlaylist(desc string) []song {

	var s []song

	// Split on carriage return
	split := strings.Split(desc, "\n")

	// pltf is the last element is the playlist, but not formated (1)
	if len(split) == 0 {
		return nil
	}

	plnf := split[len(split)-1]

	// remove trailing <p> and </p>
	// prepare regex
	reg, err := regexp.Compile(`(?:Playlist|PLAYLIST|Setlist) : (.+)`)
	if err != nil {
		return nil
	}

	// pl contain the string playlist (1)
	pl := reg.FindSubmatch([]byte(plnf))
	if pl == nil {
		return nil
	}

	// convert to string
	list := string(pl[1])

	// Split by ", " (2)
	songs := strings.Split(list, ", ")

	// if playlist no so long, do not add anything, it's a special episode
	// see https://podcast.ausha.co/yckm/yckm-beer-x-metal-burp-666-burp
	if len(songs) <= 1 {
		return nil
	}

	// for each song
	for _, e := range songs {
		elem := strings.Split(e, "/")
		// TRIM, just to be sure
		if len(elem) >= 2 {
			song := song{title: strings.Trim(elem[1], " "), artist: strings.Trim(elem[0], " "), album: "", id: ""}
			s = append(s, song)
		}
	}

	return s

}

// parse description and extract playlist, Le Bruit edition
// First step (1), split on \n
// Second step (2), isolate lines containing 💀 or 🐻
// Then (3), split on "|"
// (4) regex on line : Jon and Roy – Here (Ecouter (https://song.link/album/fr/i/1447292371))
// to get Artist and Song
func parseLeBruitPlaylist(desc string) []song {

	// some local vars
	var songs []song
	found := false

	// (1)
	split := strings.Split(desc, "\n")

	for _, elem := range split {
		// (2)
		if strings.Contains(elem, "💀") || strings.Contains(elem, "🐻") {
			found = true
			// (3)
			split := strings.Split(elem, "|")
			// (4)
			reg := regexp.MustCompile("[ ]+(.*) (?:-|–) (.*) \\(Ecouter.*")
			// res[1] contains Artist, res[2] contains Title
			res := reg.FindSubmatch([]byte(split[1]))
			if len(res) >= 3 {
				songs = append(songs, song{artist: string(res[1]), album: strings.Trim(string(res[2]), ""), title: "", id: ""})
			}
		}
	}

	// handle silent error, if no playlist found, just pass
	if found == false {
		return nil
	}

	return songs

}
