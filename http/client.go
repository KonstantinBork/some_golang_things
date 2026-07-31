package http

import (
	"io"
	"log"
	"net/http"
)

func Get(requestURL string) string {
	resp, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
