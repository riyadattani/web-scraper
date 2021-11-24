package main

import (
	"fmt"
	"log"
	"net/http"
	"web-scraper/pkgs/scraper"
)

func main() {
	addressScraper := scraper.New()
	addresses, err := addressScraper.Scrape("https://www.retinalscreening.co.uk/patient-information/screening-locations-list/#loclist-South")
	if err != nil {
		log.Fatalf("unable to scrape site err: %v", err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, address := range addresses {
			fmt.Fprintf(w, "%v\n", address)
		}
	})
	log.Fatal(http.ListenAndServe(":5000", handler))
}
