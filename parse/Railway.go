package parse

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
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

func RailwayGet(id int, depotCount int, trainCount int, photoCount int, quick bool) (*object.Railway, error) {
	var depots []*object.Depot
	uri := "https://trainpix.org/railway/" + strconv.Itoa(id)
	document, err := GetPage(uri)
	if err != nil {
		return nil, err
	}

	name := document.Find("h1").Text()

	document.Find(":contains('Списки подвижного состава')").Parent().Last().Parent().Last().Find("b").Find("a").Each(func(i int, selection *goquery.Selection) {
		if i > depotCount {
			return
		}

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

		depots = append(depots, depotGet)
	})

	return &object.Railway{
		ID:        id,
		Name:      name,
		DepotList: &depots,
	}, nil
}

func RailwaySearch(query string, depotCount int, trainCount int, photoCount int, quick bool) ([]*object.Railway, int, error) {
	uri := "https://trainpix.org/ajax2.php?action=get-cities&term=" + query
	var jsonResponse []railway
	var result []*object.Railway

	response, err := http.Get(uri)
	if err != nil {
		return nil, 0, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	json.Unmarshal(body, &jsonResponse)

	found := 0

	for _, element := range jsonResponse {
		id, _ := strconv.Atoi(element.ID)
		railwayObject, err := RailwayGet(id, depotCount, trainCount, photoCount, quick)
		if err != nil {
			return nil, 0, err
		}

		result = append(result, railwayObject)
		found++
	}

	return result, found, nil
}
