# YCKMS

[![Build Status](https://drone.github.papey.fr/api/badges/papey/yckms/status.svg)](https://drone.github.papey.fr/papey/yckms)

Small go app to sync playlists from various
french metal podcast shows to Spotify.

YCKMS currently supports :

- [YCKM](https://podcast.ausha.co/yckm)
- [Le Bruit](https://podcast.ausha.co/le-bruit)

## Build

Get deps using go mod

    go mod vendor

Build

    cd cmd
    go build ykcms.go

## Usage

### Setup

#### Oauth

1. Create a YCKMS application on [Spotify](https://developer.spotify.com/my-applications/) : https://developer.spotify.com/my-applications/
2. Set `http://localhost:8080` as callback url (in order to use a different port see env var section bellow)
3. Copy SPOTIFY_ID and SPOTIFY_SECRET

### Args, env and flags

YCKMS takes one arg, multiple env var and multiple flags

- Env vars :
    - SPOTIFY_ID : from oauth setup (3)
    - SPOTIFY_SECRET : from oauth setup (3)
    - HTTP_CALLBACK_PORT : change http callback port (default 8080)

- Arguments :
    - The URL to the podcast [RSS Feed](https://feed.ausha.co/owAEhJ0qOPkb) (eg: https://feed.ausha.co/owAEhJ0qOPkb)

- Flags :
    - --last, -l : sync last show
    - --date, -d : sync show from date --from to date --to (format YYYY-MM-DD)

If no flag is provided, sync all

### Usage example

    # Launch
    SPOTIFY_ID=XXXX SPOTIFY_SECRET=YYYY ./yckms -l https://feed.ausha.co/owAEhJ0qOPkb
    # Waiting for oauth stuff, go to url to oauth
    Please log in to Spotify by visiting the following page : https://accounts.spotify.com/authorize?client_id=blablablalbalablbala
    # Accept access to user profile and data
    # Ok, everything is fine
    You are logged in as : morty

## Notes

Le Bruit podcast is focused on album rather than songs (like YCKM). Each
playlist generated for Le Bruit is a collection of pseudo random tracks of
albums discussed in the show episode.

## Licence

See LICENCE file

## Build info

See [https://drone.github.papey.fr/papey/yckms/](https://drone.github.papey.fr/papey/yckms/)