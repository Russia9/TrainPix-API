package photo

import (
	"net/url"
	"trainpix-api/api/response"
)

type Random struct {
	Group  string
	Method string
}

func (object Random) GetGroup() string {
	return object.Group
}

func (object Random) GetMethod() string {
	return object.Method
}

func (object Random) Process(params url.Values) response.Response {

}
