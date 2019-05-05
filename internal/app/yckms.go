package app

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/araddon/dateparse"
	"github.com/mmcdole/gofeed"
	"github.com/nfnt/resize"
	"github.com/papey/yckms/internal/spoopify"
)

// Song struct
type song struct {
	// song title
	title string
	// song artist
	artist string
}
type show struct {
	// podcast title/name
	name string
	// array of songs played during show
	playlist []song
	// small description used as Spotify playlist description
	desc string
	// image used as Spotify playlist image
	image io.Reader
}

type interval struct {
	from time.Time
	to   time.Time
}

// filterFeed apply date filter on a feed
func filterFeed(feed *gofeed.Feed, from string, to string) ([]*gofeed.Item, error) {

	var filtered []*gofeed.Item

	d, err := parseDates(from, to)
	if err != nil {
		return nil, err
	}

	for _, e := range feed.Items {
		t, err := dateparse.ParseAny(e.Published)
		if err != nil {
			return nil, err
		}

		if t.Before(d.to) && t.After(d.from) {
			filtered = append(filtered, e)
		}

	}

	return filtered, err
}

// parseDates takes from and to dates as string and convert them to date struct
func parseDates(from string, to string) (*interval, error) {

	format := "2006-01-02"

	f, err := time.Parse(format, from)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse(format, to)
	if err != nil {
		return nil, err
	}

	if f.After(t) {
		return nil, fmt.Errorf("Error: %s (from) is after %s (to)", from, to)
	}

	return &interval{from: f, to: t}, nil

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

// createShows wrap stuff to create and filter show structs
func createShows(feed *gofeed.Feed, last bool, from string, to string) ([]*show, error) {

	// local vars
	var shows []*show
	var items []*gofeed.Item
	var err error

	// if last, items contains only the last show
	if last {
		items = append(items, feed.Items[0])
	} else {
		// if last is false, fallback to all
		// remove if dates set, filter
		if from != "" && to != "" {
			items, err = filterFeed(feed, from, to)
			if err != nil {
				return nil, err
			}
		} else {
			// else, get all items
			items = feed.Items
		}

	}

	// all shows, range over
	for _, e := range items {
		// create show
		s, err := createShow(e)
		if err != nil {
			return nil, err
		}

		// append only if it's ok
		if s != nil {
			shows = append(shows, s)
		}
	}

	return shows, nil

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

// Sync is the most important function of the app
// takes url, and params from cli and exec :
// - get RSS feed
// - filter
// - Spotify auth
// - create playlist(s)
func Sync(url string, last bool, from string, to string) error {

	// show episodes, YCKMS format
	var shows []*show

	// get show
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return err
	}

	shows, err = createShows(feed, last, from, to)
	if err != nil {
		return err
	}

	// auth to Spotify
	client, user, err := spoopify.AuthToSpotify()
	if err != nil {
		return err
	}

	// for all shows
	for _, elem := range shows {
		// create playlist
		err = createPlaylist(elem, user, client)
		if err != nil {
			return err
		}
	}

	return err

}
