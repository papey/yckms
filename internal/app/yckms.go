package app

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

// SyncLast will sync last show
func SyncLast(url string) error {

	// get show
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	fmt.Println(feed.Title)

	return nil
}
