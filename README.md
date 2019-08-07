# YCKMS

[![Build Status](https://drone.github.papey.fr/api/badges/papey/yckms/status.svg)](https://drone.github.papey.fr/papey/yckms)

Small go app to sync playlists from various
french metal podcast shows to Spotify.

YCKMS currently supports :

- [YCKM](https://podcast.ausha.co/yckm)
- [Le Bruit](https://podcast.ausha.co/le-bruit)

## Getting Started

### Prerequisites

- [Golang](https://golang.org)

### Installing

Get deps using go mod

```sh
go mod vendor
```

then buld

```sh
cd cmd
go build ykcms.go
```

### Usage

#### Setup

##### Oauth

1. Create a YCKMS application on [Spotify](https://developer.spotify.com/my-applications/) : https://developer.spotify.com/my-applications/
2. Set `http://localhost:8080` as callback url (in order to use a different port see env var section bellow)
3. Copy SPOTIFY_ID and SPOTIFY_SECRET

#### Args, env and flags

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

#### Example

```sh
# Launch
SPOTIFY_ID=XXXX SPOTIFY_SECRET=YYYY ./yckms -l https://feed.ausha.co/owAEhJ0qOPkb
# Waiting for oauth stuff, go to url to oauth
Please log in to Spotify by visiting the following page : https://accounts.spotify.com/authorize?client_id=blablablalbalablbala
# Accept access to user profile and data
# Ok, everything is fine
You are logged in as : morty
```

## Notes

Le Bruit podcast is focused on album rather than songs (like YCKM). Each
playlist generated for Le Bruit is a collection of pseudo random tracks of
albums discussed in the show episode.

## Running the tests

```sh
go test github.com/papey/yckms/internal/app
```

### Continuous Integration

See [drone.github.papey.fr/papey/yckms](https://drone.github.papey.fr/papey/yckms/)

## Built With

- [go-spotify](https://github.com/Krognol/go-spotify) - Spotify web API for go

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

- **Wilfried OLLIVIER** - _Main author_ - [Papey](https://github.com/papey)

## License

[LICENSE](LICENSE) file for details

## Acknowledgments

- Thanks Le Bruit and YCKM for all you podcasts !