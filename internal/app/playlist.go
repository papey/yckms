package app

import (
	"fmt"

	"math/rand"

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

	fmt.Printf("Info: playlist for show '%s' created, see %s\n", s.name, pl.URI)

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

		// forge search query
		// for tracks
		if elem.album == "" {
			search = fmt.Sprintf("artist:%s track:%s", elem.artist, elem.title)
		} else if elem.title == "" {
			// for albums
			search = fmt.Sprintf("artist:%s album:%s", elem.artist, elem.album)
		}

		// search, tracks
		res, err := client.Search(search, spotify.SearchTypeTrack)
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
				rand := rand.Intn(res.Tracks.Total)
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

	return nil

}
