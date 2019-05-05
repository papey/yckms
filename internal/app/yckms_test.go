package app

import (
	"testing"
)

func TestCreateImage(t *testing.T) {

	url := "https://image.ausha.co/SHSw9XAonLSu3xLyQBpTOT1ai6bpGCw8LKceHuan_1400x1400.jpeg?t=1556571464"

	_, err := createImage(url)
	if err != nil {
		t.Error("Error: Can't create image")
	}

}

func TestParseDates(t *testing.T) {

	from := "2019-01-11"
	to := "2019-04-01"

	_, err := parseDates(from, to)
	if err != nil {
		t.Error(err)
	}

}
