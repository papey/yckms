package app

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/nfnt/resize"
	"github.com/zmb3/spotify"
)

type song struct {
	title  string
	artist string
}
type show struct {
	name     string
	playlist []song
	desc     string
	image    io.Reader
}

// createShow handles creations of show structs
func createShow(item *gofeed.Item) (*show, error) {

	// pass last show as arg, extract songs from playlist
	songs, err := parsePlaylist(item.Description)
	if err != nil {
		return nil, err
	}

	img, err := createImage(item.Image.URL)
	if err != nil {
		return nil, err
	}

	// create show stuct
	return &show{name: item.Title, playlist: songs, desc: item.Published, image: img}, nil

}

// createImage handles show image manipulation
func createImage(url string) (io.Reader, error) {

	// get image
	res, _ := http.Get(url)
	if res.StatusCode != 200 {
		return nil, errors.New("Error, can't get show image")
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	m := resize.Resize(300, 0, img, resize.Lanczos3)

	var b bytes.Buffer

	// write new image to file
	jpeg.Encode(&b, m, nil)

	return &b, err

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

// createPlaylist is used to wrap all playlist things
func createPlaylist(s *show, user string, client *spotify.Client) error {

	// create playlist
	pl, err := client.CreatePlaylistForUser(user, s.name, s.desc, true)
	if err != nil {
		return err
	}

	// image setup
	err = client.SetPlaylistImage(pl.ID, s.image)
	if err != nil {
		return err
	}

	// add songs
	err = addSongsToPlaylist(s.playlist, pl, client)
	if err != nil {
		return err
	}

	fmt.Printf("Playlist for show '%s' created, see %s\n", s.name, pl.URI)

	return nil
}

// SyncLast will sync last show
func SyncLast(url string) error {

	// get show
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return err
	}

	// create a show struct
	s, err := createShow(feed.Items[0])

	// auth to Spotify
	client, user, err := AuthToSpotify()
	if err != nil {
		return err
	}

	// create playlist
	err = createPlaylist(s, user, client)
	if err != nil {
		return err
	}

	return err

}
