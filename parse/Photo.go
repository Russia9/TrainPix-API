package parse

import (
	"strconv"
	"trainpix-api/object"
)

func PhotoGet(id int, quick bool) (*object.Photo, error) {

}

func RandomPhotoGet() (*object.Photo, *object.Train, error) {

}

func PhotoSearch(query string, count int, params map[string]string) (*[]*object.Photo, int, error) {

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
