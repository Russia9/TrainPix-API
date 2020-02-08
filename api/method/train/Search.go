package train

import (
	"errors"
	"strconv"
)

type Search struct {
	group string
	alias string
}

func (object Search) GetGroup() string {
	return object.group
}

func (object Search) GetAlias() string {
	return object.alias
}

func (object Search) Process(params map[string]string) error {
	var err error

	return nil
}
