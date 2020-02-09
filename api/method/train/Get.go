package train

import (
	"net/url"
	"trainpix-api/api/response"
)

type Get struct {
	Group  string
	Method string
}

func (object Get) GetGroup() string {
	return object.Group
}

func (object Get) GetMethod() string {
	return object.Method
}

func (object Get) Process(params url.Values) response.Response {

}
