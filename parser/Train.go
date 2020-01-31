package parser

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
	"trainpix-api/object/infrastructure"
	"trainpix-api/object/train"
)

func TrainSearch(query string, count int, quick bool) ([]train.Train, error) {
	searchURI := "https://trainpix.org/vsearch.php?num=" + url.QueryEscape(query)
	searchDocument, err := GetPage(searchURI)
	if err != nil {
		return nil, err
	}

	var result []train.Train

	iter := 0

	searchDocument.Find(".rtable tbody *[class^=s]").Each(func(i int, selection *goquery.Selection) {
		if iter >= count {
			return
		}
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

		var trainElement train.Train

		if quick {
			trainElement = train.Train{
				Id:                   id,
				Name:                 name,
				Railway:              infrastructure.Railway{},
				Depot:                infrastructure.Depot{},
				Model:                train.Model{},
				Builder:              "",
				IdentificationNumber: "",
				SerialType:           "",
				Built:                "",
				Category:             "",
				Condition:            condition,
				Note:                 "",
				PhotoList:            nil,
			}
		} else {
			trainElement, _ = TrainGet(id)
		}

		result = append(result, trainElement)
		iter++
	})

	return result, nil
}

func TrainGet(id int) (train.Train, error) {
	stringID := strconv.Itoa(id)
	trainURI := "https://trainpix.org/vehicle/" + stringID + "/"
	searchDocument, err := GetPage(trainURI)
	if err != nil {
		return train.Train{}, err
	}

	if searchDocument.Find(":contains('Подвижной состав не найден')").Size() > 0 {
		return train.Train{}, errors.New("404")
	}

	name := searchDocument.Find("h1").First().Text()
	var railway infrastructure.Railway
	var depot infrastructure.Depot
	var model train.Model
	var builder string
	var identificationNumber string
	var serialType string
	var built string
	var category string
	condition := 1
	var note string

	searchDocument.Find(".h21").Each(func(i int, selection *goquery.Selection) {
		if selection.Children().Size() > 1 {
			key := selection.Find(".ds").Text()
			switch key {
			case "Дорога приписки:":
				linkElement := selection.Find("a")
				link, _ := linkElement.Attr("href")
				elementId, _ := strconv.Atoi(strings.Split(link, "/")[2])
				railway = infrastructure.Railway{
					Id:   elementId,
					Name: linkElement.Text(),
				}
				break
			case "Депо:":
				linkElement := selection.Find("a")
				link, _ := linkElement.Attr("href")
				elementId, _ := strconv.Atoi(strings.Split(link, "=")[1])
				depot = infrastructure.Depot{
					Id:   elementId,
					Name: linkElement.Text(),
				}
				break
			case "Серия:":
				linkElement := selection.Find("a")
				link, _ := linkElement.Attr("href")
				elementId, _ := strconv.Atoi(strings.Split(link, "=")[1])
				model = train.Model{
					Id:   elementId,
					Name: linkElement.Text(),
				}
				break
			case "Завод-изготовитель:":
				builder = selection.Find(".d").Text()
				break
			case "Сетевой №:":
				identificationNumber = selection.Find(".d").Text()
				break
			case "Заводской тип:":
				serialType = selection.Find(".d").Text()
				break
			case "Построен:":
				built = selection.Find(".d").Text()
				break
			case "Категория:":
				category = selection.Find(".d").Text()
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
				note = selection.Find(".d").Text()
				break
			}
		}
	})

	return train.Train{
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
		PhotoList:            nil,
	}, nil
}
