# YCKMS

Small go app to sync playlists from [YCKM](https://podcast.ausha.co/yckm)
french metal podcast show to Spotify

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
2. Set `http://localhost:8080` as callback url
3. Copy SPOTIFY_ID and SPOTIFY_SECRET

### Args, env and flags

YCKMS takes one arg, multiple env var and multiple flags

- Env vars :
    - SPOTIFY_ID : from oauth setup (3)
    - SPOTIFY_SECRET : from oauth setup (3)

- Arguments :
    - The URL to the podcast [RSS Feed](https://feed.ausha.co/owAEhJ0qOPkb) (eg: https://feed.ausha.co/owAEhJ0qOPkb)

- Flags :
    - --last, -l : sync last show
    - --date, -d : sync show from date --from to date --to (format MM-DD-YYYY) (TODO)

If no flag is provided, sync all (TODO)

## Licence

See LICENCE file