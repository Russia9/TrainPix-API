package parse

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func GetPage(URI string) (*goquery.Document, error) {
	page, err := http.Get(URI)
	if err != nil {
		return nil, err
	}

	defer page.Body.Close()
	if page.StatusCode != 200 {
		return nil, errors.New("Status code:" + string(page.StatusCode))
	}

	document, err := goquery.NewDocumentFromReader(page.Body)

	return document, err
}

