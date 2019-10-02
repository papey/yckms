package app

import (
	"testing"
)

func TestCreateImage(t *testing.T) {

	url := "https://image.ausha.co/kRalftnUIhCs8GKLARlcEvQjtKgNIEFeg88CxDdc_400x400.jpeg"

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
