package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type Addresses []string

type AddressScraper struct {
	Addresses Addresses
}

func New() *AddressScraper {
	return &AddressScraper{Addresses: nil}
}

func (s AddressScraper) Scrape(url string) (Addresses, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	document.Find("div.loclist-address").Each(func(i int, selection *goquery.Selection) {
		html, _ := selection.Html()
		addWithoutBr := strings.Replace(html, "<br/>", " ", -1)
		s.Addresses = append(s.Addresses, addWithoutBr)
	})

	return s.Addresses, err

}
