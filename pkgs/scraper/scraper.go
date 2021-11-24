package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

type Addresses []string

type AddressScraper struct {
	Addresses Addresses
}

func New() *AddressScraper {
	return &AddressScraper{Addresses: []string{}}
}

func (s *AddressScraper) Scrape(url string) (Addresses, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	err = s.parseRetinalScreeningURL(res.Body)
	if err != nil {
		return nil, err
	}

	return s.Addresses, err
}

func (s *AddressScraper) parseRetinalScreeningURL(body io.Reader) error {
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}

	document.Find("div.loclist-address").Each(func(i int, selection *goquery.Selection) {
		html, _ := selection.Html()
		addWithoutBr := strings.Replace(html, "<br/>", " ", -1)
		s.Addresses = append(s.Addresses, addWithoutBr)
	})

	return nil
}
