package app

import (
	"fmt"

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

	// for each songs in a show
	for _, elem := range songs {
		// forge search query
		search := fmt.Sprintf("artist:%s track:%s", elem.artist, elem.title)

		// search
		res, err := client.Search(search, spotify.SearchTypeTrack)
		if err != nil {
			return err
		}

		// search query is specific, so first item is always the best
		if res.Tracks.Total > 0 {
			// add to array
			tracks = append(tracks, res.Tracks.Tracks[0].ID)
		}

	}

	// Add all tracks to playlist
	_, err := client.AddTracksToPlaylist(pl.ID, tracks...)
	if err != nil {
		return err
	}

	return nil

}
