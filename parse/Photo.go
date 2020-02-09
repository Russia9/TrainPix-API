package parse

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
	"trainpix-api/object"
)

func PhotoGet(id int, quick bool) (*object.Photo, error) {
	pageLink := "https://trainpix.org/photo/" + strconv.Itoa(id) + "/"
	imageLink := "https://trainpix.org/photo" + getIDString(id) + "/" + strconv.Itoa(id) + ".jpg"
	thumbnailLink := "https://trainpix.org/photo" + getIDString(id) + "/" + strconv.Itoa(id) + "_s.jpg"
	var date *string
	var location *string
	var author *string
	var authorLink *string

	if !quick {
		photoDocument, err := GetPage(pageLink)
		if err != nil {
			return nil, err
		}

		if photoDocument.Find(":contains('Изображение не найдено')").Size() > 0 {
			return nil, errors.New("404")
		}

		authorElement := photoDocument.Find("span.cmt_aname").Find("a").First()
		authorName := authorElement.Text()
		author = &authorName
		authorURI, _ := authorElement.Attr("href")
		authorURI = "https://trainpix.org" + authorURI
		authorLink = &authorURI

		locationText := photoDocument.Find("center").Find("b").First().Text()
		location = &locationText

		dateText := photoDocument.Find("span.cmt_aname").Parent().Find("b").Last().Text()
		date = &dateText
	}

	return &object.Photo{
		ID:         id,
		Image:      imageLink,
		Thumbnail:  thumbnailLink,
		Page:       pageLink,
		Date:       date,
		Location:   location,
		Author:     author,
		AuthorLink: authorLink,
	}, nil
}

func RandomPhotoGet() (*object.Photo, *object.Train, error) {
	pageLink := "https://trainpix.org/ph.php"
	pageDocument, err := GetPage(pageLink)
	if err != nil {
		return nil, nil, err
	}
	photoURI, _ := pageDocument.Find("#ph").Attr("src")

	photoID, err := strconv.Atoi(strings.Split(strings.Split(photoURI, "/")[5], ".")[0])
	if err != nil {
		return nil, nil, err
	}

	trainURI, _ := pageDocument.Find(".pwrite").First().Find("a").Attr("href")

	trainID, err := strconv.Atoi(strings.Split(trainURI, "/")[2])
	if err != nil {
		return nil, nil, err
	}

	trainObject, err := TrainGet(trainID, true)
	if err != nil {
		return nil, nil, err
	}

	photoObject, err := PhotoGet(photoID, false)
	if err != nil {
		return nil, nil, err
	}

	return photoObject, trainObject, nil
}

func PhotoSearch(query string, count int, params map[string]string) (*[]*object.Photo, int, error) {
	searchURI := "https://trainpix.org/search.php?"
	for key := range params {
		searchURI = searchURI + "&" + key + "=" + params[key]
	}
	searchURI = searchURI + "&num=" + url.QueryEscape(query)

	searchDocument, err := GetPage(searchURI)
	if err != nil {
		return nil, 0, err
	}

	countFound, _ := strconv.Atoi(searchDocument.Find(".main").Find("b").First().Text())

	var result []*object.Photo

	searchDocument.Find("a.prw").Each(func(i int, selection *goquery.Selection) {
		if i > count {
			return
		}

		link, errBoolean := selection.Attr("href")
		if !errBoolean {
			return
		}

		id, err := strconv.Atoi(strings.Split(link, "/")[2])
		if err != nil {
			return
		}
		photoObject, err := PhotoGet(id, false)
		if err != nil {
			return
		}

		result = append(result, photoObject)
	})

	return &result, countFound, nil
}

func getIDString(id int) string {
	strID := strconv.Itoa(id)
	idLen := len(strID) - 1
	first := strID[:idLen]
	str := ""
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			str += "/"
		}
		if i < 6-idLen {
			str += "0"
		} else {
			str += string(first[i-(6-idLen)])
		}
	}
	return str
}
