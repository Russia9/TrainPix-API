package parser

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"trainpix-api/object"
)

func DepotGet(id int, trainCount int, quick bool) (*object.Depot, error) {
	stringID := strconv.Itoa(id)
	depotURI := "https://trainpix.org/list.php?did=" + stringID
	depotDocument, err := GetPage(depotURI)
	if err != nil {
		return nil, err
	}

	var trains []*object.Train

	if depotDocument.Find(":contains('В БД нет записей, удовлетворяющих заданным условиям.')").Size() > 0 {
		return nil, errors.New("404")
	}

	name := depotDocument.Find("h2").Text()

	depotDocument.Find(".rtable tbody *[class^=s]").Each(func(i int, selection *goquery.Selection) {
		if i > trainCount {
			return
		}

		if selection.Find(":contains('№')").Size() > 0 {
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

		var trainElement *object.Train

		if quick {
			trainElement = &object.Train{
				Id:        id,
				Name:      name,
				Condition: condition,
			}
		} else {
			trainElement, _ = TrainGet(id, false)
		}

		trains = append(trains, trainElement)
	})

	return &object.Depot{
		Id:        id,
		Name:      name,
		TrainList: &trains,
	}, nil
}

func DepotSearch(query string, depotId int) {

}