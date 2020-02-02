package parser

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
	"trainpix-api/object/infrastructure"
	"trainpix-api/object/photo"
	"trainpix-api/object/train"
)

func TrainSearch(query string, count int, quick bool, params map[string]string) ([]*train.Train, int, int, error) {
	searchURI := "https://trainpix.org/vsearch.php?"
	for key := range params {
		searchURI = searchURI + "&" + key + "=" + params[key]
	}
	searchURI = searchURI + "&num=" + url.QueryEscape(query)

	searchDocument, err := GetPage(searchURI)
	if err != nil {
		return nil, 0, 0, err
	}

	if searchDocument.Find(":contains('Ничего не найдено.')").Size() > 0 {
		return nil, 0, 0, errors.New("404")
	}

	countFound, _ := strconv.Atoi(searchDocument.Find(".main").Find("b").First().Text())

	var result []*train.Train

	iter := 0
	countParsed := 0

	searchDocument.Find(".rtable tbody *[class^=s]").Each(func(i int, selection *goquery.Selection) {
		if iter >= count {
			return
		}
		countParsed++

		idString, _ := selection.Find("a").First().Attr("href")
		id, _ := strconv.Atoi(strings.Split(idString, "/")[2])

		conditionClass, _ := selection.Attr("class")
		condition, _ := strconv.Atoi(conditionClass[1:])
		if condition > 10 {
			condition -= 10
		}
		if condition == 6 || condition == 8 {
			return
		}

		name := selection.Find("a").Text()

		var trainElement *train.Train

		if quick {
			trainElement = &train.Train{
				Id:        id,
				Name:      name,
				Condition: condition,
			}
		} else {
			trainElement, _ = TrainGet(id, true)
		}

		result = append(result, trainElement)
		iter++
	})

	return result, countFound, countParsed, nil
}

func TrainGet(id int, quick bool) (*train.Train, error) {
	stringID := strconv.Itoa(id)
	trainURI := "https://trainpix.org/vehicle/" + stringID + "/"
	searchDocument, err := GetPage(trainURI)
	if err != nil {
		return nil, err
	}

	if searchDocument.Find(":contains('Подвижной состав не найден')").Size() > 0 {
		return nil, errors.New("404")
	}

	name := searchDocument.Find("h1").First().Text()
	var railway infrastructure.Railway
	var depot infrastructure.Depot
	var model train.Model
	var builder *string
	var identificationNumber *string
	var serialType *string
	var built *string
	var category *string
	condition := 1
	var note *string
	var info *string
	var photoList []*photo.Photo

	searchDocument.Find(".p0.horlines").Find(".h21").Each(func(i int, selection *goquery.Selection) {
		if selection.Children().Size() > 1 {
			key := selection.Find(".ds").Text()
			switch key {
			case "Дорога приписки:":
				linkElement := selection.Find("a")
				link, _ := linkElement.Attr("href")
				elementId, _ := strconv.Atoi(strings.Split(link, "/")[2])
				railwayName := linkElement.Text()
				railway = infrastructure.Railway{
					Id:   elementId,
					Name: railwayName,
				}
				break
			case "Депо:":
				linkElement := selection.Find("a")
				link, _ := linkElement.Attr("href")
				elementId, _ := strconv.Atoi(strings.Split(link, "=")[1])
				depotName := linkElement.Text()
				depot = infrastructure.Depot{
					Id:   elementId,
					Name: depotName,
				}
				break
			case "Серия:":
				linkElement := selection.Find("a")
				link, _ := linkElement.Attr("href")
				elementId, _ := strconv.Atoi(strings.Split(link, "=")[1])
				modelName := linkElement.Text()
				model = train.Model{
					Id:   elementId,
					Name: modelName,
				}
				break
			case "Завод-изготовитель:":
				builderText := selection.Find(".d").Text()
				builder = &builderText
				break
			case "Сетевой №:":
				identificationNumberText := selection.Find(".d").Text()
				identificationNumber = &identificationNumberText
				break
			case "Заводской тип:":
				serialTypeText := selection.Find(".d").Text()
				serialType = &serialTypeText
				break
			case "Построен:":
				builtText := selection.Find(".d").Text()
				built = &builtText
				break
			case "Категория:":
				categoryText := selection.Find(".d").Text()
				category = &categoryText
				break
			case "Текущее состояние:":
				conditionClass, err := selection.Find("span").First().Attr("class")
				if err == true {
					value, _ := strconv.Atoi(conditionClass[1:])
					if value > 10 {
						value -= 10
					}
					condition = value
				}
				break
			case "Примечание:":
				noteText := selection.Find(".d").Text()
				note = &noteText
				break
			}
		} else {
			infoText := selection.Find(".d").Text()
			info = &infoText
		}
	})

	searchDocument.Find(".prw").Each(func(i int, selection *goquery.Selection) {
		href, status := selection.Attr("href")
		if status == false {
			return
		}

		id, err := strconv.Atoi(strings.Split(href, "/")[2])
		if err != nil {
			return
		}

		trainPhoto, err := PhotoGet(id, quick)
		photoList = append(photoList, trainPhoto)
	})

	return &train.Train{
		Id:                   id,
		Name:                 name,
		Railway:              railway,
		Depot:                depot,
		Model:                model,
		Builder:              builder,
		IdentificationNumber: identificationNumber,
		SerialType:           serialType,
		Built:                built,
		Category:             category,
		Condition:            condition,
		Note:                 note,
		Info:                 info,
		PhotoList:            photoList,
	}, nil
}
