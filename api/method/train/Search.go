package train

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

func getParams(v url.Values) map[string]string {
	params := make(map[string]string)

	if v.Get("state") != "" {
		params["state"] = v.Get("state")
	}
	if v.Get("order") != "" {
		params["order"] = v.Get("order")
	}
	if v.Get("st") != "" {
		params["st"] = v.Get("st")
	}
	return params
}
