package parse

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"strconv"
	"strings"
	"trainpix-api/object"
)

func DepotSearch(query string, railwayID int, count int, trainCount int, photoCount int, quick bool) ([]*object.Depot, error) {
	var result []*object.Depot
	uri := "https://trainpix.org/railway/" + strconv.Itoa(railwayID)
	document, err := GetPage(uri)

	if err != nil {
		return nil, err
	}

	if document.Find(":contains('Дорога не найдена')").Size() > 0 {
		return nil, errors.New("404")
	}

	iter := 0

	document.Find(":contains('Списки подвижного состава')").Parent().Last().Parent().Last().Find("b").Find("a").Each(func(i int, selection *goquery.Selection) {
		if iter > count {
			return
		}

		if fuzzy.Match(query, selection.Text()) {
			depotURI, errBool := selection.Attr("href")
			if !errBool {
				return
			}

			depotID, err := strconv.Atoi(strings.Split(depotURI, "=")[1])
			if err != nil {
				return
			}

			depotGet, err := DepotGet(depotID, trainCount, photoCount, quick)
			if err != nil {
				return
			}

			result = append(result, depotGet)
			iter++
		}
	})

	return result, nil
}

func DepotGet(id int, trainCount int, photoCount int, quick bool) (*object.Depot, error) {
	var trains []*object.Train
	uri := "https://trainpix.org/list.php?did=" + strconv.Itoa(id)
	document, err := GetPage(uri)

	if err != nil {
		return nil, err
	}

	if document.Find(":contains('В БД нет записей, удовлетворяющих заданным условиям.')").Size() > 0 {
		return nil, errors.New("404")
	}

	name := document.Find("h2").Find("a").Last().Text()

	iter := 0

	document.Find(".rtable tbody *[class^=s]").Each(func(i int, selection *goquery.Selection) {
		if iter > trainCount {
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
				ID:        id,
				Name:      name,
				Condition: condition,
			}
		} else {
			trainElement, _ = TrainGet(id, photoCount, quick)
		}

		trains = append(trains, trainElement)
		iter++
	})

	return &object.Depot{
		ID:        id,
		Name:      name,
		TrainList: &trains,
	}, nil
}
