package train

import (
	"net/url"
	"trainpix-api/api/response"
)

type QSearch struct {
	Group  string
	Method string
}

func (object QSearch) GetGroup() string {
	return object.Group
}

func (object QSearch) GetMethod() string {
	return object.Method
}

func (object QSearch) Process(params url.Values) response.Response {

}
