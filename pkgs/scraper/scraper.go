package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type Addresses []string

func Scrape(url string) (Addresses, error) {
	res, err := http.Get(url)
	if err != nil {
		panic("oh no")
	}

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var addresses Addresses

	document.Find("div.loclist-address").Each(func(i int, selection *goquery.Selection) {
		html, _ := selection.Html()
		addWithoutBr := strings.Replace(html, "<br/>", " ", -1)
		addresses = append(addresses, addWithoutBr)
	})

	return addresses, err

}