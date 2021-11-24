package scraper_test

import (
	"github.com/matryer/is"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-scraper/pkgs/scraper"
)

func TestScrape(t *testing.T) {
	t.Run("given some html, when it is scraped, then I get a list of addresses", func(t *testing.T) {
		is := is.New(t)

		content, err := ioutil.ReadFile("test.html")
		is.NoErr(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write(content)
			is.NoErr(err)
		}))

		addressScraper := scraper.New()
		addresses, err := addressScraper.Scrape(server.URL)
		is.NoErr(err)

		is.Equal(len(addresses), 3)
		is.Equal(addresses[0], "336 Lichfield Road Mere Green Sutton Coldfield B74 4BH")
	})
}
