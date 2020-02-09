package parse

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"strconv"
	"strings"
	"trainpix-api/object"
)

func DepotSearch(query string, railwayID int, count int, trains int, quick bool) ([]*object.Depot, error) {
	railwayURI := "https://trainpix.org/railway/" + strconv.Itoa(railwayID) + "/"
	railwayDocument, err := GetPage(railwayURI)
	if err != nil {
		return nil, err
	}

	num := 0

	var result []*object.Depot

	listElement := railwayDocument.Find(":contains('Списки подвижного состава')").Parent().Last().Parent().Last()
	listElement.Find("b").Find("a").Each(func(i int, selection *goquery.Selection) {
		if num >= count {
			return
		}

		if fuzzy.Match(query, selection.Text()) {
			depotURI, _ := selection.Attr("href")

			id, err := strconv.Atoi(strings.Split(depotURI, "=")[1])
			if err != nil {
				return
			}

			depot, err := DepotGet(id, trains, quick)
			if err != nil {
				return
			}

			result = append(result, depot)
			num++
		}
	})

	return result, nil
}

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
				ID:        id,
				Name:      name,
				Condition: condition,
			}
		} else {
			trainElement, _ = TrainGet(id, false)
		}

		trains = append(trains, trainElement)
	})

	return &object.Depot{
		ID:        id,
		Name:      name,
		TrainList: &trains,
	}, nil
}
