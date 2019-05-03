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

YCKMS takes one args and multiple flags

- Arguments :
    - The URL to the podcast RSS Feed (eg: https://feed.ausha.co/owAEhJ0qOPkb)

- Flags :
    - --last, -l : sync last show (TODO)
    - --date, -d : sync show from date --from to date --to (format MM-DD-YYYY) (TODO)

If no flag is provided, sync all (TODO)

## Licence

See LICENCE file