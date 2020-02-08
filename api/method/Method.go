package method

import "github.com/sirupsen/logrus"

type Method interface {
	GetGroup() string
	GetAlias() string
	Process(params map[string]string, logger logrus.Logger) error
}

