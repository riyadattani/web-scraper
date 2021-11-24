package main

import (
	"fmt"
	"log"
	"net/http"
	"web-scraper/pkgs/scraper"
)

func main() {
	addresses, err := scraper.Scrape("https://www.retinalscreening.co.uk/patient-information/screening-locations-list/#loclist-South")
	if err != nil {
		log.Fatalf("unable to scrape site err: %v", err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "addresses: %v", addresses)
	})
	log.Fatal(http.ListenAndServe(":5000", handler))
}
