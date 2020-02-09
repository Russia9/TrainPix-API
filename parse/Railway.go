package parse

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"trainpix-api/object"
)

type railway struct {
	ID    string `json:"value"`
	Label string `json:"label"`
	RID   string `json:"rid"`
	RName string `json:"rname"`
}

func RailwayGet(id int, count int, trains int, quick bool) (*object.Railway, error) {
	railwayURI := "https://trainpix.org/railway/" + strconv.Itoa(id)
	railwayDocument, err := GetPage(railwayURI)
	if err != nil {
		return nil, err
	}

	name := railwayDocument.Find("h1").Text()

	num := 0

	var result []*object.Depot

	listElement := railwayDocument.Find(":contains('Списки подвижного состава')").Parent().Last().Parent().Last()
	listElement.Find("b").Find("a").Each(func(i int, selection *goquery.Selection) {
		if num >= count {
			return
		}

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
	})

	return &object.Railway{
		ID:        id,
		Name:      name,
		DepotList: &result,
	}, nil
}

func RailwaySearch(query string) {
	var railways []railway
	queryURI := "https://trainpix.org/ajax2.php?action=get-cities&term=" + url.QueryEscape(query)
	queryResult, err := http.Get(queryURI)
	if err != nil {
		return
	}

	queryData, err := ioutil.ReadAll(queryResult.Body)
	if err != nil {
		return
	}

	json.Unmarshal(queryData, &railways)

	fmt.Println(railways[0].Label)
}
