package train

import (
	"errors"
	"strconv"
)

type Get struct {
	group string
	alias string
}

func (object Get) GetAlias() string {
	return object.alias
}

func (object Get) Process(params map[string]string) error {
	id := -1
	var err error

	stringId, contains := params["id"]
	if contains {
		id, err = strconv.Atoi(stringId)

		if err != nil {
			return errors.New("400")
		}
	} else {
		return errors.New("400")
	}

	return nil
}
