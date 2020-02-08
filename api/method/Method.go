package method

import (
	"net/url"
	"trainpix-api/api/response"
)

type Method interface {
	GetGroup() string
	GetMethod() string
	Process(params url.Values) (response.Response, error)
}
