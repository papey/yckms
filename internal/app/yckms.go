package app

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"net/http"

	"github.com/mmcdole/gofeed"
	"github.com/nfnt/resize"
	"github.com/papey/yckms/internal/spoopify"
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
	songs, err := parsePlaylist(item.ITunesExt.Summary)
	if err != nil {
		return nil, err
	}

	if songs == nil {
		fmt.Printf("Warning : no playlist to parse in show %s\n", item.Title)
		return nil, nil
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

// Sync will sync last show
func Sync(url string, last bool) error {

	var shows []*show

	// get show
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return err
	}

	// last episode only
	if last {
		// create a show struct
		s, err := createShow(feed.Items[0])
		if err != nil {
			return err
		}

		// ensure s is not nil
		if s != nil {
			// add last show to an array of one show
			shows = append(shows, s)
		}

	} else {
		// all shows, range over
		for _, e := range feed.Items {
			// create show
			s, err := createShow(e)
			if err != nil {
				return err
			}

			// append only if it's ok
			if s != nil {
				shows = append(shows, s)
			}
		}
	}

	// auth to Spotify
	client, user, err := spoopify.AuthToSpotify()
	if err != nil {
		return err
	}

	for _, elem := range shows {
		// create playlist
		err = createPlaylist(elem, user, client)
		if err != nil {
			return err
		}
	}

	return err

}
