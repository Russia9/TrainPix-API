package parse

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
	"trainpix-api/object"
)

func PhotoSearch(query string, count int, quick bool, params map[string]string) ([]*object.Photo, int, error) {
	var result []*object.Photo

	uri := "https://trainpix.org/search.php?"
	for key := range params {
		uri = uri + "&" + key + "=" + params[key]
	}
	uri = uri + "&num=" + url.QueryEscape(query)

	document, err := GetPage(uri)

	if err != nil {
		return nil, 0, err
	}

	countFound, _ := strconv.Atoi(document.Find(":contains('Найдено изображений:')").Parent().Last().Find("b").Text())

	document.Find(".prw").Each(func(i int, selection *goquery.Selection) {
		if i > count {
			return
		}

		photoURI, errBool := selection.Attr("href")
		if !errBool {
			return
		}

		id, err := strconv.Atoi(strings.Split(photoURI, "/")[2])
		if err != nil {
			return
		}

		photo, _ := PhotoGet(id, quick)
		result = append(result, photo)
	})

	return result, countFound, nil
}

func PhotoGet(id int, quick bool) (*object.Photo, error) {
	pageURI := "https://trainpix.org/photo/" + strconv.Itoa(id) + "/"
	imageURI := "https://trainpix.org/photo" + getIDString(id) + "/" + strconv.Itoa(id) + ".jpg"
	thumbnailURI := "https://trainpix.org/photo" + getIDString(id) + "/" + strconv.Itoa(id) + "_s.jpg"

	var date *string
	var location *string
	var author *string
	var authorLink *string

	if !quick {
		document, err := GetPage(pageURI)
		if err != nil {
			return nil, err
		}

		if document.Find(":contains('Изображение не найдено')").Size() > 0 {
			return nil, errors.New("404")
		}

		authorElement := document.Find("span.cmt_aname").Find("a").First()
		authorName := authorElement.Text()
		author = &authorName
		authorURI, _ := authorElement.Attr("href")
		authorURI = "https://trainpix.org" + authorURI
		authorLink = &authorURI

		locationText := document.Find("center").Find("b").First().Text()
		location = &locationText

		dateText := document.Find("span.cmt_aname").Parent().Find("b").Last().Text()
		date = &dateText
	}

	return &object.Photo{
		ID:         id,
		Image:      imageURI,
		Thumbnail:  thumbnailURI,
		Page:       pageURI,
		Date:       date,
		Location:   location,
		Author:     author,
		AuthorLink: authorLink,
	}, nil
}

func RandomPhotoGet() (*object.Photo, *object.Train, error) {
	uri := "https://trainpix.org/ph.php"
	document, err := GetPage(uri)
	if err != nil {
		return nil, nil, err
	}
	photoURI, _ := document.Find("#ph").Attr("src")

	photoID, err := strconv.Atoi(strings.Split(strings.Split(photoURI, "/")[5], ".")[0])
	if err != nil {
		return nil, nil, err
	}

	trainURI, _ := document.Find(".pwrite").First().Find("a").Attr("href")

	trainID, err := strconv.Atoi(strings.Split(trainURI, "/")[2])
	if err != nil {
		return nil, nil, err
	}

	trainObject, err := TrainGet(trainID, 5, true)
	if err != nil {
		return nil, nil, err
	}

	photoObject, err := PhotoGet(photoID, false)
	if err != nil {
		return nil, nil, err
	}

	return photoObject, trainObject, nil
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
