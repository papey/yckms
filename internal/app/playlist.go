package app

import (
	"fmt"
	"html"
	"regexp"

	"math/rand"

	"github.com/apex/log"
	"github.com/zmb3/spotify"
)

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

	log.WithFields(log.Fields{
		"show": s.name,
		"uri":  pl.URI,
	}).Info("Playlist created")

	return nil
}

// addSongsToPlaylist loops orver songs from a show and add it to a playlist
func addSongsToPlaylist(songs []song, pl *spotify.FullPlaylist, client *spotify.Client) error {

	// array of tracks IDs
	var tracks []spotify.ID
	// search query
	var search string

	// for each songs in a show
	for _, elem := range songs {

		// if we already have ids
		if elem.id != "" {
			tracks = append(tracks, spotify.ID(elem.id))
			continue
		}

		// forge search query
		// for tracks
		if elem.album == "" {
			search = fmt.Sprintf("artist:%s track:%s", elem.artist, elem.title)
		} else if elem.title == "" {
			// for albums
			search = fmt.Sprintf("artist:%s album:%s", elem.artist, elem.album)
		}

		// Unescaped first, just to be sure
		query := purgeSearchQuery(html.UnescapeString(search))

		// search, tracks
		res, err := client.Search(query, spotify.SearchTypeTrack)
		if err != nil {
			return err
		}

		// ensure that the search return tracks
		if res.Tracks.Total > 0 {
			// if it's a track (ie: not an album) add first result
			if elem.title != "" {
				tracks = append(tracks, res.Tracks.Tracks[0].ID)
			} else if elem.album != "" {
				// if it's an album (ie: not a track) compute pseudo random number (based on tracks number inside this album)
				rand := rand.Intn(len(res.Tracks.Tracks))
				// append pseudo random track
				tracks = append(tracks, res.Tracks.Tracks[rand].ID)
			}
		}

	}

	// Add all tracks to playlist
	_, err := client.AddTracksToPlaylist(pl.ID, tracks...)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"playlist": pl.ID,
		"tracks":   tracks,
	}).Debug("Songs added to playlist")

	return nil

}

func purgeSearchQuery(query string) string {
	reg := regexp.MustCompile("['&]")
	ret := reg.ReplaceAll([]byte(query), []byte(""))

	return string(ret)
}
