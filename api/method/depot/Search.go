package depot

import (
	"net/url"
	"trainpix-api/api/response"
)

type Search struct {
	Group  string
	Method string
}

func (object Search) GetGroup() string {
	return object.Group
}

func (object Search) GetMethod() string {
	return object.Method
}

func (object Search) Process(params url.Values) response.Response {

}
