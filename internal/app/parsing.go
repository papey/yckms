package app

import (
	"regexp"
	"strings"

	"github.com/apex/log"
)

// Parser interface
type Parser interface {
	// parse is used to get list of song from show
	parse() []song
}

// YCKMOrSaccage is used for show YCKM
type YCKMOrSaccage struct {
	// podcast name
	name string
	// podcast title
	title string
	// podcast description
	desc string
}

// NewYCKMOrSaccage inits a YCKM or Saccage struct since the share the same format
func NewYCKMOrSaccage(name string, title string, desc string) YCKMOrSaccage {
	return YCKMOrSaccage{name: name, title: title, desc: desc}
}

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
func (y YCKMOrSaccage) parse() []song {

	// is playlist found ?
	var found bool = false

	// playlist
	var plst string = ""

	var s []song

	// Split on carriage return
	split := strings.Split(y.desc, "\n\n")

	// ensure split is ok
	if len(split) == 0 {
		return nil
	}

	reg, err := regexp.Compile(`(Playlist|PLAYLIST|Setlist)\s?:\s?(.+)?`)
	// preprare regex
	if err != nil {
		return nil
	}

	// loop over lines
	for i, elem := range split {

		// Try to find something that looks like a playlist
		pl := reg.FindSubmatch([]byte(elem))
		// If something is found
		if len(pl) >= 1 {
			// Check length of submatch array
			if len(pl) == 3 {
				// If the second goup is empty
				if strings.TrimSpace(string(pl[2])) == "" {
					// Playlist is the next line
					plst = split[i+1]
				} else {
					// Else, playlist is the last member of match array
					plst = string(pl[2])
				}
			}
			// Yay, we found something
			found = true
			break
		}

	}

	if !found {
		return nil
	}

	// convert to string
	list := string(plst)

	// Split by ", " (2)
	songs := strings.Split(list, ", ")
	if len(songs) <= 1 {
		songs = strings.Split(list, ";")
	}

	// if playlist no so long, do not add anything, it's a special episode
	// see https://podcast.ausha.co/yckm/yckm-beer-x-metal-burp-666-burp
	if len(songs) <= 1 {
		return nil
	}

	// for each song
	for _, e := range songs {
		elem := strings.Split(e, "/")
		if len(elem) == 1 {
			elem = strings.Split(e, " - ")
		}
		// TRIM, just to be sure
		if len(elem) >= 2 {
			song := song{title: strings.Trim(elem[1], " "), artist: strings.Trim(elem[0], " "), album: "", id: ""}
			s = append(s, song)
		}
	}

	return s

}

// LB is used for show Le Bruit
type LB struct {
	// podcast name
	name string
	// podcast title
	title string
	// podcast description
	desc string
}

// NewLB inits a LB struct
func NewLB(name string, title string, desc string) LB {
	return LB{name: name, title: title, desc: desc}
}

// parse description and extract playlist, Le Bruit edition
// First step (1), split on \n
// Second step (2), isolate lines containing 💀 or 🐻
// (3) regex on line to get Artist and Song
func (l LB) parse() []song {

	// some local vars
	var songs []song
	found := false

	// (1)
	split := strings.Split(l.desc, "\n")

	for _, elem := range split {
		// (2)
		if strings.Contains(elem, "💀") || strings.Contains(elem, "🐻") {
			// most complex to simple one
			regexs := []string{
				`\s*(?:💀|🐻) \(\d+:\d+:\d+\) (.*) - (.*) (\(\d+\) )?\([Ecouter|\d+].*`,
				`\s*(?:💀|🐻) \(\d+:\d+:\d+\) \d+\. (.*) - (.*)`,
				`\s*(?:💀|🐻)(.*) - (.*)`,
			}
			found = true
			for _, reg := range regexs {
				r := regexp.MustCompile(reg)
				res := r.FindSubmatch([]byte(elem))
				if len(res) >= 2 {
					songs = append(songs, song{artist: string(res[1]), album: strings.Trim(string(res[2]), ""), title: "", id: ""})
					break
				}
			}
		}
	}

	// handle silent error, if no playlist found, just pass
	if found == false {
		return nil
	}

	return songs

}

// HC show is used for show Harry Cover
type HC struct {
	// podcast name
	name string
	// podcast title
	title string
	// podcast description
	desc string
}

// NewHC inits a HC struct
func NewHC(name string, title string, desc string) HC {
	return HC{name: name, title: title, desc: desc}
}

// parse description and extract playlist, HarryCover edition
// First step (1), split on \n
// Second step (2), isolate lines containing original version and cover version
// Third step (3), get Spotify ID from URL
func (hc HC) parse() []song {

	// some local vars
	var songs []song
	found := false

	// (1)
	split := strings.Split(hc.desc, "\n")

	for _, elem := range split {
		// (2)
		if strings.Contains(elem, "https://open.spotify.com/track") {
			found = true
			// (3)
			reg := regexp.MustCompile(`.*https://open.spotify.com/track/(\w+)?.*`)
			res := reg.FindSubmatch([]byte(elem))
			songs = append(songs, song{artist: "", album: "", title: "", id: string(res[1])})
		}
	}

	// handle silent error, if no playlist found, just pass
	if found == false {
		return nil
	}

	return songs

}

// Pifo is used for La Pifothèque show
type Pifo struct {
	// podcast name
	name string
	// podcast title
	title string
	// podcast description
	desc string
}

// NewPifo inits a HC struct
func NewPifo(name string, title string, desc string) Pifo {
	return Pifo{name: name, title: title, desc: desc}
}

// parse description and extract playlist, Le Pifothèque edition
// (1) extract epifode number
// (2) get albums from CSV
func (p Pifo) parse() []song {

	// Local var containing epifode number
	var epifode string

	reg := regexp.MustCompile(`.*Epifode (\d+)`)

	res := reg.FindSubmatch([]byte(p.title))

	// Little hack to get the first epifode
	if len(res) >= 2 {
		epifode = string(res[1])
	} else {
		return nil
	}

	// (2)
	s, err := getAlbumsFromCSV(epifode)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return s
}

// InitParse inits a parser using it's name
func InitParse(name, title, desc string) Parser {
	switch name {
	case "YCKM":
		return NewYCKMOrSaccage(name, title, formatAushaDesc(desc))
	case "SACCAGE":
		return NewYCKMOrSaccage(name, title, formatAushaDesc(desc))
	case "Le Bruit":
		return NewLB(name, title, formatAushaDesc(desc))
	case "Recoversion, le Podcast des Meilleures Reprises":
		return NewHC(name, title, desc)
	case "La Pifothèque":
		return NewPifo(name, title, formatAushaDesc(desc))
	default:
		return nil
	}
}

// formatAushaDesc will take an html input, remplace </p> with \n, and smash the rest
// 🖕 Ausha
func formatAushaDesc(in string) (out string) {
	cleaner := regexp.MustCompile("<[^>]*>")
	return string(cleaner.ReplaceAll([]byte(strings.ReplaceAll(in, "</p>", "\n")), []byte("")))
}
