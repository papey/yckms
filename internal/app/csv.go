package app

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// getAlbumsFromCSV is used to get the 4 albums from a specified epifode
func getAlbumsFromCSV(epifode string) (s []song, err error) {

	c, err := download()
	if err != nil {
		return nil, err
	}

	r, err := read(c)
	if err != nil {
		return nil, err
	}

	return get(r, epifode)

}

// download is used to get the CSV file listing albums from La Pifothèque
func download() (c []byte, err error) {
	// url
	url := "https://docs.google.com/spreadsheets/d/e/2PACX-1vS9ET46agcI-kxiAhpnUPMqer0x7WQ9zAb-ZMU8_vdN8PBpPpIjaudspd-Mb_FFBzLLlqHAVsaHe5q-/pub?output=csv"

	// client setup + timeout
	client := http.Client{
		Timeout: time.Duration(5 * int64(time.Second)),
	}

	// Get
	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.New("Cannot get file for given URL")
	}

	// Ensure it's ok
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status code is %d", resp.StatusCode)
	}

	// Ensure content type is ok
	if resp.Header["Content-Type"][0] != "text/csv" {
		return nil, errors.New("Content type is not text/csv")
	}

	// Read content from file
	c, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading body from CSV file")
	}

	// We just need it in memory, no need to write it into a file
	return c, err

}

// read is used to extract all records from CSV inside a matrix
func read(c []byte) (r [][]string, err error) {

	// Setup a reader
	reader := csv.NewReader(bytes.NewReader(c))

	// Read the entire file
	r, err = reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return r, err
}

// get is used to extract 4 specific albums from the readed CSV file
func get(r [][]string, epifode string) (a []song, err error) {

	for i, line := range r {
		if len(line) == 4 {
			if line[3] == epifode {
				a = append(a, song{artist: r[i][0], album: r[i][1], title: "", id: ""})
			}
		}
	}

	return a, err

}
