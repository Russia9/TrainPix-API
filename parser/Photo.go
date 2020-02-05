package parser

import (
	"errors"
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
		Id:         id,
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
	photoUrl, _ := pageDocument.Find("#ph").Attr("src")

	photoId, err := strconv.Atoi(strings.Split(strings.Split(photoUrl, "/")[5], ".")[0])
	if err != nil {
		return nil, nil, err
	}

	trainUrl, _ := pageDocument.Find(".pwrite").First().Find("a").Attr("href")

	trainId, err := strconv.Atoi(strings.Split(trainUrl, "/")[2])
	if err != nil {
		return nil, nil, err
	}

	trainObject, err := TrainGet(trainId, true)
	if err != nil {
		return nil, nil, err
	}

	photoObject, err := PhotoGet(photoId, false)
	if err != nil {
		return nil, nil, err
	}

	return photoObject, trainObject, nil
}

func getIDString(id int) string {
	strId := strconv.Itoa(id)
	idLen := len(strId) - 1
	first := strId[:idLen]
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
