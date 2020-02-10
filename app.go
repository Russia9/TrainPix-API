package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"trainpix-api/api"
)

var logger = logrus.New()

func main() {
	logger.Out = os.Stdout

	logger.Info("TrainPix API By Russia9")
	logger.Info("More info: https://github.com/Russia9/TrainPix-API")

	switch os.Getenv("API_LEVEL") {
	case "TRACE":
		logger.SetLevel(logrus.TraceLevel)
		break
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
		break
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
		break
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
		break
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
		break
	case "FATAL":
		logger.SetLevel(logrus.FatalLevel)
		break
	case "PANIC":
		logger.SetLevel(logrus.PanicLevel)
		break
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	port := 8080
	if os.Getenv("API_PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("API_PORT"))
		logger.Trace("Detected PORT: ", port)
	} else {
		logger.Trace("Port not specified, using: ", port)
	}

	logger.Debug("Starting server at port: ", port)

	api.Route(port, logger)
}
