package parser

import (
	"errors"
	"strconv"
	"trainpix-api/object/photo"
)

func PhotoGet(id int, quick bool) (photo.Photo, error) {
	pageLink := "https://trainpix.org/photo/" + strconv.Itoa(id) + "/"
	imageLink := "https://trainpix.org/photo" + getIDString(id) + "/" + strconv.Itoa(id) + ".jpg"
	thumbnailLink := "https://trainpix.org/photo" + getIDString(id) + "/" + strconv.Itoa(id) + "_s.jpg"
	date := ""
	location := ""
	author := ""
	authorLink := ""

	if !quick {
		photoDocument, err := GetPage(pageLink)
		if err != nil {
			return photo.Photo{}, err
		}

		if photoDocument.Find(":contains('Изображение не найдено')").Size() > 0 {
			return photo.Photo{}, errors.New("404")
		}

		authorElement := photoDocument.Find("span.cmt_aname").Find("a").First()
		author = authorElement.Text()
		authorLink, _ = authorElement.Attr("href")
		authorLink = "https://trainpix.org" + authorLink

		location = photoDocument.Find("center").Find("b").First().Text()
		date = photoDocument.Find("span.cmt_aname").Parent().Find("b").Last().Text()
	}

	return photo.Photo{
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
