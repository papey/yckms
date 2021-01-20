package app

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/apex/log"
	"github.com/araddon/dateparse"
	"github.com/mmcdole/gofeed"
	"github.com/nfnt/resize"
	"github.com/papey/yckms/internal/spoopify"
)

// Song struct
type song struct {
	// song title
	title string
	// album
	album string
	// song artist
	artist string
	// song id
	id string
}

// Show struct
type show struct {
	// podcast title/name
	name string
	// array of songs played during show
	playlist []song
	// small description used as Spotify playlist description
	desc string
	// image used as Spotify playlist image
	image io.Reader
	// generator id
	genid int
}

// create a specific type for a list of show
type shows []*show

// implement sort interface using genid as compare data
func (s shows) Len() int           { return len(s) }
func (s shows) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s shows) Less(i, j int) bool { return s[i].genid < s[j].genid }

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

// showResult is used as a value to get result from createShow go routine
type showResult struct {
	// show structure
	s *show
	// warning
	w string
	// error
	e error
}

// createShow handles creations of show structs
func createShow(item *gofeed.Item, name string, id int, pipe chan showResult) {

	// local vars
	var songs []song
	var err error

	parser := InitParse(name, item.Title, item.Description)
	if parser == nil {
		log.Fatal("Show not supported")
	}

	songs = parser.parse()

	if songs == nil {
		// not an error, but no show created, send a warning
		pipe <- showResult{s: nil, e: nil, w: fmt.Sprintf("Warning: no playlist to parse in show %s", item.Title)}
		return
	}

	// get image
	img, err := createImage(item.Image.URL)
	if err != nil {
		// if something goes wrong, send an error
		pipe <- showResult{s: nil, e: fmt.Errorf("Error: can't create show image for show %s", item.Title), w: ""}
		return
	}

	// send actual show playlist and metadata
	pipe <- showResult{s: &show{name: item.Title, playlist: songs, desc: item.Published, image: img, genid: id}, e: nil, w: ""}
}

// createShows wrap stuff to create and filter show structs
func createShows(feed *gofeed.Feed, last bool, from string, to string) ([]*show, error) {

	// local vars
	var s []*show
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

	// create pipe used to return value from goroutine
	pipe := make(chan showResult, 1)
	// ensure pipe is close at the end of the func
	defer close(pipe)

	// get playlists for all shows, using concurrency
	for i, e := range items {
		// create show
		go createShow(e, feed.Title, i, pipe)
	}

	// get return from all goroutines, ensure no deadlock
	for range items {
		// get return value
		ret := <-pipe
		// if it's an error,
		if ret.e != nil {
			// return it to caller
			return nil, ret.e
		}

		// add show to current array
		if ret.s != nil {
			s = append(s, ret.s)
		}

		// if a warning is set, print it
		if ret.w != "" {
			log.Warn(ret.w)
		}
	}

	// short shows
	sort.Sort(shows(s))

	// return sorted shows
	return s, nil

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
		log.Fatal(err.Error())
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
